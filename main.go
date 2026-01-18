package main

/**
 * TLS Checker by sahlo21
 * Repositorio: https://github.com/sahlo21/tls-checker
 Uso:
  ./tls-checker -host www.ejemplo.com [opciones]

Opciones disponibles:
  -all
        Analizar todos los endpoint del dominio. Desactivado por defecto
  -fromCache
        Usar resultados guardados en cache en caso de existir. Desactivado por defecto
  -host string
        Dominio a analizar (Requerido, Ejemplo: www.ejemplo.com)
  -ignoreMismatch
        Continuir analisis aunque el certificado no coincida con el host. Desactivado por defecto
  -maxAge int
        Antiguedad maxima del informe cuando de tomar del cache guardado(fromCache). Desactivado por defecto
  -publish
        Publicar resultados en tableros publicos. Desactivado por defecto
  -startNew
        Iniciar un analisis nuevo ignorando los guardados en cache. Desactivado por defecto
*/
import (
	"fmt"
	"log"
)

// Main es la función principal donde se iniliazaran los parametros, se genera el URL, se realiza en analisis de la api y se presentan los resultados
func main() {

	params := processParams()
	printParams(params)

	url := generateURL2(params)

	fmt.Println("Analizando el dominio...")

	response, err := waitForAnalysis(url)
	if err != nil {
		log.Fatalf("Error al obtener resultados: %v", err)
	}
	printResult(response)

}
