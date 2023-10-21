package protocol

import (
	"sync"
)
// DevToolsQueue 是一个线程安全的队列结构
type DevToolsQueue struct {
	items []map[string]func(param interface{}) // 存储队列项的切片
	lock  sync.Mutex                              // 互斥锁，用于保护队列的并发访问
}

// Enqueue 将项添加到队列末尾
func (q *DevToolsQueue) Enqueue(item map[string]func(param interface{})) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, item)
}

// Dequeue 从队列中移除并返回队首的项
func (q *DevToolsQueue) Dequeue() map[string]func(param interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// DevToolsQueueMap 是一个线程安全的 DevToolsQueue 映射
type DevToolsQueueMap struct {
	queues map[string]*DevToolsQueue
	lock   sync.Mutex // 互斥锁，用于保护映射的并发访问
}

// NewDevToolsQueueMap 创建一个新的 DevToolsQueueMap
func NewDevToolsQueueMap() *DevToolsQueueMap {
	return &DevToolsQueueMap{
		queues: make(map[string]*DevToolsQueue),
	}
}

// GetQueue 根据 key 获取对应的 DevToolsQueue，如果不存在则返回 nil
func (m *DevToolsQueueMap) GetQueue(key string) *DevToolsQueue {
	m.lock.Lock()
	defer m.lock.Unlock()

	queue, ok := m.queues[key]
	if !ok {
		return nil
	}

	return queue
}

// Enqueue 将项添加到指定 key 对应的 DevToolsQueue 的末尾
func (m *DevToolsQueueMap) Enqueue(key string, item map[string]func(param interface{})) {
	m.lock.Lock()
	defer m.lock.Unlock()

	queue, ok := m.queues[key]
	if !ok {
		queue = &DevToolsQueue{}
		m.queues[key] = queue
	}
	queue.Enqueue(item)
}

// Dequeue 从指定 key 对应的 DevToolsQueue 中移除并返回队首的项
func (m *DevToolsQueueMap) Dequeue(key string) map[string]func(param interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	queue, ok := m.queues[key]
	if !ok {
		return nil
	}

	return queue.Dequeue()
}


// Peek 获取指定 key 对应的 DevToolsQueue 中队首的项，不移除该项
func (m *DevToolsQueueMap) Peek(key string) map[string]func(param interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	queue, ok := m.queues[key]
	if !ok {
		return nil
	}

	if len(queue.items) == 0 {
		return nil
	}

	item := queue.items[0]
	return item
}

// Remove 删除指定 key 对应的 DevToolsQueue 中的队首项
func (m *DevToolsQueueMap) Remove(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	queue, ok := m.queues[key]
	if !ok {
		return
	}

	if len(queue.items) == 0 {
		return
	}

	queue.items = queue.items[1:]
}
