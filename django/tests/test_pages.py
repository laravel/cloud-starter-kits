import platform

import django
from django.test import Client
from django.urls import reverse


def test_the_welcome_page_responds(client: Client) -> None:
    response = client.get(reverse("home"))

    assert response.status_code == 200
    assert "Your Django app is up and running on Laravel Cloud." in response.text
    assert f"Django {django.get_version()} · Python {platform.python_version()}" in response.text
