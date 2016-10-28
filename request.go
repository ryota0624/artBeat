package artBeat

import (
  "io/ioutil"
  "net/http"
  "net/url"
)

type ArtBeatArgs struct  {
  Latitude float64
  Longitude float64
  MaxResult int
}

func NewArtBeatArgs(latitude float64, longitude float64) ArtBeatArgs {
  return ArtBeatArgs{
    Latitude: latitude,
    Longitude: longitude,
    MaxResult: 5,
  }
}

type Response struct {
  Event string
}

const baseUrl = "http://www.tokyoartbeat.com/list/event_searchNear"

func makeUrl(param ArtBeatArgs) string {
  u, err := url.Parse(baseUrl)
  if err != nil {
    return ""
  }

  q := u.Query()
  q.Set("Latitude", formatFtoStr(param.Latitude))
  q.Set("Longitude", formatFtoStr(param.Longitude))
  u.RawQuery = q.Encode()
  return u.String()
  //strLatitude := formatFtoStr(param.Latitude)
  //strLongitude := formatFtoStr(param.Longitude)
  //return url + "?Latitude=" + strLatitude + "&Longitude=" + strLongitude + "&MaxResult=" + strconv.Itoa(param.MaxResult)
}

func GetArtBeat(args ArtBeatArgs) string {
  response, err := http.Get(makeUrl(args))
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
