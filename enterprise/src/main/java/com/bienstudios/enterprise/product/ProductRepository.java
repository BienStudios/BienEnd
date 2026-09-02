package com.bienstudios.enterprise.product;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * <h3 lang="es">Repositorio de Productos</h3>
 *
 * <p lang="es">Interfaz que define las operaciones CRUD para la entidad {@link Product}.</p>
 * 
 * @author BienStudios Develops
 * @since 26.00.0000_SNAPSHOT
 */
@Repository
public interface ProductRepository extends JpaRepository<Product, Long> {
    // Custom query methods can be defined here if needed
}
