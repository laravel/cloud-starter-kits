import platform
from pathlib import Path

import fastapi
from fastapi import FastAPI
from fastapi.responses import HTMLResponse

app = FastAPI()

welcome = (
    (Path(__file__).parent / "welcome.html")
    .read_text()
    .replace("{{version}}", fastapi.__version__)
    .replace("{{python_version}}", platform.python_version())
)


@app.get("/", response_class=HTMLResponse, include_in_schema=False)
def home() -> str:
    return welcome


@app.get("/api")
def api() -> dict[str, str]:
    return {"message": "Hello from Laravel Cloud"}
