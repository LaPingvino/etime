package main

import (
	"strconv"
	"time"
	"github.com/lapingvino/etime"
	"honnef.co/go/js/dom"
)

func timer(clock dom.Element) {
	clock.SetInnerHTML(currentTime() + " <br /> @ quarterday " + 
	strconv.Itoa(time.Now().UTC().YearDay()*4 -
	(3 - int(etime.Now().Element()))))
	dom.GetWindow().Document().(dom.HTMLDocument).SetTitle(currentTime() + " Elementary Time")
	time.AfterFunc(time.Millisecond*10, func() {timer(clock)})
}

func currentTime() string {
	return etime.Now().String() + " " + etime.Now().Element().String()
}

func main() {
	clock := dom.GetWindow().Document().GetElementByID("clock")
	timer(clock)
}
