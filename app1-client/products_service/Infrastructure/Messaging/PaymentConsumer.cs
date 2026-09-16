using System.Text;
using System.Text.Json;
using Microsoft.EntityFrameworkCore;
using ProductsService.Infrastructure.Data;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;

namespace ProductsService.Infrastructure.Messaging;

/// <summary>
/// Worker asíncrono que escucha eventos 'payment.transaction.succeeded' provenientes de Go.
/// </summary>
public class PaymentConsumer : BackgroundService
{
    private readonly IServiceProvider _serviceProvider;
    private readonly IConfiguration _configuration;
    private readonly ILogger<PaymentConsumer> _logger;

    public PaymentConsumer(IServiceProvider serviceProvider, IConfiguration configuration, ILogger<PaymentConsumer> logger)
    {
        _serviceProvider = serviceProvider;
        _configuration = configuration;
        _logger = logger;
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
        var factory = new ConnectionFactory
        {
            HostName = _configuration["RabbitMQ:Host"] ?? "localhost"
        };

        var connection = await factory.CreateConnectionAsync(stoppingToken);
        var channel = await connection.CreateChannelAsync(stoppingToken);

        // 1. Declarar Exchange de entrada
        await channel.ExchangeDeclareAsync("payment.events", ExchangeType.Topic, durable: true, cancellationToken: stoppingToken);

        // 2. Declarar Cola y su Dead Letter Exchange (DLX) para resiliencia
        var queueName = "product-service.payment-succeeded.queue";
        await channel.QueueDeclareAsync(
            queue: queueName,
            durable: true,
            exclusive: false,
            autoDelete: false,
            arguments: new Dictionary<string, object?>
            {
                { "x-dead-letter-exchange", "dlx.events" }
            },
            cancellationToken: stoppingToken
        );

        // 3. Vincular cola con la routing key del evento de Go
        await channel.QueueBindAsync(queueName, "payment.events", "payment.transaction.succeeded", cancellationToken: stoppingToken);

        var consumer = new AsyncEventingBasicConsumer(channel);

        consumer.ReceivedAsync += async (model, ea) =>
        {
            try
            {
                var body = ea.Body.ToArray();
                var jsonMessage = Encoding.UTF8.GetString(body);
                _logger.LogInformation("Evento recibido desde Pagos (Go): {Message}", jsonMessage);

                // Parsear evento
                using var doc = JsonDocument.Parse(jsonMessage);
                var root = doc.RootElement;
                
                // Extraer datos del producto vendido
                if (root.TryGetProperty("data", out var dataElement) && 
                    dataElement.TryGetProperty("product_id", out var productIdElement))
                {
                    var productId = Guid.Parse(productIdElement.GetString()!);

                    // DbContext en .NET es Scoped, por lo que creamos un Scope manual en el BackgroundService
                    using var scope = _serviceProvider.CreateScope();
                    var dbContext = scope.ServiceProvider.GetRequiredService<AppDbContext>();

                    var product = await dbContext.Products.FirstOrDefaultAsync(p => p.Id == productId, stoppingToken);
                    if (product != null)
                    {
                        product.DecreaseStock(1);
                        await dbContext.SaveChangesAsync(stoppingToken);
                        _logger.LogInformation("Stock actualizado para el producto {ProductId}", productId);
                    }
                }

                // Confirmación manual (ACK)
                await channel.BasicAckAsync(ea.DeliveryTag, multiple: false, cancellationToken: stoppingToken);
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "Error al procesar el pago. Enviando a Dead Letter Queue.");
                // Rechazar mensaje sin re-encolar (se va automáticamente a DLQ)
                await channel.BasicNackAsync(ea.DeliveryTag, multiple: false, requeue: false, cancellationToken: stoppingToken);
            }
        };

        await channel.BasicConsumeAsync(queue: queueName, autoAck: false, consumer: consumer, cancellationToken: stoppingToken);
    }
}
