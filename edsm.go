package edsm

import (
	"encoding/json"
	"net/http"
)

type Credentials struct {
	ApiKey string `json:"-,omitempty"`
}

type Service struct {
	Creds *Credentials
}

func NewService() *Service {
	res := &Service{}
	return res
}

const (
	rqUrlSystem = "https://www.edsm.net/api-v1/system"
)

func (creds *Credentials) Clear() {
	creds.ApiKey = "" // TODO is this secure… releasing that memory???
}

type RespSystem struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Coords struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"coords"`
}

func (srv *Service) System(name string) *RespSystem {
	rq, _ := http.NewRequest("GET", rqUrlSystem, nil)
	q := rq.URL.Query()
	q.Set("systemName", name)
	q.Set("showCoordinates", "1")
	q.Set("showId", "1")
	rq.URL.RawQuery = q.Encode()
	rq.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	res := &RespSystem{}
	json.NewDecoder(resp.Body).Decode(res)
	if len(res.Name) == 0 {
		return nil
	} else {
		return res
	}
}
