package torsaver

import (
	"context"
	"time"
)

import "github.com/zyxar/argo/rpc"

// DefaultAriaRPC ...
var DefaultAriaRPC = "http://localhost:6800/jsonrpc"

// DefaultAriaSecret ...
var DefaultAriaSecret = ""

// DefaultTimeOutSecond ...
var DefaultTimeOutSecond time.Duration = 3

// NewRPCClient ...
func NewRPCClient(ctx context.Context) (cli rpc.Client, e error) {
	return rpc.New(ctx, DefaultAriaRPC, DefaultAriaSecret, DefaultTimeOutSecond*time.Second, nil)
}
