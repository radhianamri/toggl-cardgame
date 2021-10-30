package enums

import (
	"database/sql/driver"
	"errors"

	"github.com/radhianamri/toggl-cardgame/lib/log"
)

// Bool - custom datatype to read bit column types in database
type Bool bool

const (
	True  Bool = true
	False Bool = false
)

func (b Bool) Value() (driver.Value, error) {

	log.Debug("zxc")
	if b {
		return []byte{1}, nil
	}
	return []byte{0}, nil
}

func (b *Bool) Scan(src interface{}) error {
	log.Debug("asdsa")
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad Bool []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}
