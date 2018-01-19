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
	t.Time = t.Time.UTC().Add(-6*time.Duration(t.Time.UTC().Hour()%6)*time.Hour)
	return t.Format("03:04:05")
}

func (t Time) Element() Element {
	return Element(t.Time.UTC().Hour() % 6)
}

func (e Element) String() string {
	return []string{"Earth", "Water", "Air", "Fire"}[e]
}

func Now() Time {
	return Time{time.Now()}
}
