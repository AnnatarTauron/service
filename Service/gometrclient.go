package main

import (
	"context"
	"time"
)

type GoMetrClient struct {
	serviceID string
	timeout   int
}

func NewGoMetrClient(serviceID string) *GoMetrClient {
	return &GoMetrClient{serviceID: serviceID}
}

func (g GoMetrClient) getHealth(ctx context.Context) HealthCheck {
	return HealthCheck{
		ServiceID: g.GetID(),
		Status:    "healthy",
	}
}

func (g GoMetrClient) GetMetrics(ctx context.Context) string {
	return ""
}
func (g GoMetrClient) Health(ctx context.Context) bool {
	cont, _ := context.WithTimeout(ctx, time.Duration(g.timeout))
	select {
	case <-cont.Done():
		return false
	default:
		return g.getHealth(ctx).Status == "healthy"

	}
}

func (g GoMetrClient) GetID() string {
	return g.serviceID
}
