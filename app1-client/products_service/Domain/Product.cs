namespace ProductsService.Domain;

/// <summary>
/// Representa un Producto dentro del inventario del e-commerce.
/// </summary>
public class Product
{
    // Guid es equivalente a un UUID de 128 bits.
    public Guid Id { get; private set; }
    
    public string Name { get; private set; } = string.Empty;
    public string Description { get; private set; } = string.Empty;
    
    // decimal es el tipo numérico de precisión exacta ideal para dinero/precios.
    public decimal Price { get; private set; }
    public int Stock { get; private set; }

    // Constructor privado para obligar a usar el método de fábrica (Encapsulamiento).
    private Product() { }

    public Product(string name, string description, decimal price, int stock)
    {
        if (string.IsNullOrWhiteSpace(name))
            throw new ArgumentException("El nombre del producto no puede estar vacío.", nameof(name));
        
        if (price <= 0)
            throw new ArgumentOutOfRangeException(nameof(price), "El precio debe ser mayor a cero.");

        Id = Guid.NewGuid();
        Name = name;
        Description = description;
        Price = price;
        Stock = stock;
    }

    /// <summary>
    /// Regla de negocio: Descuenta el stock ante una compra confirmada.
    /// </summary>
    public void DecreaseStock(int quantity)
    {
        if (quantity > Stock)
            throw new InvalidOperationException($"Stock insuficiente para el producto {Name}.");

        Stock -= quantity;
    }
}
