package parallel

import "github.com/aleveille/telegraf"

type Parallel interface {
	Enqueue(telegraf.Metric)
	Stop()
}
