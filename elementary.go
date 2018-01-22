package etime

import (
	"time"
)

type Element int

type Time struct{
	time.Time
}

const (
	FIRE Element = iota
	AIR
	WATER
	EARTH
)

func (t Time) String() string {
	t.Time = t.Time.UTC().Add(-6*time.Duration(t.Element())*time.Hour)
	return t.Format("15:04:05")
}

func (t Time) Element() Element {
	return Element(t.Time.UTC().Hour() / 6)
}

func (e Element) String() string {
	return e.Emoji() + " " + e.Name()
}

func (e Element) Name() string {
	return []string{"Fire", "Air", "Water", "Earth"}[e]
}

func (e Element) Emoji() string {
	return []string{"\U0001f525", "\U0001f32c", "\U0001f30a", "\U0001F331"}[e]
}

func (t Time) QuarterDay() int {
	daysSinceEpoch := t.Time.Unix() / (60 * 60 * 24)
	modDay := int(daysSinceEpoch%50)
	return modDay - (3-int(t.Element())) + 1
}

func Now() Time {
	return Time{time.Now()}
}
