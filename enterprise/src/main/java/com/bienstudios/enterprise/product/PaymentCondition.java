package com.bienstudios.enterprise.product;

/**
 * <h3 lang="es">Condiciones de Pago</h3>
 * 
 * <p lang="es">{@code Enum} ordinal que representa las
 * condiciones de pago para un producto, adjuntables a uno
 * o  más métodos de pago para la misma.</p>
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
public enum PaymentCondition {
    DISCOUNT,
    DISCOUNT_ON_BULK_PURCHASE,
    DISCOUNT_ON_FIRST_PURCHASE,
    DISCOUNT_ON_SUBSCRIPTION,
    DISCOUNT_ON_REFERRAL,
    DISCOUNT_ON_WHOLESALE,
    DISCOUNT_ON_SEASONAL_PROMOTION,
    QUOTES_WITHOUT_INTEREST
}
