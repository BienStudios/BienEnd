package com.bienstudios.enterprise.product;

/**
 * <h3 lang="es">Método de Pago</h3>
 * 
 * <p lang="es">{@code Enum} ordinal adjuntable a
 * {@code PaymentCondition}, para anexar una condición
 * para la adquisición del producto mediante uno o más
 * métodos de pago. Formas de pago pueden ser enlistadas
 * en otro {@code Enum} para mostrar las formas de pago
 * aceptadas para el producto individual, pero generalmente
 * las formas de pago son <strong>generales</strong> a
 * nivel comercial.</p>
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
public enum PaymentMethod {
    CASH,
    CREDIT_CARD,
    DEBIT_CARD,
    PAYPAL,
    BANK_TRANSFER,
    CRYPTOCURRENCY
}
