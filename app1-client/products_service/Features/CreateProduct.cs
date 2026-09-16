using Microsoft.EntityFrameworkCore;
using ProductsService.Domain;
using ProductsService.Infrastructure.Data;
using ProductsService.Infrastructure.Messaging;

namespace ProductsService.Features.Products;

public static class CreateProduct
{
    // DTOs exclusivos para el comando de creación
    public record Request(string Name, string Description, decimal Price, int Stock);
    public record Response(Guid Id, string Name, string Description, decimal Price, int Stock);

    /// <summary>
    /// Manejador de la creación de un nuevo producto.
    /// </summary>
    public static async Task<IResult> HandleAsync(
        Request request, 
        AppDbContext dbContext, 
        IRabbitMqPublisher publisher)
    {
        // 1. Instanciar Entidad con validaciones de negocio del Dominio
        var product = new Product(request.Name, request.Description, request.Price, request.Stock);

        // 2. Persistencia en la Base de Datos
        dbContext.Products.Add(product);
        await dbContext.SaveChangesAsync();

        // 3. Notificación asíncrona a RabbitMQ
        var productCreatedEvent = new
        {
            EventId = Guid.NewGuid(),
            EventType = "product.created",
            Timestamp = DateTime.UtcNow,
            Data = new { ProductId = product.Id, product.Name, product.Price }
        };

        await publisher.PublishProductCreatedEventAsync(productCreatedEvent, "product.created");

        // 4. Respuesta HTTP 201 Created
        var response = new Response(product.Id, product.Name, product.Description, product.Price, product.Stock);
        return Results.Created($"/api/products/{product.Id}", response);
    }
}
