using System.Text;
using System.Text.Json;
using RabbitMQ.Client;

namespace ProductsService.Infrastructure.Messaging;

public interface IRabbitMqPublisher
{
    Task PublishProductCreatedEventAsync<T>(T message, string routingKey);
}

/// <summary>
/// Servicio encargado de publicar eventos asíncronos hacia RabbitMQ.
/// </summary>
public class RabbitMqPublisher : IRabbitMqPublisher
{
    private readonly IConfiguration _configuration;

    public RabbitMqPublisher(IConfiguration configuration)
    {
        _configuration = configuration;
    }

    public async Task PublishProductCreatedEventAsync<T>(T message, string routingKey)
    {
        var factory = new ConnectionFactory
        {
            HostName = _configuration["RabbitMQ:Host"] ?? "localhost",
            Port = int.Parse(_configuration["RabbitMQ:Port"] ?? "5672")
        };

        // En C# 8+, 'await using' destruye y libera el recurso/conexión al finalizar el bloque.
        await using var connection = await factory.CreateConnectionAsync();
        await using var channel = await connection.CreateChannelAsync();

        // Declaramos el Topic Exchange del dominio
        await channel.ExchangeDeclareAsync(
            exchange: "product.events", 
            type: ExchangeType.Topic, 
            durable: true
        );

        // Serializamos el payload a JSON binario (UTF-8)
        var json = JsonSerializer.Serialize(message);
        var body = Encoding.UTF8.GetBytes(json);

        // Publicación del mensaje
        await channel.BasicPublishAsync(
            exchange: "product.events",
            routingKey: routingKey,
            body: body
        );
    }
}
