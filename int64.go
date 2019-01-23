package nulltype

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// NullInt64 is null friendly type for int64.
type NullInt64 struct {
	i sql.NullInt64
}

// Valid return the value is valid. If true, it is not null value.
func (i *NullInt64) Valid() bool {
	return i.i.Valid
}

// Value return the value.
func (i *NullInt64) Value() int64 {
	return i.i.Int64
}

// Reset set nil to the value.
func (i *NullInt64) Reset() {
	i.i.Int64 = 0
	i.i.Valid = false
}

// Set set the value.
func (i *NullInt64) Set(value int64) {
	i.i.Valid = true
	i.i.Int64 = value
}

// Scan is a method for database/sql.
func (i *NullInt64) Scan(value interface{}) error {
	return i.i.Scan(value)
}

// String return string indicated the value.
func (i NullInt64) String() string {
	if !i.i.Valid {
		return "nil"
	}
	return fmt.Sprint(i.i.Int64)
}

// MarshalJSON encode the value to JSON.
func (i NullInt64) MarshalJSON() ([]byte, error) {
	if !i.i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.i.Int64)
}

// UnmarshalJSON decode data to the value.
func (i *NullInt64) UnmarshalJSON(data []byte) error {
	var value *int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	i.i.Valid = value != nil
	if value == nil {
		i.i.Int64 = 0
	} else {
		i.i.Int64 = *value
	}
	return nil
}
