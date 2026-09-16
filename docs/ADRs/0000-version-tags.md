# Prototipado de versiones:

La separación de etiquetas de versión para el repositorio, seguirá la convención:

- `YY`: Dos dígitos representando el año del inicio del desarrollo de la versión.
- `VV`: Dos dígitos indicando la versión **dentro del año** del inicio del desarrollo/actualización. Se inicia por el `00`.
- `PPPP`: Cuatro dígitos indicando el número de parche sobre la versión desarrollada. Se inicia por el `0000`.
- `LIFECYCLE`: Ciclo de vida de la aplicación, con convención *"UPPERCASE"*; los cuales serán:

    - `ALPHA`: Primer lanzamiento de una versión, aún no aprobado para publicación.
    - `BETA`: Pruebas con público limitado. Se esperan reportes de problemas no arbitrados en pruebas internas.
    - `SNAPSHOT`: Versión pública de mayor alcance. Se esperan reportes mayores, con aserciones referentes a rendimiento, escalabilidad, modificabilidad, y facilidad para el desarrollador para no cometer errores.
    - `RELEASE`: Versión pública lista para producción. Se esperan reportes de mejoras de rendimiento, implementabilidad, esquemáticas, o de cualquier tipo que contribuyan a mejorar más el sistema.

La separación entre `YY`, `VV` Y `PPPP` será por "punto" ("."), mientras que el `PPPP` se separará del ciclo de vida por un "guión bajo" ("_"); siendo un ejemplo de versión:

- **Versión**: *26.00.0000_ALPHA*.

La cual indica:

- El año es que se comenzó a desarrollar la versión es 2026.
- Es la primera versión del año.
- Es el primer parche de la versión (aún no hay correcciones).
- Está en el ciclo *"ALPHA"* de producción, por lo que aún no se considera apta para uso público; si no, más bien, para pruebas privadas dentro del equipo de desarrollo.
