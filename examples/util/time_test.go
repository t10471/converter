package util

import (
	"reflect"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{"Success", time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)},
	}
	for _, tt := range tests {
		SetNowFunc(func() time.Time {
			return time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := Now(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Now() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocation(t *testing.T) {
	tests := []struct {
		name string
		want *time.Location
	}{
		{"Success", time.UTC},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLocation(time.UTC)
			if got := Location(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Location() = %v, want %v", got, tt.want)
			}
		})
	}
}
