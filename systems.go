package edsm

import (
	"encoding/json"
	"net/http"
)

const (
	pathSystem = "api-v1/system"
)

const (
	SYSTEM_ID uint32 = (1 << iota)
	SYSTEM_COOS
	SYSTEM_PERMIT
	SYSTEM_INFO
	SYSTEM_PRIMSTAR
	SYSTEM_HIDDEN
)

const SYSTEM_ALL uint32 = SYSTEM_ID | SYSTEM_COOS | SYSTEM_PERMIT |
	SYSTEM_INFO | SYSTEM_PRIMSTAR

type RespSystem struct {
	Flags  uint32 `json:-`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Coords struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"coords"`
	Info struct {
		Allegiance   string `json:"allegiance"`
		Government   string `json:"government"`
		Faction      string `json:"faction"`
		FactionState string `json:"factionState"`
		Population   int    `json:"population"`
		Reserve      string `json:"reserve"`
		Security     string `json:"security"`
		Economy      string `json:"economy"`
	} `json:"information"`
	PrimStar struct {
		Type      string `json:"type"`
		Name      string `json:"name"`
		Scoopable bool   `json:"isScoopable"`
	} `json:"primaryStar"`
}

func (srv *Service) System(name string, flags uint32) (*RespSystem, error) {
	rq, _ := http.NewRequest("GET", srv.url(pathSystem), nil)
	q := rq.URL.Query()
	q.Set("systemName", name)
	if (flags & SYSTEM_ID) != 0 {
		q.Set("showId", "1")
	}
	if (flags & SYSTEM_COOS) != 0 {
		q.Set("showCoordinates", "1")
	}
	if (flags & SYSTEM_PERMIT) != 0 {
		q.Set("showPermit", "1")
	}
	if (flags & SYSTEM_INFO) != 0 {
		q.Set("showInformation", "1")
	}
	if (flags & SYSTEM_PRIMSTAR) != 0 {
		q.Set("showPrimaryStart", "1")
	}
	if (flags & SYSTEM_HIDDEN) != 0 {
		q.Set("showHidden", "1")
	}
	rq.URL.RawQuery = q.Encode()
	rq.Header.Set("Accept", "application/json")
	resp, err := srv.Http.Do(rq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := &RespSystem{Flags: flags}
	json.NewDecoder(resp.Body).Decode(res)
	if len(res.Name) == 0 {
		return nil, nil // TODO does this need an error?
	}
	return res, nil
}
