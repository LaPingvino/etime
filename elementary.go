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

var World = time.FixedZone("World", 53100) // 18:15 in Tehran time converted to UTC converted to seconds

func (t Time) BaseAdd(d time.Duration) Time {

	t.Time = t.Time.In(World).Add(-d)
	return t
}

func (t Time) String() string {
	return t.BaseAdd(6*time.Duration(t.Element())*time.Hour).Format("15:04:05")
}

func (t Time) Element() Element {
	return Element(t.BaseAdd(0).Time.Hour() / 6)
}

func (e Element) String() string {
	return []string{"\U0001f525 Fire", "\U0001f32c Air", "\U0001f30a Water", "\U0001F331 Earth"}[e]
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
