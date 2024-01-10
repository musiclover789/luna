package luna_utils

import (
	"sync"
)

// 定义全局变量
var IdGen = &IDGenerator{}

// IDGenerator 生成递增数字的结构体
type IDGenerator struct {
	mu sync.Mutex
	id int
}

// NextID 生成下一个递增数字
func (g *IDGenerator) NextID() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.id++
	return g.id
}
