package datamodels

import (
	"fmt"
	"time"
)

type LocalTime time.Time

// MarshalJSON satify the json marshal interface
func (l LocalTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(l).UnixNano()/1e6)
	return []byte(stamp), nil
}

type LocalDate time.Time

// MarshalJSON satify the json marshal interface
func (l LocalDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(l).Format("2006-01-02"))
	return []byte(stamp), nil
}
