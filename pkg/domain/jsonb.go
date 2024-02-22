package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONB map[string]interface{}

func NewJsonB() JSONB {
	return make(JSONB)
}

// Value transforms the type to database driver compatible type.
func (j JSONB) Value() (driver.Value, error) {
	v, err := json.Marshal(j)
	return v, err
}

// Scan take the raw data that comes from database and convert it as a Go type.
// The reverse process of Value method.
func (j *JSONB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return nil
	}

	*j, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interfaceP{}) failed")
	}

	return nil
}
