plugins {
	java
	id("org.springframework.boot") version "4.1.1"
	id("io.spring.dependency-management") version "1.1.7"
}

group = "com.bienstudios"
version = "26.00.0000-SNAPSHOT"

java {
	toolchain {
		languageVersion = JavaLanguageVersion.of(25)
	}
}

repositories {
	mavenCentral()
}

dependencies {

	// Dependencias de Spring Boot
	implementation("org.springframework.boot:spring-boot-starter-webmvc")
	implementation("org.springframework.boot:spring-boot-starter-data-jpa")

	// Lombok para reducir el código boilerplate
	compileOnly("org.projectlombok:lombok")
	annotationProcessor("org.projectlombok:lombok")

	// Para manejo de flujos OAuth2 y OpenID Connect
	implementation("org.springframework.boot:spring-boot-starter-oauth2-client")

	// Manejar seguridad de APIs con JWT
	implementation("org.springframework.boot:spring-boot-starter-oauth2-resource-server")

	// DevTools para reinicio automático y otras utilidades de desarrollo
	developmentOnly("org.springframework.boot:spring-boot-devtools")

	// Driver PostgreSQL
	runtimeOnly("org.postgresql:postgresql")

	// Test
	testImplementation("org.springframework.boot:spring-boot-starter-webmvc-test")
	testCompileOnly("org.projectlombok:lombok")
	testRuntimeOnly("org.junit.platform:junit-platform-launcher")
	testAnnotationProcessor("org.projectlombok:lombok")
}

tasks.withType<Test> {
	useJUnitPlatform()
}
