<!--
################################################################
--------------------- REFERENCIAS EXTERNAS ---------------------
################################################################
-->
[appc_users_rm]: ./../users_service/README.es.md

<!-- Fin de Referencias Externas -->

# BienEnd - Servicio de pasarela de pagos

## Contenido

- [¿Por qué Go y Gin?](#por-qué-golang-y-gin)

## ¿Por qué Golang y Gin?

Ya que es un servicio el cual no necesita un consumo excesivo de memoria; al ser realizadas las primeras validaciones esenciales en el [Servicio de Usuarios y Autenticación][appc_users_rm] y necesitar revalidarlas de manera simplificada, y al contar con una memoria de pila ligera, Go es una de las mejores opciones que consideramos; y gracias a la arquitectura de Gin, se puede trabajar el inicio y final del servicio de manera sencilla, modificando y adaptando únicamente lo que sucede en el medio (*middleware*), y el árbol radix de rutas, le dan una gran ventaja de velocidad.
