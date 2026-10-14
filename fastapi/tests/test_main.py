import platform

import fastapi
from fastapi.testclient import TestClient

from app.main import app

client = TestClient(app)


def test_the_welcome_page_responds() -> None:
    response = client.get("/")

    assert response.status_code == 200
    assert "Your FastAPI app is up and running on Laravel Cloud." in response.text
    assert f"FastAPI {fastapi.__version__} · Python {platform.python_version()}" in response.text


def test_the_api_responds() -> None:
    response = client.get("/api")

    assert response.status_code == 200
    assert response.json() == {"message": "Hello from Laravel Cloud"}
