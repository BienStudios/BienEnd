namespace ProductsService.Features.Products;

public static class ProductEndpoints
{
    /// <summary>
    /// Registra los endpoints de Minimal API para el módulo de Productos.
    /// </summary>
    public static IEndpointRouteBuilder MapProductEndpoints(this IEndpointRouteBuilder endpoints)
    {
        var group = endpoints.MapGroup("/api/products")
                             .WithTags("Products");

        // Asignamos la ruta directamente al método estático HandleAsync de cada feature
        group.MapPost("/", CreateProduct.HandleAsync)
             .WithName("CreateProduct")
             .WithSummary("Crea un nuevo producto y emite el evento a RabbitMQ.");

        group.MapGet("/{id:guid}", GetProductById.HandleAsync)
             .WithName("GetProductById")
             .WithSummary("Obtiene un producto por su UUID en modo lectura optimizada.");

        return endpoints;
    }
}
