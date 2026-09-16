# Desarrollo políglota, microservicios escalables.

## Contenido:

- [Desarrollo políglota](#desarrollo-políglota)
- [Contenerización de los Microservicios](#contenerización-de-los-microservicios)

## Desarrollo políglota

Decisión impulsada por *"Workload Fit"*; Es conocido que cada ecosistema tiene marcada más una cierta orientación frente a la competencia, impulsada mayormente incluso por sus propias comunidades. Estudiando el tema, separamos diferentes lógicas basadas en las fortalezas de los siguientes ecosistemas:

- **Node.js + Express**, *Backend For Frontend*: El *Event Loop* optimiza operaciones asíncronas de entrada y salida (I/O) de datos, lo que permite que atender en casi simultaneidad a múltiples clientes sin congelar el hilo físico asignado. Para operaciones bloqueantes, se implementa **RabbitMQ** como servicio de mensajería interno.
- **Java + Spring Boot**, *Usuarios y Autenticación*: La velocidad del Java Moderno, la madurez de su entorno y continuo crecimiento de Spring Boot y su ecosistema perfecto para validaciones de seguridad, además de una comunidad excesivamente enorme; lo hacen casi perfecto para una de las capas más importantes de casi cualquier sistema empresarial. Esta capa será la que persista usuarios en la base de datos, y la encargada de la protección de sesiones.
- **C# + .NET**, *Productos*: Ya que los productos corresponden a datos que son generalmente más leídos que modificados; *.NET* es la respuesta para satisfacer correctamente la necesidad de *Query* sin tracking, y de *Commands* de modificación sólo cuando son necesarios; por lo que ofrece balancear la carga para el Servidor satisfaciendo la necesidad de mayor demanda de *Queries* por sobre los *Commands*, por lo que es posible optimizar la entrega de datos, y trabajar los mismos únicamente cuando deben validarse para modificaciones en la Base de Datos.
- **Go + Gin**, *Pasarela de Pagos*: Siendo Gin un Wrapper de *net/http* con arquitectura *Radix Tree*, la cual no ocupa alojamientos en memoria para las solicitudes, lo convierten en un proceso ligero y optimizado para llevar a cabo las transacciones de movimientos bancarios, y realizar el registro de las mismas en una Base de Datos no relacional (para escalado horizontal; *MongoDB* en este caso).
- **Go + Gin**, *Carrito de compras*: Apoyado por *Valkey*, al no necesitar una lógica compleja ni de escritura ni de lectura, y al no necesitar persistir a lo largo de tiempo los datos, se optó por una base de datos en caché ligera, orquestada como *goroutines* para evitar carga innecesaria de memoria.
- **Python + FastAPI**, *ChatBot*: Al ser el ecosistema más creciente y maduro en librerías de Inteligencia Artificial (por su simpleza), es el rey indiscutido para esta área.
- **Python + FastAPI**, *Mejora de imágenes/fotografías*: Nuevamente gana por su ecosistema; al ser un servicio no tan solicitado dentro del servidor, y estár comunicado por medio de *RabbitMQ* con el entorno para no bloquear una petición, no necesita mayores optimizaciones o tiempos de respuesta ultra-rápidos. La necesidad principal, abarca la simpleza en el código para adaptarlo a las necesidades de cada desarrollador/servicio.

## Contenerización de los Microservicios

La imagen es unificada y contiene los entornos necesarios para el proyecto completo. Los contenedores están pensados para múltiples despliegues en un mismo servidor, y se reparten en:

- **Contenedor compartido**: Está pensado para las Base de Datos, los servicios de mensajería y el microservicio de imágenes; y cualquier otro servicio que sea accedido de manera *pública* o sin modificaciones importantes por cada aplicación por sí misma.
- **Contenedor único por aplicación**: Está pensado para los microservicios propios de cada aplicación backend. Contendrá un pequeño módulo *"BUS"* para conectar el *frontend* con el *backend for frontend*.
