package main

import (
	"time"
	"github.com/lapingvino/etime"
	"honnef.co/go/js/dom"
)

func timer(clock dom.Element) {
	clock.SetInnerHTML(etime.Now().String() + " " + etime.Now().Element().String())
	time.AfterFunc(time.Millisecond*10, func() {timer(clock)})
}

func main() {
	clock := dom.GetWindow().Document().GetElementByID("clock")
	explain := dom.GetWindow().Document().GetElementByID("explain")
	timer(clock)
	explain.SetInnerHTML("Earth, Water, Air and Fire are used in a cycle to provide four blocks of six hours each. This time is the same everywhere on earth. It is meant to have the same mental idea of time everywhere in the world without relying on morning/evening etc in one specific location. In your local timezone, midnight is at "+etime.Descriptive("0")+ ", morning is at "+etime.Descriptive("6")+", noon is at "+etime.Descriptive("12")+ " and evening is at "+etime.Descriptive("18")+".")
}