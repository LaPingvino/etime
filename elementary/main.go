package main

import (
	"strconv"
	"time"
	"github.com/brualemar/etime"
	"honnef.co/go/js/dom"
)

func timer(clock dom.Element) {
	clock.SetInnerHTML(currentTime(false))
	dom.GetWindow().Document().(dom.HTMLDocument).SetTitle(currentTime(true) + " Hora Elementar")
	time.AfterFunc(time.Second - time.Duration(time.Now().Nanosecond()), func() {timer(clock)})
}

func currentTime(short bool) string {
	if short {
		return etime.Now().Element().Emoji() + " " + etime.Now().String()[:5] + " " + etime.Now().Element().Name()[:1] + " @ " + strconv.Itoa(etime.Now().QuarterDay())
	} else {
		return etime.Now().String() + " " + etime.Now().Element().String() + " @ " + strconv.Itoa(etime.Now().QuarterDay())
	}
}

func main() {
	clock := dom.GetWindow().Document().GetElementByID("clock")
	timer(clock)
}
