FROM python:3.10-slim

WORKDIR /app

COPY . /app

RUN python -m venv venv \
    && ./venv/bin/pip install --upgrade pip \
    && ./venv/bin/pip install -r requirements.txt

ENV PATH="/app/venv/bin:$PATH"

EXPOSE 80

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "80"]
