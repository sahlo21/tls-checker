package main

import (
	"strings"
	"testing"
)

func TestGenerateURL(t *testing.T) {
	misParams := Params{
		Host:    "truora.com",
		Publish: true,
	}

	resultadoUrl := generateURL(misParams)

	if !strings.Contains(resultadoUrl, "host=truora.com") {
		t.Error("La url no contiene el host correcto")
	}
	if !strings.Contains(resultadoUrl, "publish=on") {
		t.Error("La url no activo el parametro publish")
	}
}

func TestAPI(t *testing.T) {
	casos := []struct {
		nombreTest string
		params     Params
	}{
		{
			nombreTest: "SSL labs en Cache",
			params:     Params{Host: "ssllabs.com"},
		},
		{
			nombreTest: "Truora en Cache",
			params:     Params{Host: "truora.com", FromCache: true},
		},
		{
			nombreTest: "Analisis Nuevo",
			params:     Params{Host: "badssl.com", StartNew: true},
		},
		{
			nombreTest: "Reporte detallado (All)",
			params:     Params{Host: "ssllabs.com", All: true, FromCache: true},
		},
		{
			nombreTest: "Publicar Resultados",
			params:     Params{Host: "ssllabs.com", Publish: true, FromCache: true, MaxAge: 10},
		},
		{
			nombreTest: "Ignorar Mismatch(certificado no coincide con el host )",
			params:     Params{Host: "badssl.com", IgnoreMismatch: true, FromCache: true},
		},
	}

	for i, caso := range casos {

		t.Run(caso.nombreTest, func(t *testing.T) {
			miUrl := generateURL(caso.params)

			respuesta, err := waitForAnalysis(miUrl, caso.params)

			if err != nil {
				t.Fatalf("CASO %d La conexión falló para %s: %v", i+1, caso.params.Host, err)
			}

			if respuesta.Host == "" {
				t.Errorf("CASO %d Error: El campo Host llego vacio de la API", i+1)
			}

			if respuesta.Status == "" {
				t.Errorf("CASO %d Error: No se recibio un estado (Status)", i+1)
			}

			t.Logf("CASO %d Resultado exitoso para %s. Estado: %s", i+1, respuesta.Host, respuesta.Status)
		})
	}
}
