package main

import (
	"strconv"
	"time"
	"github.com/lapingvino/etime"
	"honnef.co/go/js/dom"
)

func timer(clock dom.Element) {
	clock.SetInnerHTML(currentTime(false) + " <br /> @ quarterday " + 
	strconv.Itoa(time.Now().UTC().YearDay()*4 -
	(3 - int(etime.Now().Element()))))
	dom.GetWindow().Document().(dom.HTMLDocument).SetTitle(currentTime(true) + " Elementary Time")
	time.AfterFunc(time.Millisecond*10, func() {timer(clock)})
}

func currentTime(short bool) string {
	if short {
		return etime.Now().Element().Emoji() + " " + etime.Now().String()[:5] + " " + etime.Now().Element().String()[:1]
	} else {
		return etime.Now().String() + " " + etime.Now().Element().String()
	}
}

func main() {
	clock := dom.GetWindow().Document().GetElementByID("clock")
	timer(clock)
}
