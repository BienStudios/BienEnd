# Python 3.14 base
FROM python:3.14-rc-slim

WORKDIR /app

# Dependencias del sistema necesarias para OpenCV o Pillow
RUN apt-get update && apt-get install -y --no-install-recommends \
    libgl1 \
    libglib2.0-0 \
    && rm -rf /var/lib/apt/lists/*

# Entorno virtual de Python (guardado en el Path)
RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

# Copia de las librerías dentro del venv
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 8000

#    Servidor ASGI Python | Aplicación Python | Servidor interno |   Puerto
CMD ["uvicorn",             "main:app",         "--host", "0.0.0.0", "--port", "8000"]
