package com.bienstudios.enterprise.product.dto;

import java.math.BigDecimal;
import java.time.Instant;

import lombok.Builder;

/**
 * <h3 lang="es">{@code Product} de respuesta (<i>Devolución
 * al cliente</i>)</h3>
 * 
 * <p lang="es">Respuesta del servidor hacia el cliente
 * en formato interpretable y legible para el mismo, con
 * los datos que el mismo pueda requerir.</p>
 * 
 * @param id
 * <p lang="es">Identificador único en la base de datos.</p>
 * <p lang="en">Unique identifier in the database.</p>
 * 
 * @param brand
 * <p lang="es">Marca del producto.</p>
 * <p lang="en">Product brand.</p>
 * 
 * @param model
 * <p lang="es">Modelo genérico del producto.</p>
 * 
 * @param description
 * <p lang="es">Descripción del producto.</p>
 * 
 * @param condition
 * <p lang="es">Condición del producto (Nuevo, Usado, Restaurado...).</p>
 * 
 * @param health
 * <p lang="es">Integridad física del producto</p>
 * 
 * @param legality
 * <p lang="es">Estado legal del producto (Papeles al día, Copia Original, etc.).</p>
 * 
 * @param paymentConditions
 * <p lang="es">Condiciones de Pago (Ofertas, descuentos, beneficios, etc)</p>
 * 
 * @param price
 * <p lang="es">Precio del producto (con 2 decimales)</p>
 * 
 * @param createdAt
 * <p lang="es">{@code TIMESTAMP} de la creación del producto en la base de datos
 * (no modificable).</p>
 * 
 * @param lastUpdate
 * <p lang="es">{@code TIMESTAMP} de la última modificación del producto
 * (modificable).</p>
 * 
 * @author Brian Merino - BienStudios Develops
 * @since 26.00.0000_ALPHA
 */
@Builder
public record ProductResponse(
    Long id,
    String brand,
    String model,
    String description,
    String condition,
    String health,
    String legality,
    String paymentConditions,
    BigDecimal price,
    Instant createdAt,
    Instant lastUpdate // TIMESTAMP WITH TIME ZONE (PostgreSQL)
) {}
