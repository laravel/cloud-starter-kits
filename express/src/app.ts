import express from 'express';

export const app = express();

app.get('/', (_request, response) => {
    response.json({ message: 'Hello from Laravel Cloud' });
});
