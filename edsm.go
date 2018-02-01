package edsm

type Credentials struct {
	ApiKey string `json:"-,omitempty"`
}

type Service struct {
	Creds *Credentials
}

func (creds *Credentials) Clear() {
	creds.ApiKey = "" // TODO is this secure… releasing that memory???
}
