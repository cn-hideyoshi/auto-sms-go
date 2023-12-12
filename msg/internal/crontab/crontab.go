package crontab

import (
	"log"
	"time"
)

type Crontab struct {
	closeCh chan struct{}
}

func NewCrontab() *Crontab {
	return &Crontab{
		closeCh: make(chan struct{}),
	}
}

func (c *Crontab) Start() {
	log.Println("crontab init success...")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-c.closeCh:
			return
		case <-ticker.C:
			c.doCrontab()
		}
	}
}

func (c *Crontab) StartD() {
	go c.Start()
}

func (c *Crontab) StopCrontab() {
	c.closeCh <- struct{}{}
}

func (c *Crontab) doCrontab() {
	crontab := SMSCrontab{}
	crontab.CheckSmsTime()
}
