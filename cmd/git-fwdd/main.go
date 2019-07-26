package main // import "git.shymega.org.uk/shymega/git-fwdd"

import (
	log "github.com/inconshreveable/log15"
)

func checkError(e error) {
	if e != nil {
		log.Crit(e.Error())
	}
}

func init() {
	log.Info("Initalising!")
}

func main() {
	log.Info("Loaded.")
}
