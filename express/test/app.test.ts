import assert from 'node:assert/strict';
import type { Server } from 'node:http';
import type { AddressInfo } from 'node:net';
import { after, before, test } from 'node:test';
import { app } from '../src/app.ts';

let server: Server;
let baseUrl: string;

before(async () => {
    server = app.listen(0);
    await new Promise((resolve) => server.once('listening', resolve));
    baseUrl = `http://localhost:${(server.address() as AddressInfo).port}`;
});

after(() => {
    server.close();
});

test('the welcome page responds', async () => {
    const response = await fetch(`${baseUrl}/`);

    assert.equal(response.status, 200);
    assert.match(await response.text(), /Your Express app is up and running on Laravel Cloud\./);
});

test('the api responds', async () => {
    const response = await fetch(`${baseUrl}/api`);

    assert.equal(response.status, 200);
    assert.deepEqual(await response.json(), { message: 'Hello from Laravel Cloud' });
});
