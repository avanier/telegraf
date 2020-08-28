// +build !linux,!freebsd

package zfs

import (
	"github.com/aleveille/telegraf"
	"github.com/aleveille/telegraf/plugins/inputs"
)

func (z *Zfs) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("zfs", func() telegraf.Input {
		return &Zfs{}
	})
}
