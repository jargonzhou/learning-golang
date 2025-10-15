package tools

import (
	"log"
	"os"
)

// Set up log.
func SetUpLog() {
	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)
}
