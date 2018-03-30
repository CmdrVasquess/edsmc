package edsm

import (
	"fmt"
	"net/http"
	"time"
)

type Endpoint string

//go:generate versioner -bno build_no -pkg edsm ./VERSION ./version.go
const (
	Software = "goEDSMc"

	// Life is the service endpoint for the life EDSM system
	Life Endpoint = "https://www.edsm.net/"
	// Test is the service endpoint recommended for testing clients
	Test Endpoint = "https://beta.edsm.net/"

	ConentType = "application/json; charset=utf-8"
)

var vStr string

func VersionStr() string {
	if len(vStr) == 0 {
		vStr = fmt.Sprintf("%d.%d.%d%s", Major, Minor, Bugfix, Quality)
	}
	return vStr
}

type Credentials struct {
	EdsmCmdr string `json:",omitempty"`
	ApiKey   string `json:",omitempty"`
}

type Service struct {
	Endp  string
	Creds *Credentials
	Game  GameStateRd
	Http  http.Client
}

func NewService(endpoint Endpoint) *Service {
	res := &Service{Endp: string(endpoint)}
	res.Http.Timeout = 8 * time.Second
	return res
}

func (creds *Credentials) Clear() {
	creds.ApiKey = "" // TODO is this secure… releasing that memory???
}

func (s *Service) url(path string) string {
	return s.Endp + path
}
