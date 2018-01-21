package main

import (
	"strconv"
	"time"
	"github.com/lapingvino/etime"
	"honnef.co/go/js/dom"
)

func timer(clock dom.Element) {
	clock.SetInnerHTML(etime.Now().String() + " " +
	etime.Now().Element().String() + " <br /> @ quarterday " + 
	strconv.Itoa(time.Now().UTC().YearDay()*4 -
	(3 - int(etime.Now().Element()))))
	time.AfterFunc(time.Millisecond*10, func() {timer(clock)})
}

func main() {
	clock := dom.GetWindow().Document().GetElementByID("clock")
	timer(clock)
}
