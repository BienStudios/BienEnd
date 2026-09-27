<!--
================================================================
DEFINICIÓN DE RUTAS. CUALQUIER MODIFICACIÓN DE RUTA O NOMBRE
DE DIRECTORIO DE UN DOCUMENTO EXTERNO APUNTADO, MODIFICAR AQUÍ.
SÓLO DOCUMENTOS EN ESPAÑOL.
================================================================

*appc -> app1-client

*rm -> READ ME

-->
[appc_users_rm]: ./app1-client/users_service/README.es.md
[appc_products_rm]: ./app1-client/products_service/README.es.md
[appc_payment_rm]: ./app1-client/payment_gateway_service/README.es.md
[appc_gateway_rm]: ./app1-client/bff_service/README.es.md

# BienEnd - Backend Web Simplificado

## Contenido

- [Enlaces externos](#enlaces-externos)
- [Resumen de Microservicios](#resumen-de-microservicios)
    - [Usuarios y Autenticación](#usuarios-y-autenticación)
- [Flujo de ejecución con Docker](#flujo-de-ejecución-con-docker)

## Enlaces externos

- [Microservicio de Usuarios y Autenticación][appc_users_rm]
- [Microservicio de Productos][appc_products_rm]
- [Microservicio de Pasarela de Pagos][appc_payment_rm]
- [Microservicio de Backend-For-Frontend (Gateway)][appc_gateway_rm]

## Resumen de Microservicios

### Usuarios y Autenticación

Servicio encargado del registro de usuarios, tanto a nivel local como procedientes de otros proveedores (*Google*, *Microsoft*, etc) y centralizador de autenticación de los mismos. Su tarea principal será:

- Registrar, actualizar y eliminar usuarios en la base de datos, tanto usuarios locales como usuarios provenientes de otros proveedores de autenticidad.
- Validar usuarios y sesiones, proveyendo a los demás microservicios internos un canal de comunicación que permita consultar esa validez.

### Productos

### Pasarela de Pagos

### Backend-For-Frontend (Gateway)

Se encargará de ser la puerta entre el Cliente y la aplicación.
Conectará mediante eventos publicados en RabbitMQ con las aplicaciones internas cuando sea necesario.
Gestionará los metadatos de entrada y responderá según *User-Agent*, para proveer al cliente de:
- Metadatos SEO para agentes de búsqueda (*Googlebot*, *Bingbot*, etc).
- */llms.txt* para Modelos de Inteligencia Artifical.
- Datos JSON para frontend convencional.

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
