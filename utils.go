package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"
)

// printParams Imprime el encabezado y las opciones seleccionadas por el usuario (Parametros)
func printParams(params Params) {
	//encabezado
	fmt.Println("========================================")
	fmt.Println("   SSL Labs TLS Security Analyzer")
	fmt.Println("========================================")

	//Opciones seleccionados por el usuario (Parametros)
	fmt.Println("\n\nParamentros a analizar: ")
	fmt.Printf("%-20s: %v\n", "Host", params.Host)
	fmt.Printf("%-20s: %v\n", "Publish", params.Publish)
	fmt.Printf("%-20s: %v\n", "StartNew", params.StartNew)
	fmt.Printf("%-20s: %v\n", "FromCache", params.FromCache)
	fmt.Printf("%-20s: %v\n", "MaxAge", params.MaxAge)
	fmt.Printf("%-20s: %v\n", "All", params.All)
	fmt.Printf("%-20s: %v\n\n", "IgnoreMismatch", params.IgnoreMismatch)
}

// printResult: Imprime de forma mas organizada y legible los resultados,
// incluido la respuesta completa y de cada endopoint retornada por la API
func printResult(response *Response, params Params) {
	// Se imprimen los resultados completos enviados por la APi

	fmt.Println("\n\n========================================")
	fmt.Printf("%-20s: %v\n", "Domain", response.Host)
	fmt.Printf("%-20s: %v\n", "Port", response.Port)
	fmt.Printf("%-20s: %v\n", "Protocol", response.Protocol)
	fmt.Printf("%-20s: %v\n", "IsPublic", response.IsPublic)
	fmt.Printf("%-20s: %v\n", "Status", response.Status)
	fmt.Printf("%-20s: %v\n", "StartTime", formatMs(response.StartTime))
	fmt.Printf("%-20s: %v\n", "TestTime", formatMs(response.TestTime))
	fmt.Printf("%-20s: %v\n", "EngineVersion", response.EngineVersion)
	fmt.Printf("%-20s: %v\n", "CriteriaVersion", response.CriteriaVersion)

	// Se imprimen los endpoints
	for i, endp := range response.Endpoints {
		fmt.Println("\n-------------------------------------")
		fmt.Printf("             Endpoint #%d              \n", i+1)
		fmt.Println("-------------------------------------")
		fmt.Printf("%-20s: %v\n", "Ip address", endp.IpAddress)
		if endp.StatusMessage == "" {
			endp.StatusMessage = "N/A"
		}
		fmt.Printf("%-20s: %v\n", "StatusMessage:", endp.StatusMessage)
		if endp.Grade == "" {
			endp.Grade = "N/A"
		}
		fmt.Printf("%-20s: %v\n", "Grade", endp.Grade)
		if endp.GradeTrustIgnored == "" {
			endp.GradeTrustIgnored = "N/A"
		}
		fmt.Printf("%-20s: %v\n", "GradeTrustIgnored", endp.GradeTrustIgnored)
		fmt.Printf("%-20s: %v\n", "HasWarnings", endp.HasWarnings)
		fmt.Printf("%-20s: %v\n", "IsExceptional", endp.IsExceptional)
		fmt.Printf("%-20s: %v\n", "Progress", endp.Progress)
		fmt.Printf("%-20s: %v\n", "Duration", formatDuration(endp.Duration))
		fmt.Printf("%-20s: %v\n", "Eta", endp.Eta)
		fmt.Printf("%-20s: %v\n", "Delegation", endp.Delegation)

	}
	if params.All {
		folderName := "json_results"
		nombreArchivo := fmt.Sprintf("%s/%s_report.json", folderName, params.Host)
		err := os.WriteFile(nombreArchivo, []byte(response.RawJSON), 0644)

		if err != nil {
			fmt.Printf("Error al guardar el reporte: %v\n", err)
		} else {
			fmt.Printf("\n El reporte completo ha sido guardado en: ./%s\n", nombreArchivo)
		}
	}
}

//utilidades

// formatMs: convierte milisegundos a un formato legible
func formatMs(ms int64) string {
	if ms == 0 {
		return "N/A"
	}
	return time.Unix(0, ms*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
}

// formatDuration: convierte milisegundos en un formato dado por minutos y segundos
func formatDuration(ms int64) string {
	duration := time.Duration(ms) * time.Millisecond
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%dm %ds", minutes, seconds) // [cite: 4, 10, 16]
}

// processParams:
// este proces se encarga de definir y gestionar los parametros y comandos que el usuario
// podrá utilizar los cuales serán interpertrados para enviar a la API
// Retorna los paramentros inicializados
func processParams() Params {
	host := flag.String("host", "", "Dominio a analizar (Requerido, Ejemplo: www.ejemplo.com)")
	publish := flag.Bool("publish", false, "Publicar resultados en tableros publicos. Desactivado por defecto")
	startNew := flag.Bool("startNew", false, "Iniciar un analisis nuevo ignorando los guardados en cache. Desactivado por defecto")
	fromCache := flag.Bool("fromCache", false, "Usar resultados guardados en cache en caso de existir. Desactivado por defecto")
	maxAge := flag.Int("maxAge", 0, "Antiguedad maxima del informe cuando de tomar del cache guardado(fromCache). Desactivado por defecto")
	all := flag.Bool("all", false, "        Analizar todos los endpoint del dominio. Desactivado por defecto ( se genera un .json debido al gran tamaño de la informacińn)")
	ignoreMismatch := flag.Bool("ignoreMismatch", false, "Continuir analisis aunque el certificado no coincida con el host. Desactivado por defecto")

	flag.Usage = func() {
		fmt.Println("========================================")
		fmt.Println("   SSL Labs TLS Security Analyzer")
		fmt.Println("========================================")
		fmt.Println("Uso:")
		fmt.Println("  ./tls-checker -host www.ejemplo.com [opciones]")
		fmt.Println("\nOpciones disponibles:")
		flag.PrintDefaults()
	}

	flag.Parse()
	//validaciones basicas del dominio
	if *host == "" {
		fmt.Println("Error!! Debe ingresar el dominio a analizar")
		flag.Usage()
		os.Exit(1)
	}

	validHost := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
	if !validHost.MatchString(*host) {
		fmt.Println("Error!! El dominio ingresado no es valido")
		flag.Usage()
		os.Exit(1)
	}
	//validaciones de funciones que no son compatibles
	if *startNew && *fromCache {
		fmt.Println("Error!! Ambas funciones no son compatibles(startNew y fromCache)")
		flag.Usage()
		os.Exit(1)

	}
	if *maxAge > 0 && *fromCache == false {
		fmt.Println("Error!! maxAge solo se puede usar con informes desde el cache")
		flag.Usage()
		os.Exit(1)

	}

	return Params{
		Host:           *host,
		Publish:        *publish,
		StartNew:       *startNew,
		FromCache:      *fromCache,
		MaxAge:         *maxAge,
		All:            *all,
		IgnoreMismatch: *ignoreMismatch,
	}
}
