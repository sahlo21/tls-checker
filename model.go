package main

// Endpoint representa la respuesta completa de la API de SSL labs
type Response struct {
	Host            string     `json:"host"`
	Port            int        `json:"port"`
	Protocol        string     `json:"protocol"`
	IsPublic        bool       `json:"isPublic"`
	Status          string     `json:"status"`
	StartTime       int64      `json:"startTime"`
	TestTime        int64      `json:"testTime"`
	EngineVersion   string     `json:"engineVersion"`
	CriteriaVersion string     `json:"criteriaVersion"`
	Endpoints       []Endpoint `json:"endpoints"`
}

// Endpoint representa los resultados de cada endpoint analizado por SSL labs
type Endpoint struct {
	IpAddress         string `json:"ipAddress"`
	StatusMessage     string `json:"statusMessage"`
	Grade             string `json:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	HasWarnings       bool   `json:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional"`
	Progress          int    `json:"progress"`
	Duration          int64  `json:"duration"`
	Eta               int64  `json:"eta"`
	Delegation        int    `json:"delegation"`
}

// Params contiene los parametros de configuración que el usuario puede utilizar, se utlizara con la libreria Flag
type Params struct {
	Host           string
	Publish        bool
	StartNew       bool
	FromCache      bool
	MaxAge         int
	All            bool
	IgnoreMismatch bool
}
