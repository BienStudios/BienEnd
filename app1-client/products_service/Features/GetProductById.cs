using Microsoft.EntityFrameworkCore;
using ProductsService.Infrastructure.Data;

namespace ProductsService.Features.Products;

public static class GetProductById
{
    // DTO exclusivo para la consulta
    public record Response(Guid Id, string Name, string Description, decimal Price, int Stock);

    /// <summary>
    /// Manejador de la consulta de productos por ID (Lectura Optimizada).
    /// </summary>
    public static async Task<IResult> HandleAsync(
        Guid id, 
        AppDbContext dbContext)
    {
        // Uso de AsNoTracking() para maximizar velocidad y reducir uso de memoria RAM
        var product = await dbContext.Products
            .AsNoTracking()
            .FirstOrDefaultAsync(p => p.Id == id);

        if (product is null)
        {
            return Results.NotFound(new { Message = $"El producto con ID {id} no existe." });
        }

        var response = new Response(product.Id, product.Name, product.Description, product.Price, product.Stock);
        return Results.Ok(response);
    }
}
