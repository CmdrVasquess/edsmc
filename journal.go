package edsm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	pathJournal = "api-journal-v1"
	pathDiscard = "api-journal-v1/discard"
)

func (srv *Service) Discard(events *[]string) error {
	rq, _ := http.NewRequest("GET", srv.url(pathDiscard), nil)
	q := rq.URL.Query()
	rq.URL.RawQuery = q.Encode()
	rq.Header.Set("Accept", "application/json")
	resp, err := srv.Http.Do(rq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(events)
	return err
}

type Command int

const Unknown = -1

const (
	InCommand Command = 0
	NoCommand Command = 1
)

type GameStateRd interface {
	CmdrName() string
	SysAddr() int64
	SysName() string
	SysCoo() []float64
	StationId() int64
	StationName() string
	ShipId() int
	Command() Command
}

type GameState interface {
	GameStateRd
	SetSysAddr(v int64)
	SetSysName(v string)
	SetSysCoo(v []float64)
	SetStationId(v int64)
	SetStationName(v string)
	SetShipId(v int)
	SetCommand(v Command)
}

func UpdateState(gs GameState, event map[string]interface{}) {
	evtNm := event["event"].(string)
	switch evtNm {
	case "LoadGame":
		gs.SetSysAddr(-1)
		gs.SetSysName("")
		gs.SetSysCoo(nil)
		gs.SetStationId(-1)
		gs.SetStationName("")
	case "SetUserShipName", "ShipyardSwap", "Loadout":
		gs.SetShipId(int(event["ShipId"].(float64)))
	case "ShipyardBuy":
		gs.SetShipId(-1)
	case "Undocked":
		gs.SetStationId(-1)
		gs.SetStationName("")
	case "Location", "FSDJump", "Docked":
		evtSysNm := event["StarSystem"].(string)
		if evtSysNm != gs.SysName() {
			gs.SetSysCoo(nil)
		}
		if evtSysNm != "ProvingGround" && evtSysNm != "CQC" {
			if sysAddr := event["SystemAddress"]; sysAddr != nil {
				gs.SetSysAddr(int64(sysAddr.(float64)))
			}
			gs.SetSysName(evtSysNm)
			if pos := event["StarPos"]; pos != nil {
				coo := pos.([]interface{})
				gs.SetSysCoo([]float64{
					coo[0].(float64),
					coo[1].(float64),
					coo[2].(float64),
				})
			}
		} else {
			gs.SetSysAddr(-1)
			gs.SetSysName("")
			gs.SetSysCoo(nil)
		}
		if market := event["MarketID"]; market != nil {
			gs.SetStationId(int64(market.(float64)))
		}
		if statNm := event["StationName"]; statNm != nil {
			gs.SetStationName(statNm.(string))
		}
	case "JoinACrew", "QuitACrew":
		if evtNm == "JoinACrew" {
			if captain := event["Captain"]; captain == nil {
				// TODO verify
				if len(gs.CmdrName()) == 0 {
					gs.SetCommand(InCommand)
				} else {
					gs.SetCommand(NoCommand)
				}
			} else if captain.(string) == gs.CmdrName() {
				gs.SetCommand(InCommand)
			} else {
				gs.SetCommand(NoCommand)
			}
		} else {
			gs.SetCommand(InCommand)
		}
		gs.SetSysAddr(-1)
		gs.SetSysName("")
		gs.SetSysCoo(nil)
		gs.SetStationId(-1)
		gs.SetStationName("")
	}
}

type postJournal struct {
	Cmdr       string      `json:"commanderName"`
	ApiKey     string      `json:"apiKey,omitempty"`
	SWare      string      `json:"fromSoftware"`
	SWareVers  string      `json:"fromSoftwareVersion"`
	Event      string      `json:"message"`
	TrSysAddr  interface{} `json:"_systemAddress,omitempty"`
	TrSysName  string      `json:"_systemName,omitempty"`
	TrSysCoos  []float64   `json:"_systemCoordinates,omitempty"`
	TrMarketId interface{} `json:"_marketId,omitempty"`
	TrStatNm   string      `json:"_stationName,omitempty"`
	TrShipId   interface{} `json:"_shipId,omitempty"`
}

func (srv *Service) Journal(cmdr string, event string) error {
	bdy := postJournal{
		Cmdr:      cmdr,
		SWare:     Software,
		SWareVers: VersionStr(),
		Event:     event,
	}
	if srv.Creds != nil {
		bdy.ApiKey = srv.Creds.ApiKey
	}
	if srv.Game != nil {
		g := srv.Game
		if g.SysAddr() >= 0 {
			bdy.TrSysAddr = g.SysAddr()
		}
		bdy.TrSysName = g.SysName()
		bdy.TrSysCoos = g.SysCoo()
		if g.StationId() >= 0 {
			bdy.TrMarketId = g.StationId()
		}
		bdy.TrStatNm = g.StationName()
		if g.ShipId() >= 0 {
			bdy.TrShipId = g.ShipId()
		}
	}
	rqStr, err := json.Marshal(bdy)
	if err != nil {
		return err
	}
	rd := bytes.NewBuffer(rqStr)
	resp, err := srv.Http.Post(srv.url(pathJournal), ConentType, rd)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, string(msg))
	}
	return nil
}
