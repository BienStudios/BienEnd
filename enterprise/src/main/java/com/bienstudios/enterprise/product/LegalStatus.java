package com.bienstudios.enterprise.product;

/**
 * <h3 lang="es">Estado Legal</h3>
 * 
 * <p lang="es">{@code Enum} ordinal que representa
 * el estado de legalidad del producto.</p>
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
public enum LegalStatus {
    // Paperwork and legal documentation for the product is in
    // order and compliant with regulations.
    LEGAL,

    // Paperwork and legal documentation for the product is not
    // in order and may not be compliant with regulations.
    ILLEGAL,

    // Paperwork and legal documentation for the product is
    // restricted and may require special permissions or licenses
    // to be sold or used.
    RESTRICTED
}
