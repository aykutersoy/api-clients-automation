// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { ClientOptions } from '@algolia/client-common';
import {
  DEFAULT_CONNECT_TIMEOUT_NODE,
  DEFAULT_READ_TIMEOUT_NODE,
  DEFAULT_WRITE_TIMEOUT_NODE,
  createMemoryCache,
  createNullCache,
} from '@algolia/client-common';
import { createHttpRequester } from '@algolia/requester-node-http';

import type { PredictClient, Region } from '../src/predictClient';
import { createPredictClient, REGIONS } from '../src/predictClient';

export { apiClientVersion, PredictClient } from '../src/predictClient';
export * from '../model';

export function predictClient(
  appId: string,
  apiKey: string,
  region: Region,
  options?: ClientOptions
): PredictClient {
  if (!appId || typeof appId !== 'string') {
    throw new Error('`appId` is missing.');
  }

  if (!apiKey || typeof apiKey !== 'string') {
    throw new Error('`apiKey` is missing.');
  }

  if (!region) {
    throw new Error('`region` is missing.');
  }

  if (region && (typeof region !== 'string' || !REGIONS.includes(region))) {
    throw new Error(
      `\`region\` must be one of the following: ${REGIONS.join(', ')}`
    );
  }

  return createPredictClient({
    appId,
    apiKey,
    region,
    timeouts: {
      connect: DEFAULT_CONNECT_TIMEOUT_NODE,
      read: DEFAULT_READ_TIMEOUT_NODE,
      write: DEFAULT_WRITE_TIMEOUT_NODE,
    },
    requester: createHttpRequester(),
    algoliaAgents: [{ segment: 'Node.js', version: process.versions.node }],
    responsesCache: createNullCache(),
    requestsCache: createNullCache(),
    hostsCache: createMemoryCache(),
    ...options,
  });
}
