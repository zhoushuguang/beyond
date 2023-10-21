package pipeline

import (
	"context"
	"sync"
	"time"
)

type Config struct {
	MaxSize  int
	Interval time.Duration
	Buffer   int
	Worker   int
}

func (c *Config) init() {
	if c.MaxSize <= 0 {
		c.MaxSize = 1024
	}
	if c.Interval <= 0 {
		c.Interval = time.Second
	}
	if c.Buffer <= 0 {
		c.Buffer = 1024
	}
	if c.Worker <= 0 {
		c.Worker = 10
	}
}

type msg struct {
	key   string
	value interface{}
}

type Pipeline struct {
	Do       func(ctx context.Context, values map[string][]interface{})
	Sharding func(key string) int
	conf     *Config
	chans    []chan *msg
	wait     sync.WaitGroup
}

func New(conf *Config) *Pipeline {
	if conf == nil {
		conf = &Config{}
	}
	conf.init()
	p := &Pipeline{
		conf:  conf,
		chans: make([]chan *msg, conf.Worker),
	}
	for i := 0; i < conf.Worker; i++ {
		p.chans[i] = make(chan *msg, conf.Buffer)
	}

	return p
}
