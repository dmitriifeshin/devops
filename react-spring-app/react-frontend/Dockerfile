FROM node:20-alpine AS build

WORKDIR /app
COPY package*.json ./
RUN npm install --frozen-lockfile && npm cache clean --force
COPY . .
RUN npm run build
