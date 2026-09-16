plugins {
	java
	id("org.springframework.boot") version "4.1.1"
	id("io.spring.dependency-management") version "1.1.7"
}

group = "com.bienstudios"
version = "26.00.0000_ALPHA"

java {
	toolchain {
		languageVersion = JavaLanguageVersion.of(25)
	}
}

repositories {
	mavenCentral()
}

// Centralización de versiones
val mapstructVersion = "1.6.3"
val lombokVersion = "1.18.36"
val lombokMapstructBindingVersion = "0.2.0"
val nimbusJoseJwtVersion = "10.0.1"

dependencies {
	// Web y validación de inputs
	implementation("org.springframework.boot:spring-boot-starter-webmvc")

	// Validación e DTOs
	implementation("org.springframework.boot:spring-boot-starter-validation")

	// Seguridad, Filtros y JWT/JWKS
	implementation("org.springframework.boot:spring-boot-starter-security")

	// Nimbus JOSE + JWT para firmar tokens y convertir claves públicas a JWKS
	implementation("com.nimbusds:nimbus-jose-jwt:$nimbusJoseJwtVersion")

	// Persistencia (Jakarta JPA / Hibernate / PostgreSQL)
	implementation("org.springframework.boot:spring-boot-starter-data-jpa")
	runtimeOnly("org.postgresql:postgresql")

	// Mapping DTO por MapsStruct, evitar Reflection para mejorar rendimiento
	implementation("org.mapstruct:mapstruct:$mapstructVersion")
	annotationProcessor("org.mapstruct:mapstruct-processor:$mapstructVersion")

	// Lombok y Binding con MapStruct
	compileOnly("org.projectlombok:lombok:$lombokVersion")
	annotationProcessor("org.projectlombok:lombok:$lombokVersion")
	// Binding
	annotationProcessor("org.projectlombok:lombok-mapstruct-binding:$lombokMapstructBindingVersion")

	// Monitoreo y escalabilidad
	implementation("org.springframework.boot:spring-boot-starter-actuator")

	// test simplificado
	testImplementation("org.springframework.boot:spring-boot-starter-test")
	testImplementation("org.springframework.security:spring-security-test")
}

tasks.withType<Test> {
	useJUnitPlatform()
}
