import { algoliasearch, apiClientVersion } from '../builds/node';

describe('bundle', () => {
  const client = algoliasearch('APP_ID', 'API_KEY');

  it('sets the ua', () => {
    expect(client.transporter.algoliaAgent).toEqual({
      add: expect.any(Function),
      value: expect.stringContaining(
        `Algolia for JavaScript (${apiClientVersion}); Search (${apiClientVersion}); Node.js`,
      ),
    });
  });
});
