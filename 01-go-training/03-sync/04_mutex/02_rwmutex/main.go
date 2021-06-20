package main

import "sync"

// RWMutexConfig 读写锁实现
type RWMutexConfig struct {
	rw   sync.RWMutex
	data []int
}

// Get get config data
func (c *RWMutexConfig) Get() []int {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.data
}

// Set set config data
func (c *RWMutexConfig) Set(n []int) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.data = n
}

// MutexConfig 互斥锁实现
type MutexConfig struct {
	data []int
	mu   sync.Mutex
}

// Get get config data
func (c *MutexConfig) Get() []int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data
}

// Set set config data
func (c *MutexConfig) Set(n []int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = n
}
