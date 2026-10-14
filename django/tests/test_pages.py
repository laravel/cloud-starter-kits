from django.test import Client
from django.urls import reverse


def test_the_home_page_responds(client: Client) -> None:
    response = client.get(reverse("home"))

    assert response.status_code == 200
    assert b"Hello from Laravel Cloud" in response.content
