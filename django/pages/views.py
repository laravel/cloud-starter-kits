import platform

import django
from django.http import HttpRequest, HttpResponse
from django.shortcuts import render


def home(request: HttpRequest) -> HttpResponse:
    return render(
        request,
        "pages/home.html",
        {"version": django.get_version(), "python_version": platform.python_version()},
    )
