package model
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"database/sql/driver"
)
type DeleteNullTime sql.NullTime

// Scan implements the Scanner interface.
func (n *DeleteNullTime) Scan(value interface{}) error {
	return (*sql.NullTime)(n).Scan(value)
}

// Value implements the driver Valuer interface.
func (n DeleteNullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

func (n DeleteNullTime) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Time)
	}
	return json.Marshal(nil)
}

func (n *DeleteNullTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Time)
	if err == nil {
		n.Valid = true
	}
	return err
}

type StatusNullString sql.NullString

// Scan implements the Scanner interface.
func (s *StatusNullString) Scan(value interface{}) error {
	return (*sql.NullString)(s).Scan(value)
}



