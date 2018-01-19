package etime

import (
	"time"
)

type Element int

type Time struct{
	time.Time
}

const (
	EARTH Element = iota
	WATER
	AIR
	FIRE
)

func (t Time) String() string {
	t.Time = t.Time.UTC().Add(-6*time.Duration(t.Element())*time.Hour)
	return t.Format("03:04:05")
}

func (t Time) Element() Element {
	return Element(t.Time.UTC().Hour() / 6)
}

func (e Element) String() string {
	return []string{"Earth", "Water", "Air", "Fire"}[e]
}

func Descriptive(h string) string {
	local, err := time.LoadLocation("Local")
	if err != nil {
		return "[problem loading location]"
	}
	r, err := time.ParseInLocation("15", h, local)
	if err != nil {
		return "[problem parsing reference hour]"
	}
	t := Time{r}
	seg := t.Time.UTC().Hour() % 6
	var segs string
	switch seg {
	case 0,1:
		segs = "early"
	case 2,3:
		segs = "mid"
	case 4,5:
		segs = "late"
	default:
		segs = "[there is a bug in the software]"
	}
	return segs + " " + t.Element().String()
}

func Now() Time {
	return Time{time.Now()}
}
