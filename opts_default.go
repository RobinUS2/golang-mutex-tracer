package muxtracer

import (
	"sync"
	"time"
)

var defaultGlobalOpts Opts
var defaultGlobalOptsMux sync.RWMutex

func obtainGlobalOpts() Opts {
	defaultGlobalOptsMux.RLock()
	c := defaultGlobalOpts
	defaultGlobalOptsMux.RUnlock()
	return c
}

func init() {
	// default global opts
	o := Opts{
		Threshold: 100 * time.Millisecond,
	}
	SetGlobalOpts(o)
}

func SetGlobalOpts(o Opts) {
	defaultGlobalOptsMux.Lock()
	defaultGlobalOpts = o
	defaultGlobalOptsMux.Unlock()
}
