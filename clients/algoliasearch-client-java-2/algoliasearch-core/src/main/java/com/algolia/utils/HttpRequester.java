package com.algolia.utils;

import com.algolia.exceptions.*;
import com.algolia.utils.retry.RetryStrategy;
import com.algolia.utils.retry.StatefulHost;
import java.io.IOException;
import java.lang.reflect.Type;
import java.util.List;
import java.util.concurrent.TimeUnit;
import okhttp3.Call;
import okhttp3.Interceptor;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.logging.HttpLoggingInterceptor;

public class HttpRequester implements Requester {

  private RetryStrategy retryStrategy;
  private OkHttpClient httpClient;
  private HttpLoggingInterceptor loggingInterceptor;
  private LogLevel level;

  public HttpRequester() {
    this.retryStrategy = new RetryStrategy();

    OkHttpClient.Builder builder = new OkHttpClient.Builder();
    builder.addInterceptor(retryStrategy.getRetryInterceptor());

    this.loggingInterceptor = new HttpLoggingInterceptor();
    loggingInterceptor.setLevel(LogLevel.NONE.value());
    builder.addInterceptor(this.loggingInterceptor);
    
    builder.retryOnConnectionFailure(false);

    httpClient = builder.build();
  }

  @Override
  public Call newCall(Request request) {
    return httpClient.newCall(request);
  }

  @Override
  public <T> T handleResponse(Response response, Type returnType) throws AlgoliaRuntimeException {
    if (response.isSuccessful()) {
      if (returnType == null || response.code() == 204) {
        // returning null if the returnType is not defined, or the status code is 204 (No Content)
        if (response.body() != null) {
          try {
            response.body().close();
          } catch (Exception e) {
            throw new AlgoliaApiException(response.message(), e, response.code());
          }
        }
        return null;
      } else {
        return deserialize(response, returnType);
      }
    } else {
      if (response.body() != null) {
        try {
          response.body().string();
        } catch (IOException e) {
          throw new AlgoliaApiException(response.message(), e, response.code());
        }
      }
      throw new AlgoliaApiException(response.message(), response.code());
    }
  }

  private <T> T deserialize(Response response, Type returnType) throws AlgoliaRuntimeException {
    if (response == null || returnType == null) {
      return null;
    }

    if ("byte[]".equals(returnType.toString())) {
      // Handle binary response (byte array).
      try {
        return (T) response.body().bytes();
      } catch (IOException e) {
        throw new AlgoliaRuntimeException(e);
      }
    }

    String respBody;
    try {
      if (response.body() != null) respBody = response.body().string(); else respBody = null;
    } catch (IOException e) {
      throw new AlgoliaRuntimeException(e);
    }

    if (respBody == null || "".equals(respBody)) {
      return null;
    }

    String contentType = response.headers().get("Content-Type");
    if (contentType == null) {
      contentType = "application/json";
    }

    return JSON.deserialize(respBody, returnType);
  }

  @Override
  public void setLogLevel(LogLevel level) {
    if (level != this.level) {
      this.loggingInterceptor.setLevel(level.value());
    }
    this.level = level;
  }

  @Override
  public int getConnectTimeout() {
    return httpClient.connectTimeoutMillis();
  }

  @Override
  public void setConnectTimeout(int connectionTimeout) {
    httpClient = httpClient.newBuilder().connectTimeout(connectionTimeout, TimeUnit.MILLISECONDS).build();
  }

  @Override
  public int getReadTimeout() {
    return httpClient.readTimeoutMillis();
  }

  public void setReadTimeout(int readTimeout) {
    httpClient = httpClient.newBuilder().readTimeout(readTimeout, TimeUnit.MILLISECONDS).build();
  }

  @Override
  public int getWriteTimeout() {
    return httpClient.writeTimeoutMillis();
  }

  @Override
  public void setWriteTimeout(int writeTimeout) {
    httpClient = httpClient.newBuilder().writeTimeout(writeTimeout, TimeUnit.MILLISECONDS).build();
  }

  @Override
  public void setHosts(List<StatefulHost> hosts) {
    this.retryStrategy.setHosts(hosts);
  }

  public void addInterceptor(Interceptor interceptor) {
    httpClient = httpClient.newBuilder().addInterceptor(interceptor).build();
  }
}
