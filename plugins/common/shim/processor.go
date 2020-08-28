package shim

import (
	"bufio"
	"fmt"
	"sync"
	"time"

	"github.com/aleveille/telegraf"
	"github.com/aleveille/telegraf/agent"
	"github.com/aleveille/telegraf/plugins/parsers"
	"github.com/aleveille/telegraf/plugins/processors"
)

// AddProcessor adds the processor to the shim. Later calls to Run() will run this.
func (s *Shim) AddProcessor(processor telegraf.Processor) error {
	setLoggerOnPlugin(processor, NewLogger())
	p := processors.NewStreamingProcessorFromProcessor(processor)
	return s.AddStreamingProcessor(p)
}

// AddStreamingProcessor adds the processor to the shim. Later calls to Run() will run this.
func (s *Shim) AddStreamingProcessor(processor telegraf.StreamingProcessor) error {
	setLoggerOnPlugin(processor, NewLogger())
	if p, ok := processor.(telegraf.Initializer); ok {
		err := p.Init()
		if err != nil {
			return fmt.Errorf("failed to init input: %s", err)
		}
	}

	s.Processor = processor
	return nil
}

func (s *Shim) RunProcessor() error {
	acc := agent.NewAccumulator(s, s.metricCh)
	acc.SetPrecision(time.Nanosecond)

	parser, err := parsers.NewInfluxParser()
	if err != nil {
		return fmt.Errorf("Failed to create new parser: %w", err)
	}

	err = s.Processor.Start(acc)
	if err != nil {
		return fmt.Errorf("failed to start processor: %w", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s.writeProcessedMetrics()
		wg.Done()
	}()

	scanner := bufio.NewScanner(s.stdin)
	for scanner.Scan() {
		m, err := parser.ParseLine(scanner.Text())
		if err != nil {
			fmt.Fprintf(s.stderr, "Failed to parse metric: %s\b", err)
			continue
		}
		s.Processor.Add(m, acc)
	}

	close(s.metricCh)
	s.Processor.Stop()
	wg.Wait()
	return nil
}
