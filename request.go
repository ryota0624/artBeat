package artBeat

import (
  "io/ioutil"
  "net/http"
  "strconv"
)
type ArtBeat struct  {
  Latitude float64
  Longitude float64
  MaxResult int
}

func NewArtBeat(latitude float64, longitude float64) ArtBeat {
  return ArtBeat{
    Latitude: latitude,
    Longitude: longitude,
    MaxResult: 5,
  }
}

type Response struct {
  Event string
}

const url = "http://www.tokyoartbeat.com/list/event_searchNear"

func makeUrl(param ArtBeat) string {
  strLatitude := formatFtoStr(param.Latitude)
  strLongitude := formatFtoStr(param.Longitude)
  return url + "?Latitude=" + strLatitude + "&Longitude=" + strLongitude + "&MaxResult=" + strconv.Itoa(param.MaxResult)
}

func (beat ArtBeat) Get() string {
  response, err := http.Get(makeUrl(beat))
  if err != nil {
    print(err)
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    print(err)
  }
  return string(body)
}