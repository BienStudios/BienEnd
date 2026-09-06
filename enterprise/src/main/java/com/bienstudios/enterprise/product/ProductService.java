package com.bienstudios.enterprise.product;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

/**
 * <h3 lang="es">Capa de Servicio de {@code Product}</h3>
 * 
 * <p lang="es">Paginación y limpieza de la <i>entidad</i>.
 * Simplificación y modelo "base", modifique según necesidades
 * específicas.</p>
 * 
 * @author Brian Merino - BienStudios Develops
 * @since 26.00.0000_ALPHA
 */
public interface ProductService {

    Page<Product> getAll(Pageable pageable);
    Product create(Product product);
}
