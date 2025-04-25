FROM golang:1.24-alpine AS build
WORKDIR /src

COPY digitalclock/ .
RUN go mod download && go build -o digitalclock .

FROM alpine:3.21
WORKDIR /app

RUN apk add --no-cache tzdata

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=build /src/digitalclock .

USER appuser
ENTRYPOINT ["./digitalclock"]
