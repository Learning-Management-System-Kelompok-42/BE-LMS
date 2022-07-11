package timezone

import (
	"time"

	"github.com/labstack/gommon/log"
)

var countryTimezone = map[string]string{
	"Jakarta":  "Asia/Jakarta",
	"Makassar": "Asia/Makassar",
	"Jayapura": "Asia/Jayapura",
}

func ConvertToTZ(country string, now string) time.Time {
	location, err := time.LoadLocation(countryTimezone[country])
	if err != nil {
		log.Error(err)
	}

	return time.Now().In(location)
}
