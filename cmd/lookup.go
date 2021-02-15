package cmd

import (
	"log"

	"github.com/bhoriuchi/dossier/config"
)

// Lookup performs a lookup
func Lookup(opts config.Options) (err error) {
	conf := config.Config{}

	if err = conf.Read(opts); err != nil {
		return
	}

	log.Printf("%+v", conf)

	return
}
