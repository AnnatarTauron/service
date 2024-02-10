package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type HealthCheck struct {
	ServiceID string
	Status    string
}

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type Checkable interface {
	GetID() string
	Health(context.Context) bool
	GetMetrics(context.Context) string
}

type Checker struct {
	targets []Checkable
	mu      sync.Mutex
}

func NewChecker() *Checker {
	return &Checker{}
}

// !mutex for concurency
func (c *Checker) Add(item Checkable) {
	c.mu.Lock()
	c.targets = append(c.targets, item)
	c.mu.Unlock()
}

func (c *Checker) String() string {

	newstring := ""
	for _, obj := range c.targets {
		newstring = newstring + obj.GetID()
	}
	return newstring //strings.Join()
}

func (c *Checker) check(ctx context.Context, cInterface Checkable) {
	if cInterface.Health(ctx) {
		fmt.Println("%v работает", cInterface.GetID())
		return
	}
	fmt.Println("%v не работает", cInterface.GetID())

}

func (c *Checker) Run() {
	c.mu.Lock()
	ctx := context.Background()
	c.run(ctx)
	c.mu.Unlock()
}

func (c *Checker) run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)

	select {
	case <-ticker.C:
		c.Check()
	case <-ctx.Done():
		fmt.Println("method run completed")
		return
	}

}

func (c *Checker) Stop() {
	fmt.Println("Stopped")
}

func (c *Checker) Check() {
	ctx := context.Background()
	for _, obj := range c.targets {
		go c.check(ctx, obj)
	}
}
