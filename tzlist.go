package tzlist

import (
	_ "embed"
	"encoding/csv"
	"strings"
)

//go:embed data/zone1970.tab
var zone1970 string

type Record struct {
	CountryCodes []string
	Coordinates  string
	TZ           string
}

var cachedRecords []Record
var cachedZones []string

func parseZone1970() {
	zoneReader := strings.NewReader(zone1970)
	csvReader := csv.NewReader(zoneReader)
	csvReader.Comma = '\t'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1
	records, err := csvReader.ReadAll()
	// Note on the panics: dear user, do not worry.
	// The data being parsed is unchanging due to the file embedding, and so any errors will discovered by the unit tests.
	if err != nil {
		panic("BUG: csv (tabbed) parser cannot parse embedded data: " + err.Error())
	}
	for _, tabRecord := range records {
		if len(tabRecord) < 3 {
			panic("BUG: unexpectedly short record")
		}
		record := Record{
			CountryCodes: strings.Split(tabRecord[0], ","),
			Coordinates:  tabRecord[1],
			TZ:           tabRecord[2],
		}
		cachedRecords = append(cachedRecords, record)
		cachedZones = append(cachedZones, record.TZ)
	}
}

func GetRecords() []Record {
	if cachedRecords == nil {
		parseZone1970()
	}

	return cachedRecords
}

func GetTimezones() []string {
	if cachedZones == nil {
		parseZone1970()
	}

	return cachedZones
}
