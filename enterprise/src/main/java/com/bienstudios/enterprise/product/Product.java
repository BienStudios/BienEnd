package com.bienstudios.enterprise.product;

import jakarta.persistence.*;

import java.math.BigDecimal;
import java.time.Instant;
import java.util.List;
import java.util.Map;

import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.JdbcTypeCode;
import org.hibernate.annotations.SQLRestriction;
import org.hibernate.annotations.UpdateTimestamp;
import org.hibernate.type.SqlTypes;

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
@SQLRestriction("application = CURRENT_SETTING('app.current_application', true)")
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

    // Aplicación a la que se atribuye
    @Column(name = "application", nullable = false, updatable = false)
    private String application;

    // Marca del producto
    @Column(name = "brand", nullable = false)
    private String brand;

    // Modelo genérico
    @Column(name ="model", nullable = false)
    private String model;

    // Descripción del producto
    @Column(name = "description", columnDefinition = "TEXT")
    private String description;

    @Column(name = "condition")
    @Enumerated(EnumType.STRING)
    private Condition condition;

    @Column(name = "health")
    @Enumerated(EnumType.STRING)
    private HealthStatus health;

    @Column(name = "legality")
    @Enumerated(EnumType.STRING)
    private LegalStatus legality;

    /*
    ----------------------------------------
    ---------------- Mapas -----------------
    ----------------------------------------
    */

    @JdbcTypeCode(SqlTypes.JSON)
    @Column(name = "categories", columnDefinition = "jsonb")
    private Map<Category, List<SubCategory>> categories;

    @JdbcTypeCode(SqlTypes.JSON)
    @Column(name = "payment_conditions", columnDefinition = "jsonb")
    private Map<PaymentCondition, List<PaymentMethod>> paymentConditions;

    /*
    ----------------------------------------
    ---------------- STAMP -----------------
    ----------------------------------------
    */

    @Column(name = "price", nullable = false)
    private BigDecimal price;

    @CreationTimestamp
    @Column(name = "created_at", updatable = false)
    private Instant createdAt;

    @UpdateTimestamp
    @Column(name = "last_update")
    private Instant lastUpdate;
}
