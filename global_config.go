package e

import "sync"

type Config struct {
	Layer int
}

var once sync.Once
var gc *Config

func NewConfig() *Config {
	once.Do(func() {
		gc = &Config{}
	})
	return gc
}

func (g *Config) SetLayer(layer int) *Config {
	g.Layer = layer
	return g
}

func (g *Config) GetLayer() int {
	return g.Layer
}