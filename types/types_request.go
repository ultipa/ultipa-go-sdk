package types

import (
	"time"
)

type Request_Property = struct {
	Dataset DBType;
}

type Request_Common struct {
	GraphSetName string
	TimeoutSeconds time.Duration
	Retry *Retry
	UseHost string
}