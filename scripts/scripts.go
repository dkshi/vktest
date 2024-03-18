package scripts

import (
	"fmt"
	"math"
	"time"
)

func StringToDate(strDate string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, strDate)
	return t, err

}

func СonvertFloat64ToInt(floatNumber float64) (int, error) {
	if floatNumber < math.MinInt64 || floatNumber > math.MaxInt64 {
		return 0, fmt.Errorf("float64 value %f is too large to be converted to int", floatNumber)
	}
	return int(floatNumber), nil
}
