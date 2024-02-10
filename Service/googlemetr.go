package main

import (
	"context"
	"time"
)

type GoogleMetrClient struct {
	serviceID string
	timeout   int
}

func NewGoogleMetrClient(serviceID string) *GoogleMetrClient {
	return &GoogleMetrClient{serviceID: serviceID}
}

func (g GoogleMetrClient) getHealth(ctx context.Context) HealthCheck {
	return HealthCheck{
		ServiceID: c.GetID(),
		Status:    "healthy",
	}
}

func (g GoogleMetrClient) GetMetrics(ctx context.Context) string {
	return ""
}
func (g GoogleMetrClient) Health(ctx context.Context) bool {
	cont, _ := context.WithTimeout(ctx, time.Duration(g.timeout))
	select {
	case <-cont.Done():
		return false
	default:
		return g.getHealth(ctx).Status == "healthy"

	}
}

func (g GoogleMetrClient) GetID() string {
	return g.serviceID
}
