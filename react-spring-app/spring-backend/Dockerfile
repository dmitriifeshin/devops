FROM maven:3.9-eclipse-temurin-17-alpine AS builder

WORKDIR /app
COPY . .
RUN mvn clean package -DskipTests && rm -rf ~/.m2

FROM eclipse-temurin:17-jre-alpine

WORKDIR /app
COPY --from=builder /app/target/*.jar app.jar

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /app
USER appuser

ENTRYPOINT ["java", "-jar", "app.jar"]
