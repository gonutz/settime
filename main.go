package main

import (
	"time"

	"github.com/beevik/ntp"
	"github.com/gonutz/w32"
)

func main() {
	for {
		t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
		if err != nil {
			time.Sleep(time.Second)
		} else {
			ok := w32.SetLocalTime(&w32.SYSTEMTIME{
				Year:         uint16(t.Year()),
				Month:        uint16(t.Month()),
				Day:          uint16(t.Day()),
				Hour:         uint16(t.Hour()),
				Minute:       uint16(t.Minute()),
				Second:       uint16(t.Second()),
				Milliseconds: 0,
			})
			if !ok {
				panic("failed to set time, must be run as admin")
			}
			return
		}
	}
}
