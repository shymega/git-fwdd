package main // import "github.com/shymega/git-fwdd"

import (
	log "github.com/inconshreveable/log15"
)

func checkError(e error) {
	if e != nil {
		log.Crit(e.Error())
	}
}

func main() {
}
