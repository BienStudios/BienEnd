using Microsoft.EntityFrameworkCore;
using ProductsService.Features.Products;
using ProductsService.Infrastructure.Data;
using ProductsService.Infrastructure.Messaging;

var builder = WebApplication.CreateBuilder(args);

// 1. Registro de Servicios (Inyección de Dependencias)
builder.Services.AddDbContext<AppDbContext>(options =>
    options.UseInMemoryDatabase("ProductsDb"));

builder.Services.AddTransient<IRabbitMqPublisher, RabbitMqPublisher>();
builder.Services.AddHostedService<PaymentConsumer>(); // Escucha RabbitMQ en 2º plano

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// 2. Middlewares HTTP
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

// 3. Mapeo de Endpoints por Módulo
app.MapProductEndpoints();

// 4. Ejecución
await app.RunAsync();
