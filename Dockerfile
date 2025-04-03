FROM alpine:3.20.6

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN apk update && \
    apk upgrade && \
    apk add nginx && \
    mkdir -p /run/nginx && \
    mkdir -p /var/www/html/my_site && \
    mkdir -p /etc/nginx/conf.d && \
    mkdir -p /var/lib/nginx/tmp && \
    mkdir -p /var/lib/nginx/logs

RUN chown -R appuser:appgroup /var/lib/nginx && \
    chown -R appuser:appgroup /run/nginx && \
    chown -R appuser:appgroup /var/log/nginx

EXPOSE 80
VOLUME [ "/var/www/html/my_site" ]

COPY nginx.conf /etc/nginx/nginx.conf

USER appuser

ENTRYPOINT ["nginx", "-g", "daemon off;"]