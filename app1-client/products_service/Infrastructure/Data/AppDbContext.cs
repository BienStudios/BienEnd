using Microsoft.EntityFrameworkCore;
using ProductsService.Domain;

namespace ProductsService.Infrastructure.Data;

/// <summary>
/// Contexto de Entity Framework Core para la gestión de la base de datos (PostgreSQL/SQL Server/In-Memory).
/// </summary>
public class AppDbContext : DbContext
{
    public AppDbContext(DbContextOptions<AppDbContext> options) : base(options) { }

    // DbSet representa la tabla 'Products' en la base de datos.
    public DbSet<Product> Products => Set<Product>();

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);

        // Mapeo fluido (Fluent API) para definir restricciones de tabla.
        modelBuilder.Entity<Product>(entity =>
        {
            entity.HasKey(p => p.Id);
            entity.Property(p => p.Name).HasMaxLength(150).IsRequired();
            entity.Property(p => p.Price).HasPrecision(18, 2); // 18 dígitos, 2 decimales
        });
    }
}
