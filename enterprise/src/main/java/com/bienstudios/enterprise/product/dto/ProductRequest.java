package com.bienstudios.enterprise.product.dto;

import com.bienstudios.enterprise.product.Category;
import com.bienstudios.enterprise.product.Condition;
import com.bienstudios.enterprise.product.HealthStatus;
import com.bienstudios.enterprise.product.LegalStatus;
import com.bienstudios.enterprise.product.PaymentCondition;
import com.bienstudios.enterprise.product.PaymentMethod;
import com.bienstudios.enterprise.product.SubCategory;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Positive;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

/**
 * <h3 lang="es">Solicitud de {@code Product} (<i>entrada
 * desde el cliente</i>)</h3>
 * 
 * <p lang="es">Representación de un producto el
 * cual se solicita crear dentro del sistema.</p>
 * 
 * <p lang="es">Esta clase se utiliza para recibir
 * información desde el cliente y validar el esqueleto
 * de un producto para su posterior creación</p>
 * 
 * @param brand La marca del producto.
 * @param model El modelo del producto.
 * @param categories Las categorías y subcategorías del producto.
 * @param description La descripción del producto.
 * @param condition El estado del producto (Nuevo, Usado, Reacondicionado).
 * @param health El estado de salud del producto (En condiciones, Deteriorado,
 * Destruído, etc).
 * @param legality La legalidad del producto (Papeles, Trámites legales, etc).
 * @param paymentConditions Las condiciones de pago del producto (Descuento,
 * Cuotas sin interés, etc).
 * @param price El precio del producto.
 * 
 * @author Brian Merino - BienStudios Develops
 * @since 26.00.0000_ALPHA
 */
public record ProductRequest(

    // El origen (aplicación a la que pertenece) se inyectará
    // automáticamente en el Servicio.

    @NotBlank(message = "La marca del producto no puede quedar vacío.")
    String brand,

    @NotBlank(message = "El modelo del producto no puede quedar vacío.")
    String model,

    @NotEmpty(
        message = "El producto debe contener al menos una "
        + "categoría con sus subcategorías.")
    Map<Category, List<SubCategory>> categories,

    @NotBlank(
        message = "La descripción es necesaria para explicar el "
        + "producto y sus características, así como para proporcionar "
        + "información adicional a los usuarios."
    )
    String description,

    // Nuevo, Usado, Reacondicionado
    @NotNull(message = "Es necesario indicar el estado del producto.")
    Condition condition,

    // En condiciones, Deteriorado, Destruído, etc
    @NotNull(
        message = "Se requiere informar sobre la integridad física "
        + "del producto."
    )
    HealthStatus health,

    // Papeles (automotor, inmobiliario), o trámites legales (certificados,
    // permisos, licencias, etc). Opcional, pero recomendable para productos
    // que requieran legalidad.
    LegalStatus legality,

    // Condición que aplique sobre el pago de un producto, atado a su forma de
    // pago. Por ejemplo, <Descuento, Efectivo>, <Cuotas_sin_interés,
    // Crédito_tarjeta_Visa>. Opcional para precios variables según la oferta
    // de mercado, pero recomendable para productos que requieran condiciones
    // de pago específicas.
    Map<PaymentCondition, List<PaymentMethod>> paymentConditions,

    // Precio del producto. Limitar a dos decimales en la capa de servicio, para
    // evitar problemas de redondeo y precisión en la capa de persistencia.
    @NotNull(message = "El precio es obligatorio.")
    @Positive(message = "El precio debe ser un valor positivo")
    BigDecimal price
) {}
