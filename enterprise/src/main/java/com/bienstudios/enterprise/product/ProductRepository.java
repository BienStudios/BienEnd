package com.bienstudios.enterprise.product;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * <h3 lang="es">Repositorio de {@code Product}</h3>
 *
 * <p lang="es">Intermediario entre la Base de datos y la Capa de Servicio.</p>
 * 
 * <p lang="es">Versión simplificada provista por JPA ({@code JpaRepository}).
 * puede extenderse según sea necesario y modificarse para consultas y comunicación
 * más robustas.</p>
 * 
 * @author BienStudios Develops
 * @since 26.00.0000_ALPHA
 */
@Repository
public interface ProductRepository extends JpaRepository<Product, Long> {
    // Métodos JPA incluidos. Añadir más de ser necesarios.
}
