package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// waitForAnalysis: Realiza la petición a la api para que realice el analisis
// Realiza varias iteraciones hasta tener el analisis listo o reportar un error
// Durante cada intento:
//   - Si hay un error de conexion, se reintenta después de 5 segundos
//   - Si la API devuelve un código distinto de 200, se reintenta después de 5 segundos
//   - Si la decodificación JSON falla, se reintenta después de 5 segundos
//
// La función maneja diferentes estados de la respuesta de la API:
//   - "READY": el analisis terminó correctamente. Se valida que existan endpoints antes de devolver la respuesta
//   - "ERROR": la API reporta un error. La función retorna un error inmediatamente
//   - "IN_PROGRESS", "DNS", "IN_QUEUE": el analisis está en curso, se imprime el progreso y se espera antes del siguiente intento
//   - Otros estados: se imprimen como desconocidos.
//
// Retorna la respuesta de la api o un error
func waitForAnalysis(url string, params Params) (*Response, error) {
	const maxAttempts = 50
	const waitTime = 10 * time.Second
	//se declara un timeout de maximo 15 segundos para la conexion http
	client := &http.Client{Timeout: 15 * time.Second}

	//evitar reinicios multiples en startNew
	urlSeguimiento := strings.ReplaceAll(url, "&startNew=on", "")

	for attempt := 1; attempt <= maxAttempts; attempt++ {

		if attempt > 1 {
			url = urlSeguimiento
		}
		//se realiza la peticion http a la API
		resp, err := client.Get(url)
		//si falla la conexion lo reporta y reintenta en 5 segundos
		if err != nil {
			fmt.Printf("Intento %d fallido: %v\n", attempt, err)
			time.Sleep(5 * time.Second)
			continue
		}
		//si el codigo de respuesta no es 200 lo reporta y reintenta en 5 segundos
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Intento %d Status code %d\n", attempt, resp.StatusCode)
			resp.Body.Close()
			time.Sleep(5 * time.Second)
			continue
		}
		var response = &Response{}
		//se procesa la respuesta dependiendo si se solicito el reporte completo o no
		if !params.All {
			//  Decodificación directa ( rapida)
			err = json.NewDecoder(resp.Body).Decode(response)
			resp.Body.Close()
			if err != nil {
				fmt.Printf("Intento %d Error codificando: %v\n", attempt, err)
				time.Sleep(5 * time.Second)
				continue
			}
		} else {
			// Captura total de los datos (-all)
			bodyBytes, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Printf("Intento %d Error leyendo: %v\n", attempt, err)
				time.Sleep(5 * time.Second)
				continue
			}

			var todoGenerico interface{}
			if err := json.Unmarshal(bodyBytes, &todoGenerico); err == nil {
				json.Unmarshal(bodyBytes, response)

				pretty, _ := json.MarshalIndent(todoGenerico, "", "  ")
				response.RawJSON = string(pretty)
			}
		}

		//se valida que el host no este vacio para descartar dominios invalidos
		if response.Host == "" {
			return nil, fmt.Errorf("Error!! La API no devolvió el Host( posible dominio incorrecto)")

		}
		//manejo de los diferentes estados de la respuesta
		switch response.Status {
		case "READY":
			//se valida que existan endpoints en la respuesta
			if len(response.Endpoints) == 0 {
				return nil, fmt.Errorf("el análisis terminó pero no se encontraron endpoints")
			}
			return response, nil

		case "ERROR":
			// la API reporta un error
			return nil, fmt.Errorf("la API de SSL Labs reportó un error en el análisis")

		case "IN_PROGRESS", "DNS", "IN_QUEUE":
			// el análisis está en curso, se imprime el progreso
			progreso := 0
			if len(response.Endpoints) > 0 {
				progreso = response.Endpoints[0].Progress
			}
			fmt.Printf("[%d/%d] Progreso: %d%% (Estado: %s)\n",
				attempt, maxAttempts, progreso, response.Status)

		default:
			// otros estados desconocidos
			fmt.Printf("Estado desconocido: %s\n", response.Status)
		}
		//espera antes del siguiente intento
		time.Sleep(waitTime)

	}
	//si se alcanzan el maximo de intentos sin exito, se retorna un error
	return nil, fmt.Errorf("La API no respondió correctamente después de %d intentos", maxAttempts)
}

// generateURL: Construye el url completo para la consulta en la API de SSL labs con los parametros dados por el usuario
// Retorna un string con el url completo
func generateURL(params Params) string {

	baseUrl := "https://api.ssllabs.com/api/v2/analyze"
	url := baseUrl + "?host=" + params.Host
	//se concatenan en el string en caso de ser true o lo correspondiente
	if params.Publish {
		url += "&publish=on"
	}
	if params.StartNew {
		url += "&startNew=on"
	}
	if params.FromCache {
		url += "&fromCache=on"
	}
	if params.MaxAge > 0 {
		url += fmt.Sprintf("&maxAge=%d", params.MaxAge)
	}
	if params.All {
		url += "&all=on"
	}

	if params.IgnoreMismatch {
		url += "&ignoreMismatch=on"
	}

	return url
}

// generateURL2: Construye el url completo para la consulta en la API de SSL labs con los parametros dados por el usuario
// Retorna un string con el url completo
func generateURL2(params Params) (string, error) {
	base, err := url.Parse("https://api.ssllabs.com/api/v2/analyze?")
	if err != nil {
		return "", err
	}
	q := base.Query()
	q.Set("host", params.Host)
	if params.Publish {
		q.Set("publish", "on")
	}
	if params.StartNew {
		q.Set("startNew", "on")
	}
	if params.FromCache {
		q.Set("fromCache", "on")
	}
	if params.MaxAge > 0 {
		q.Set("maxAge", fmt.Sprintf("%d", params.MaxAge))
	}
	if params.All {
		q.Set("all", "on")
	}
	if params.IgnoreMismatch {
		q.Set("ignoreMismatch", "on")
	}
	base.RawQuery = q.Encode()
	return base.String(), nil
}
