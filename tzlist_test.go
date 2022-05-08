package tzlist

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetTimezones(t *testing.T) {
	zones := GetTimezones()
	expectedLen := 338
	if len(zones) != expectedLen {
		t.Fatalf("unexpected timezone count, expected %v was %v", expectedLen, len(zones))
	}
	assertEqual(t, "Europe/Andorra", zones[0])
	assertEqual(t, "Africa/Johannesburg", zones[expectedLen-1])
	for i, zone := range zones {
		if !strings.Contains(zone, "/") {
			t.Errorf("zone at %v does not look like a zone: %v", i, zone)
		}
	}
}

func TestGetRecords(t *testing.T) {
	records := GetRecords()
	expectedLen := 338
	if len(records) != expectedLen {
		t.Fatalf("unexpected record count, expected %v was %v", expectedLen, len(records))
	}
	expectedFirst := Record{
		CountryCodes: []string{"AD"},
		Coordinates:  "+4230+00131",
		TZ:           "Europe/Andorra",
	}
	expectedLast := Record{
		CountryCodes: []string{"ZA", "LS", "SZ"},
		Coordinates:  "-2615+02800",
		TZ:           "Africa/Johannesburg",
	}
	assertEqual(t, expectedFirst, records[0])
	assertEqual(t, expectedLast, records[expectedLen-1])
	for i, r := range records {
		if !strings.Contains(r.TZ, "/") {
			t.Errorf("zone looks odd for record at %v: %v", i, r)
		}
		if len(r.CountryCodes) == 0 {
			t.Errorf("no country codes for record at %v: %v", i, r)
		}
		if r.Coordinates == "" {
			t.Errorf("no coordinates for record at %v: %v", i, r)
		}
	}
}

func TestIsValidTz(t *testing.T) {
	validTz := "America/Rio_Branco"
	invalidTz := "foo"
	if !IsValidTZ(validTz) {
		t.Errorf("expected %v to be valid tz", validTz)
	}
	if IsValidTZ(invalidTz) {
		t.Errorf("expected %v to be invalid tz", invalidTz)
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	t.Helper()
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Logf("expected %v but received %v", expected, actual)
		t.Fatal(diff)
	}
}
