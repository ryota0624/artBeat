package artBeat

import "strconv"

func formatFtoStr(fl float64) string {
  return strconv.FormatFloat(fl, 'f', 4, 64)
}