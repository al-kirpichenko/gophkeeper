package main

import (
	"os"

	"github.com/al-kirpichenko/gophkeeper/cmd/client/cfg"
	"github.com/al-kirpichenko/gophkeeper/internal/client"
)

func main() {

	conf := cfg.NewCfg()

	cli := client.NewClient(conf)

	cli.Welcome()

	os.Exit(0)
}
