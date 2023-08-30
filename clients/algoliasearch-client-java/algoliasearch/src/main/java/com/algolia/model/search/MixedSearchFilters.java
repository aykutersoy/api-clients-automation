// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

import com.algolia.exceptions.AlgoliaRuntimeException;
import com.algolia.utils.CompoundType;
import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.core.*;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.deser.std.StdDeserializer;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import java.io.IOException;
import java.util.List;
import java.util.logging.Logger;

/** MixedSearchFilters */
@JsonDeserialize(using = MixedSearchFilters.Deserializer.class)
@JsonSerialize(using = MixedSearchFilters.Serializer.class)
public interface MixedSearchFilters<T> extends CompoundType<T> {
  static MixedSearchFilters<List<String>> of(List<String> inside) {
    return new MixedSearchFiltersListOfString(inside);
  }

  static MixedSearchFilters<String> of(String inside) {
    return new MixedSearchFiltersString(inside);
  }

  class Serializer extends StdSerializer<MixedSearchFilters> {

    public Serializer(Class<MixedSearchFilters> t) {
      super(t);
    }

    public Serializer() {
      this(null);
    }

    @Override
    public void serialize(MixedSearchFilters value, JsonGenerator jgen, SerializerProvider provider) throws IOException {
      jgen.writeObject(value.get());
    }
  }

  class Deserializer extends StdDeserializer<MixedSearchFilters> {

    private static final Logger LOGGER = Logger.getLogger(Deserializer.class.getName());

    public Deserializer() {
      this(MixedSearchFilters.class);
    }

    public Deserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public MixedSearchFilters deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException {
      JsonNode tree = jp.readValueAsTree();

      // deserialize List<String>
      if (tree.isArray()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          List<String> value = parser.readValueAs(new TypeReference<List<String>>() {});
          return MixedSearchFilters.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf List<String> (error: " + e.getMessage() + ") (type: List<String>)");
        }
      }

      // deserialize String
      if (tree.isValueNode()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          String value = parser.readValueAs(new TypeReference<String>() {});
          return MixedSearchFilters.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf String (error: " + e.getMessage() + ") (type: String)");
        }
      }
      throw new AlgoliaRuntimeException(String.format("Failed to deserialize json element: %s", tree));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public MixedSearchFilters getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "MixedSearchFilters cannot be null");
    }
  }
}

class MixedSearchFiltersListOfString implements MixedSearchFilters<List<String>> {

  private final List<String> value;

  MixedSearchFiltersListOfString(List<String> value) {
    this.value = value;
  }

  @Override
  public List<String> get() {
    return value;
  }
}

class MixedSearchFiltersString implements MixedSearchFilters<String> {

  private final String value;

  MixedSearchFiltersString(String value) {
    this.value = value;
  }

  @Override
  public String get() {
    return value;
  }
}
