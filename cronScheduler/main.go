package main

import (
	"context"
	"time"
)

type JobFunc func(ctx context.Context)

type Job struct {
	Name     string
	Run      JobFunc
	interval time.Duration
}
