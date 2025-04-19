### Dockerfile
FROM python:3.13-alpine

WORKDIR /app

# Runtime ENV variables for consistency
ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    DATABASE_URL=postgresql://postgres:postgres@db:5432/taskdb

RUN apk add --no-cache gcc musl-dev libffi-dev postgresql-dev

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["sh", "-c", "alembic upgrade head && python app.py"]