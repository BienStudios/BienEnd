# BienEnd - BienStudios Backend, backend basado en microservicios

Léelo en otros idiomas:

- [English](README.md)

## Contenido

- [Sistema de Versiones](#sistema-de-versiones)
- [Propósito](#propósito)

## Sistema de Versiones

Las versiones serán desplegadas tal que: `AA.LL.PPPP_C...`; donde:

- **AA**: Año del lanzamiento de la versión en 2 dígitos; por ejemplo: 2026 -> `26`.
- **LL**: Lanzamiento ordenado (siendo el primer índice el *0*) dentro del año de la versión; por ejemplo: Primera versión dentro del 2026: `00`. La segunda versión del año sería `01`.
- **PPPP**: Parche ordenado del lanzamiento. Cambios mínimos para corregir errores, añadir documentación o similares. Por ejemplo: Primer lanzamiento: `0000`. Se añade documentación al código: `0001`. Se parcha un error de validación: `0002`.
- **C...**: Etapa en el ciclo de vida:
  - **ALPHA**: Versión en testing; no aprobada aún para uso personal.
  - **BETA**: Versión en testing a escala: acceso a un grupo de usuarios, que reporten errores, problemas o posibles mejoras.
  - **SNAPSHOT**: Versión al público, "pulida" (testing *Beta* aprobado). Puede todavía contener errores o problemas aun no reportados ni corregidos.
  - **RELEASE**: Lanzamiento oficial al público. Complétamente aprobado para el uso público. Para mejoras, se trabajará en un nuevo lanzamiento.

***Ejemplo completo***:
- Primer lanzamiento del año 2026, aún sin parches y en etapa de pruebas internas:
  - **Versión**: `26.00.0000_ALPHA`.
- Cuarto parche desde el lanzamiento `ALPHA`, aún no aprobado para `BETA`:
  - **Versión**: `26.00.0003_ALPHA`.
- Se logra superar las pruebas internas y pasa a estudio con un grupo miniritario de gente:
  - **Versión**: `26.00.0000_BETA`.
  *Nótese que volvemos a `0000` el Parche al pasar a un nuevo "ciclo de vida"*.
- Luego de pasar pruebas `BETA`, lanzarse el `SNAPSHOT` por un tiempo prudente para recibir notificaciones de errores, publicarse el primer `RELEASE` al estar todo bien, el equipo nota que se pueden hacer mejoras de rendimiento en uno de los ecosistemas o microservicio, mientras aún estamos en *2026*:
  - **Versión**: `26.01.0000_ALPHA`.

Se considerará como tiempo prudente de prueba de `SNAPSHOT` a 60 días desde lanzada la versión. Si los problemas son menores, como falta de documentación, ortografía en algún *Log*, o similar, se pasarán a revisar y corregir en un siguiente lanzamiento.

## Propósito

Funcionar como una plantilla de fácil adaptacion a proyectos web para comercios electrónicos. Incluirá integración completa a microservicios, con un BFF (Backend For Frontend) el cual limitará a qué servicios estará limitada la aplicación según necesidades.
Se eligió que cada microservicio se ambiente en su mejor ecosistema:

- **BFF**: *Node.js* + *Express*. Servidor con lenguaje unificado con el frontend. El Bucle de Eventos asícrono gestiona bien múltiples solicitudes concurrentes.
- **Enterprise**: *Java* + *Spring Boot* + *PostgreSQL*. Ofrece un entorno maduro y es el estándar de la industria. Es ideal para la capa de negocio, que trabaja con los datos persistentes y las validaciones. Facilita el escalado vertical de la aplicación.
- **Transacciones, crecimiento horizontal**: *Go* + *Gin*. Compilación a binario, genera ejecutables livianos y con poco consumo de RAM. Ideal para procesos en segundo plano o como puente con APIs externas sin interrumpir los demás procesos.
- **Edición de imágenes**: *Python* + *FastAPI* + *rembg* + *Pillow (PIL)*. Eliminar fondo de las fotografías utilizadas para que combinen con el fondo de la página, realzar colores y mejorar iluminación. Python ofrece herramientas muy avanzadas, y al no ser el encargado de todo el resto de tareas, no bloquea los hilos principales.
- **Chatbot**: *En estudio*. Aún no se tiene estudiado *qué* ni *cómo*. Durante el avance, se irá actualizando el contexto.
