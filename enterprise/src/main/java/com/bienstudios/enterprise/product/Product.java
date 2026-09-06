package com.bienstudios.enterprise.product;

import java.math.BigDecimal;
import java.time.Instant;

import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.UpdateTimestamp;

import jakarta.persistence.*;

import lombok.*;

/**
 * <h3 lang="es">Producto (<i>entidad</i>)</h3>
 *
 * <p lang="es">Entidad de datos para operaciones internas
 * y en común con la base de datos.</p>
 * 
 * @author BienStudios Develops
 * @since 26.00.0000_ALPHA
 */
@Getter
@Setter
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Entity
@Table(name = "products")
public class Product {

    /**
     * <h4 lang="es">Identificador único del producto</h4>
     * 
     * <p lang="es">Este campo representa el identificador único del producto
     * guardado al momento de registrarlo en la base de datos. Es un valor
     * generado automáticamente por la misma. El nombre común en la tabla de
     * la base de datos es normalmente {@code id}</p>
     */
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    // Marca del producto
    @Column(nullable = false)
    private String brand;

    // Modelo genérico
    @Column(nullable = false)
    private String model;

    // Descripción del producto
    @Column(columnDefinition = "TEXT")
    private String description;

    // Condición. Cambie a EnumType.STRING si así lo desea para que
    // sea más legible en la base de datos y más fácil la posterior
    // modificación de los enums Condition, HealthStatus, y LegalStatus
    @Enumerated(EnumType.ORDINAL)
    private Condition condition;

    @Enumerated(EnumType.ORDINAL)
    private HealthStatus health;

    @Enumerated(EnumType.ORDINAL)
    private LegalStatus legality;

    // Sin mapeo en entidad por temas de simplicidad.
    // De ser necesario, añadir con el mapeo correspondiente ó simplificar
    // como un string con estructura "JSON" si así lo desea.
    // Map<PaymentCondition, List<PaymentMethod>> paymentConditions;

    @Column(nullable = false)
    private BigDecimal price;

    @CreationTimestamp
    @Column(name = "created_at", updatable = false)
    private Instant createdAt;

    @UpdateTimestamp
    @Column(name = "last_update")
    private Instant lastUpdate;
}
