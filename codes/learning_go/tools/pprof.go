package tools

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/pprof"
	"time"
)

// Usage of net/http/pprof.
// more: https://pkg.go.dev/net/http/pprof
func StartPprof(addr string) {
	go func() {
		log.Println(http.ListenAndServe(addr, nil))
	}()
}

const (
	ProfileGoroutine    = "goroutine"
	ProfileHeap         = "heap"
	ProfileThreadcreate = "threadcreate"
	ProfileBlock        = "block"
	ProfileMutex        = "mutex"
)

// Dump profiles.
// name: goroutine, heap, threadcreate, block, mutex.
func DumpProfiles(name string, d time.Duration) {
	go func() {
		profile := pprof.Lookup(name)
		for range time.Tick(d) {
			log.Printf("%s count: %d\n", profile.Name(), profile.Count())
			profile.WriteTo(log.Writer(), 1)
		}
	}()
}

// Create profile.
func NewProfile(name string) *pprof.Profile {
	profile := pprof.Lookup(name)
	if profile == nil {
		profile = pprof.NewProfile(name)
	}
	return profile
}
