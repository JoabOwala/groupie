package data

import (
	"testing"
)

func TestGetArtists(t *testing.T) {
	data, err := GetArtists()
	if err != nil && data != nil {
		t.Fatalf(`GetArtists() = %v, want match for %v, nil`, len(data), nil)
	}
	want := true

	res := len(data) >= 1
	if !res {
		t.Fatalf(`GetArtists() = %v, want match for %v, nil`, res, want)
	}
}
