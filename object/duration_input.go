package object

import (
	"encoding/json"
	"time"
)

// DurationInput is a user input duration
type DurationInput int64

// GetSeconds returns the seconds value of the duration input
func (d DurationInput) GetSeconds() int64 {
	intVal := int64(d)
	if intVal != 0 && time.Duration(intVal) > ((time.Hour*24)*3650) {
		return intVal / int64(time.Second)
	}
	return int64(d)
}

// GetDuration returns a time.Duration for the user input
func (d DurationInput) GetDuration() time.Duration {
	intVal := d
	if intVal != 0 && int64(intVal) < int64(time.Second) {
		return time.Second * time.Duration(intVal)
	}
	return time.Duration(d)
}

// GetRawValue returns the raw user input
func (d DurationInput) GetRawValue() int64 {
	return int64(d)
}

// MarshalJSON always returns seconds
func (d *DurationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.GetSeconds)
}
