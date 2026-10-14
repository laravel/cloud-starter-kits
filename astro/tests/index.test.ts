import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { expect, test } from 'vitest';
import Index from '../src/pages/index.astro';

test('the welcome page renders', async () => {
    const container = await AstroContainer.create();

    const html = await container.renderToString(Index);

    expect(html).toContain('Your Astro app is up and running on Laravel Cloud.');
});
