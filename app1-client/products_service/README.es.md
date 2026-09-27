# BienEnd - Servicio de Productos

## Contenido:

- [¿Por qué C# y .NET?](#por-qué-c-y-net)

## ¿Por qué C# y .NET?

La desición fué tomada por simplificar el catálogo de productos delegándolo a un framework el cuál optimiza justo lo que necesita un catálogo de comercio: Se leerá más veces de las que se modificará. Al plantear un patrón de diseño que optimiza las lecturas (*Queries*) y permite separar la capa de modificaciones (*Command*), y realizar las respectivas validaciones en esta misma, se decició que este ecosistema cumple precisamente los requerimientos técnicos.
