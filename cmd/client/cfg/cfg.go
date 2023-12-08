package cfg

const ServerURL = "http://localhost:8080"

type Cfg struct {
	ServerURL string
}

func NewCfg() *Cfg {
	return &Cfg{ServerURL: ServerURL}
}
