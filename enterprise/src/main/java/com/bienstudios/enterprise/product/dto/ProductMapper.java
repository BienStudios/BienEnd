package com.bienstudios.enterprise.product.dto;

import com.bienstudios.enterprise.product.Product;

import org.mapstruct.Mapper;

/**
 * <h3 lang="es">Mapeador de <i>entidad</i> {@code Product}</h3>
 * 
 * <p lang="es">Mapeador de {@code ProductRequest} a {@code Product}
 * (<i>entity</i>) para la entrada de datos desde el cliente, y de
 * {@code Product} a {@code ProductResponse} para enviar el producto
 * en formato de respuesta de regreso al cliente.</p>
 * 
 * @author Brian Merino - BienStudios Develops
 * @since 26.00.0000_ALPHA
 */
@Mapper
public interface ProductMapper {

    // Recibe el producto como "ProductRequest" y lo convierte
    // en una entidad para obtener datos precisos y ordenados
    // del producto
    Product toEntity(ProductRequest request);

    // Transforma la entidad interna "Product" a una entidad
    // que el cliente puede recibir e interpretar, "ProductResponse"
    ProductResponse toResponse(Product entity);
}
