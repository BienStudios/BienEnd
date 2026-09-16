# Empaquetado
FROM gradle:8.7-jdk25 AS build
WORKDIR /app
COPY . .
RUN gradle bootJar --no-daemon

# Ejecución
FROM eclipse-temurin:25-jre-alpine
WORKDIR /app
COPY --from=build /app/build/libs/*.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
