package com.bienstudios.enterprise.product;

import jakarta.persistence.*;

import lombok.*;

/**
 * <h3 lang="es">Producto</h3>
 *
 * <p lang="es">Utilice esta entidad como representación de un producto
 * en la base de datos. Las propiedades aquí definidas son moldeables a
 * las necesidades específicas de cada caso de uso.</p>
 * 
 * @author BienStudios Develops
 * @since 26.00.0000_SNAPSHOT
 */
@Getter
@Setter
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Entity
@Table(
    name = "products",
    indexes = {
        @Index(name = "idx_products_name", columnList = "name")
    },
    uniqueConstraints = {
        @UniqueConstraint(name = "uk_products_name", columnNames = "name")
    }
)
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

}
