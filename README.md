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
```
Compilación (Forma óptima)
Si prefieres generar un archivo ejecutable para usarlo en cualquier lugar:
```bash
go build -o tls-checker
./tls-checker -host google.com
```
## 🏗️ Arquitectura del Proyecto

El código ha sido estructurado siguiendo principios de **diseño modular** para separar las responsabilidades y facilitar la escalabilidad. A continuación se detalla la función de cada componente:

### 📂 Estructura de Archivos

* **`main.go`**: Es el punto de entrada de la aplicación. Se encarga de orquestar el flujo general, llamar a los procesos de análisis y gestionar el ciclo de vida de la ejecución.
* **`client.go`**: Actúa como la capa de comunicación externa. Implementa la lógica para realizar peticiones HTTP a la API de SSL Labs, gestionando los reintentos y los tiempos de espera (*timeouts*).
* **`model.go`**: Contiene la definición de las estructuras de datos (`structs`). Estas estructuras permiten el mapeo tipado de las respuestas JSON recibidas, asegurando la integridad de los datos en todo el programa.
* **`utils.go`**: Reúne funciones de utilidad general, como el procesamiento de los *flags* de la línea de comandos (`-host`) y el formateo estético de la salida en consola.
* **`go.mod`**: El archivo de manifiesto del módulo que garantiza que las dependencias y la versión de Go sean consistentes en cualquier entorno.

---

### 🔄 Flujo de Datos

1.  **Entrada**: `main.go` captura el dominio a través de `utils.go`.
2.  **Procesamiento**: `client.go` realiza la petición a la API externa.
3.  **Mapeo**: La respuesta JSON se transforma en objetos de Go usando las definiciones en `model.go`.
4.  **Salida**: El programa procesa los resultados y los muestra al usuario final.

## 🛠️ Tecnologías utilizadas

* **Go (Golang)**: Lenguaje de programación principal, seleccionado por su eficiencia en herramientas de CLI y concurrencia.
* **Git**: Sistema de control de versiones para el seguimiento del código.
* **SSL Labs API**: Fuente de datos externa utilizada para realizar el análisis profundo de los certificados TLS.

---
**Desarrollado por [sahlo21](https://github.com/sahlo21)**
