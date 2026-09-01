# BienEnd - BienStudios Backend, backend basado en microservicios

Léelo en otros idiomas:

- [English](README.md)

## Contenido

- [Propósito](#propósito)

## Propósito

Funcionar como una plantilla de fácil adaptacion a proyectos web para comercios electrónicos. Incluirá integración completa a microservicios, con un BFF (Backend For Frontend) el cual limitará a qué servicios estará limitada la aplicación según necesidades.
Se eligió que cada microservicio se ambiente en su mejor ecosistema:

- **BFF**: *Node.js* + *Express*. Servidor con lenguaje unificado con el frontend. El Bucle de Eventos asícrono gestiona bien múltiples solicitudes concurrentes.
- **Enterprise**: *Java* + *Spring Boot* + *PostgreSQL*. Ofrece un entorno maduro y es el estándar de la industria. Es ideal para la capa de negocio, que trabaja con los datos persistentes y las validaciones. Facilita el escalado vertical de la aplicación.
- **Transacciones, crecimiento horizontal**: *Go* + *Gin*. Compilación a binario, genera ejecutables livianos y con poco consumo de RAM. Ideal para procesos en segundo plano o como puente con APIs externas sin interrumpir los demás procesos.
- **Edición de imágenes**: *Python* + *FastAPI* + *rembg* + *Pillow (PIL)*. Eliminar fondo de las fotografías utilizadas para que combinen con el fondo de la página, realzar colores y mejorar iluminación. Python ofrece herramientas muy avanzadas, y al no ser el encargado de todo el resto de tareas, no bloquea los hilos principales.
- **Chatbot**: *En estudio*. Aún no se tiene estudiado *qué* ni *cómo*. Durante el avance, se irá actualizando el contexto.
