package customtypes

import (
	"time"
)

type Date struct {
	time.Time
}

const DateFormat = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	str := string(b)
	// Strip the quotes
	str = str[1 : len(str)-1]
	t, err := time.Parse(DateFormat, str)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Format(DateFormat) + `"`), nil
}

func (d *Date) String() string {
	return d.Format(DateFormat)
}
