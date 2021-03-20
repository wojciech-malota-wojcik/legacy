package main

import (
	"log"

	"github.com/wojciech-malota-wojcik/legacy/secrets"
	"github.com/wojciech-malota-wojcik/legacy/util"
)

func main() {
	util.WorkingDir(0)
	if err := secrets.Integrate(); err != nil {
		log.Fatal(err)
	}
}
