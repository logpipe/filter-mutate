package main

import (
	"github.com/logpipe/logpipe/engine"
	"testing"
)

func TestMutateFilter(t *testing.T) {
	filter := &MutateFilter{
		Ops: []MutateOp{
			&AddTagOp{Tag: "ATag"},
			&AddFieldOp{Field: "AField", Value: "FiledValue"},
		},
	}
	engine.DebugFilter(filter)
	engine.Wait()
}
