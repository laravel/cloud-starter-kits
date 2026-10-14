import { readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import express from 'express';

const { version } = createRequire(import.meta.url)('express/package.json') as { version: string };

const welcome = readFileSync(new URL('./welcome.html', import.meta.url), 'utf8').replace('{{version}}', version);

export const app = express();

app.get('/', (_request, response) => {
    response.type('html').send(welcome);
});

app.get('/api', (_request, response) => {
    response.json({ message: 'Hello from Laravel Cloud' });
});
