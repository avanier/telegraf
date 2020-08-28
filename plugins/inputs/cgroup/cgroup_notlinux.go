// +build !linux

package cgroup

import (
	"github.com/aleveille/telegraf"
)

func (g *CGroup) Gather(acc telegraf.Accumulator) error {
	return nil
}
