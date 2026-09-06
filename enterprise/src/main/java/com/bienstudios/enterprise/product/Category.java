package com.bienstudios.enterprise.product;

/**
 * <h3 lang="es">Categoría</h3>
 * 
 * <p lang="es">{@code Enum} ordinal que representa las
 * categorías de los productos o servicios comerciales
 * de la empresa.</p>
 * <p lang="es">Puede modificarse a las necesidades
 * específicas (<i>Previo a guardar elementos en la
 * base de datos</i>).</p>
 * <p lang="es">Puede ser alterada a un {@code Enum} de
 * <i>cadenas de texto</i>, añadiéndole la traducción a
 * cada elemento y modificando la <strong>entidad {@code
 * Product}</strong> en los campos <i>{@code
 * @Enumerated(EnumType.ORDINAL)}</i> cambiándolos por
 * <i>{@code @Enumerated(EnumType.STRING)}. Este cambio
 * permitirá simplificar cualquier modificación futura
 * del enum incluyendo el orden de los elementos.</i></p>
 * 
 * @author Brian Merino - BienStudios Develops
 * @since 26.00.0000_ALPHA
 */
public enum Category {
    ELECTRONICS,
    FASHION,
    HOME,
    BEAUTY,
    SPORTS,
    TOYS,
    AUTOMOTIVE,
    BOOKS,
    MUSIC,
    GROCERY
}
