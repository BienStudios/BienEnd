-- Creación de las bases de datos aisladas.
-- TODO: renombrar de acuerdo a las necesidades de desarrollo
CREATE DATABASE enterprise_db;
CREATE DATABASE enterprise_db_test;

-- Creación de Usuarios SQL
CREATE USER app_user WITH ENCRYPTED PASSWORD "change_me_on_first_login";
CREATE USER test_user WITH ENCRYPTED PASSWORD "change_me_too";

-- Asignación de privilegios
GRANT ALL PRIVILEGES ON DATABASE enterprice_db TO app_user;
GRANT ALL PRIVILEGES ON DATABASE enterprice_db_test TO test_user;

-- Conexión e instalación de extensiones
\c enterprise_db;
CREATE EXTENSION IF NOT EXISTS "uuid_ossp";
-- otras extensiones para producción

\c enterprise_db_test;
CREATE EXTENSION IF NOT EXISTS "uuid_ossp";
-- otras extensiones para testing
