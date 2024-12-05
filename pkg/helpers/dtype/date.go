package dtype

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

const dateLayout = "2006-01-02"

// MarshalJSON formats the date as yyyy-mm-dd
func (d Date) MarshalJSON() ([]byte, error) {
	formatted := d.Format(dateLayout)
	return []byte(`"` + formatted + `"`), nil
}

// UnmarshalJSON parses yyyy-mm-dd into a Date
func (d *Date) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(`"`+dateLayout+`"`, string(data))
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// Value implements the driver.Valuer interface for database storage
func (d Date) Value() (driver.Value, error) {
	return d.Format(dateLayout), nil
}

// Scan implements the sql.Scanner interface for database retrieval
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		d.Time = v
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}
