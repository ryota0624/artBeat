package artBeat

import (
  "testing"
)

func TestMakeUrl(t *testing.T) {
  param := NewArtBeatArgs(35.671208, 139.76517)
  url := makeUrl(param)
  if url != "http://www.tokyoartbeat.com/list/event_searchNear?Latitude=35.671208&Longitude=139.76517" {
    t.Errorf("url notmath")
  }
}

func TestParseArtBeatXml(t *testing.T) {
  artBeat := NewArtBeatArgs(35.671208, 139.76517)
}