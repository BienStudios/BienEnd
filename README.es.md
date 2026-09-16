# BienEnd - Backend Web Simplificado

## Contenido:

- [Flujo de ejecución con Docker](#flujo-de-ejecución-con-docker)

## Flujo de ejecución con Docker

Cuando los microservicios se adapten a las necesidades del proyecto, desplegar primeramente el contenedor compartido dentro del Servidor:

```bash
cd /infra-core # Asegúrate de estar accediendo al directorio correcto
# Acceder al archivo (-f) "docker-compose.shared.yml" y levantarlo en segundo plano
# (-d) para no detenerlo al cerrar la terminal.
docker compose -f docker-compose.shared.yml up -d
```

Seguidamente, iniciar las Aplicaciones/Clientes según sea el caso:

```bash
cd /app1-client # O como se haya nombrado la aplicación
docker compose -f docker-compose.app.yml up -d
```
