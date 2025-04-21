FROM python:3.13-alpine

WORKDIR /app

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    DATABASE_URL=postgresql://postgres:postgres@db:5432/taskdb

RUN apk add --no-cache libpq && \
    adduser -D -u 1000 appuser

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

RUN chown -R appuser:appuser /app
USER appuser

ENTRYPOINT [ "python3", "app.py" ]
CMD ["alembic", "upgrade", "head"]
