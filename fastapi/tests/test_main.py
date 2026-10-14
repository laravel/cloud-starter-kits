from fastapi.testclient import TestClient

from app.main import app

client = TestClient(app)


def test_the_home_page_responds() -> None:
    response = client.get("/")

    assert response.status_code == 200
    assert response.json() == {"message": "Hello from Laravel Cloud"}
