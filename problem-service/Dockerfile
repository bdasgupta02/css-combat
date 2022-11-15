FROM python:3.9-slim

WORKDIR /problem-service

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

COPY . .

RUN python -m pip install --upgrade pip

RUN apt-get update && apt-get install -y --no-install-recommends gcc

RUN pip install pipenv && pipenv install --dev --system --deploy

RUN pipenv install

RUN adduser -u 5678 --disabled-password --gecos "" appuser && chown -R appuser /problem-service

USER appuser

EXPOSE 8040

CMD ["python", "server.py"]