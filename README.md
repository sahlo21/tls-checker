# TLS Checker CLI 🛡️

Una herramienta de línea de comandos (CLI) desarrollada en **Go** que permite analizar la seguridad y configuración de los certificados TLS de cualquier dominio utilizando la API pública de **SSL Labs**.

## 📋 Descripción

Este proyecto permite consultar el estado de los certificados SSL/TLS de forma remota, devolviendo información detallada sobre la calificación de seguridad del servidor. Ha sido diseñado bajo un enfoque modular para asegurar que el código sea fácil de leer, mantener y testear.

**Características principales:**

* ✨ Análisis en tiempo real o desde caché.
* ✨ Soporte para múltiples endpoints (IPv4/IPv6).
* ✨ **Exportación de reportes detallados en formato JSON.**
* ✨ Validación de dominios y manejo robusto de errores.

## 🚀 Instalación y Uso

### Requisitos previos

* **Go** 1.21 o superior instalado localmente.
* Conexión a internet para realizar las peticiones a la API externa.

### Configuración

1. **Clonar el repositorio:**
```bash
git clone https://github.com/sahlo21/tls-checker.git
cd tls-checker

```



### Ejecución

Puedes ejecutar la herramienta de varias formas según lo que necesites:

**1. Análisis básico (Rápido):**

```bash
go run . -host google.com

```

**2. ✨ Forzar un nuevo análisis (Ignorar caché):**
Ideal para obtener el estado en tiempo real.

```bash
go run . -host google.com -startNew

```

**3. ✨ Generar reporte completo JSON:**
Esto guardará un archivo detallado en la carpeta `json_results/`.

```bash
go run . -host google.com -all

```

**4. ✨ Usar caché (Resultados instantáneos):**

```bash
go run . -host google.com -fromCache

```

### Opciones Disponibles (Flags)

| Flag | Descripción |
| --- | --- |
| `-host` | **(Requerido)** El dominio a analizar (ej: `google.com`). |
| `-startNew` | Inicia un análisis nuevo ignorando resultados previos. |
| `-fromCache` | Intenta obtener resultados guardados para una respuesta inmediata. |
| `-all` | Genera un reporte detallado `.json` en la carpeta `json_results`. |
| `-publish` | Publica los resultados en los tableros públicos de SSL Labs. |
| `-ignoreMismatch` | Continúa el análisis aunque el certificado no coincida con el host. |
| `-maxAge` | Define la antigüedad máxima (en horas) aceptada para el caché. |

## 🏗️ Arquitectura del Proyecto

El código ha sido estructurado siguiendo principios de **diseño modular** para separar las responsabilidades y facilitar la escalabilidad.

### 📂 Estructura de Archivos

* **`main.go`**: Punto de entrada. Orquesta el flujo, inicializa parámetros y gestiona el ciclo de vida.
* **`transport.go`** (Antes `client.go`): Capa de comunicación HTTP. Gestiona las peticiones a la API, el *polling* de estado, la limpieza de URLs y los *timeouts*.
* **`model.go`**: Definición de estructuras (`structs`) para el mapeo tipado de las respuestas JSON de SSL Labs.
* **`utils.go`**: Funciones auxiliares para validación de regex, manejo de flags, creación de carpetas de reportes y formateo de salida.
* **`utils_test.go`**: ✨ Set de pruebas unitarias para validar la generación de URLs y la integridad de los parámetros.

---

### 🔄 Flujo de Datos

1. **Entrada**: Se validan los flags y el formato del dominio (Regex).
2. **Petición**: `transport.go` decide si iniciar un scan (`startNew`) o consultar caché.
3. **Polling**: Si el análisis está en curso, el sistema consulta periódicamente el estado hasta obtener un grado (Grade) o finalizar.
4. **Resultados**:
* Si es básico: Se imprime en consola.
* Si es `-all`: Se estructura el JSON y se guarda en `json_results/[dominio]_report.json`.



## 🛠️ Tecnologías utilizadas

* **Go (Golang)**: Lenguaje principal, uso de `net/http`, `encoding/json` y `flag`.
* **Testing**: Paquete nativo `testing` de Go para pruebas unitarias.
* **SSL Labs API v2**: Motor de análisis de seguridad.

---

**Desarrollado por [sahlo21**](https://github.com/sahlo21)