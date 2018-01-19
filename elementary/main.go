package main

import (
	"time"
	"github.com/lapingvino/etime"
	"honnef.co/go/js/dom"
)

func timer(clock dom.Element) {
	clock.SetInnerHTML(etime.Now().String() + " " + etime.Now().Element().String() + " <br /> @ " + time.Now().UTC().Format("02 Jan 2006") + " aUm")
	time.AfterFunc(time.Millisecond*10, func() {timer(clock)})
}

func main() {
	clock := dom.GetWindow().Document().GetElementByID("clock")
	explain := dom.GetWindow().Document().GetElementByID("explain")
	timer(clock)
	explain.SetInnerHTML("Fire, Air, Water and Earth are used in a cycle to provide four blocks of six hours each. This time is the same everywhere on earth. It was created to give a similar idea of time everywhere in the world without relying on morning/evening etc in one specific location. In your timezone, midnight is at "+etime.Descriptive("0")+ ", morning is at "+etime.Descriptive("6")+", noon is at "+etime.Descriptive("12")+ " and evening is at "+etime.Descriptive("18")+". A more visual representation is coming soon! Day is given in aUm, after UTC midnight.")
}
