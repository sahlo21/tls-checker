# TLS Checker CLI 🛡️

Una herramienta de línea de comandos (CLI) desarrollada en **Go** que permite analizar la seguridad y configuración de los certificados TLS de cualquier dominio utilizando la API pública de **SSL Labs**.

## 📋 Descripción
Este proyecto permite consultar el estado de los certificados SSL/TLS de forma remota, devolviendo información detallada sobre la calificación de seguridad del servidor. Ha sido diseñado bajo un enfoque modular para asegurar que el código sea fácil de leer, mantener y testear.

## 🚀 Instalación y Uso

### Requisitos previos
* **Go** 1.21 o superior instalado localmente.
* Conexión a internet para realizar las peticiones a la API externa.

### Configuración
1.  **Clonar el repositorio:**
    ```bash
    git clone [https://github.com/sahlo21/tls-checker.git](https://github.com/sahlo21/tls-checker.git)
    cd tls-checker
    ```

### Ejecución
Para ejecutar la herramienta directamente sin necesidad de generar un binario:
```bash
go run . -host google.com
Compilación (Forma óptima)
Si prefieres generar un archivo ejecutable para usarlo en cualquier lugar:

Bash
go build -o tls-checker
./tls-checker -host google.com
🏗️ Arquitectura del Proyecto
El código se ha dividido en múltiples archivos para separar las responsabilidades (Separation of Concerns):

main.go: Orquestador principal del programa. Maneja el flujo de ejecución y la interacción inicial.

client.go: Contiene la lógica necesaria para realizar las peticiones HTTP a la API de SSL Labs y manejar los tiempos de espera.

model.go: Define las estructuras de datos (structs) que representan la respuesta JSON de la API.

utils.go: Incluye funciones de soporte para el procesamiento de argumentos de la terminal y el formateo de los resultados impresos.

go.mod: Define el módulo del proyecto y gestiona las versiones de Go.

🛠️ Tecnologías utilizadas
Go (Golang): Lenguaje de programación principal.

Git: Control de versiones.

SSL Labs API: Fuente de datos para el análisis de TLS.

Desarrollado por sahlo21
