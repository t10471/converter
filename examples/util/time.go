package util

import (
	"time"
)

type NowFunc func() time.Time

var (
	jst      = time.FixedZone("Asia/Tokyo", 9*60*60)
	location *time.Location
	nowFunc  NowFunc
)

func init() {
	SetNowFunc(time.Now)
	SetLocation(jst)
}

func Now() time.Time {
	return nowFunc()
}

func SetNowFunc(f NowFunc) {
	nowFunc = f
}

func Location() *time.Location {
	return location
}

func SetLocation(l *time.Location) {
	location = l
}
