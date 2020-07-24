package e

import "sync"

// Stores ...
type Stores struct {
	layer    int
	handlers HandlersChain
}

var once sync.Once
var gc *Stores

// NewStores ...
func NewStores() *Stores {
	once.Do(func() {
		gc = &Stores{layer: 1}
	})
	return gc
}

// Setlayer ...
func (g *Stores) Setlayer(layer int) {
	g.layer = layer
}

// Getlayer ...
func (g *Stores) Getlayer() int {
	return g.layer
}

// Setlayer 设置跟踪错误堆栈的层数
func Setlayer(layer int) {
	NewStores().Setlayer(layer)
}

// Getlayer 获取设置的跟踪错误堆栈的层数
func Getlayer() int {
	return NewStores().Getlayer()
}
