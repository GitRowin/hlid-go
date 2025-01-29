package hlid

import (
	"crypto/rand"
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ID [16]byte

var (
	_ fmt.Stringer               = (*ID)(nil)
	_ encoding.BinaryMarshaler   = (*ID)(nil)
	_ encoding.BinaryUnmarshaler = (*ID)(nil)
	_ encoding.TextMarshaler     = (*ID)(nil)
	_ encoding.TextUnmarshaler   = (*ID)(nil)
	_ driver.Valuer              = (*ID)(nil)
	_ sql.Scanner                = (*ID)(nil)
	_ json.Marshaler             = (*ID)(nil)
	_ json.Unmarshaler           = (*ID)(nil)
)

func NewWithTime(t time.Time) ID {
	var id ID

	// 100 microsecond resolution
	ts := t.UnixMicro() / 100

	// 6 bytes (48 bits) to store the timestamp
	// This won't overflow until 2861-12-16 05:21:11.0656 UTC
	id[0] = byte(ts >> 40)
	id[1] = byte(ts >> 32)
	id[2] = byte(ts >> 24)
	id[3] = byte(ts >> 16)
	id[4] = byte(ts >> 8)
	id[5] = byte(ts)

	// 10 bytes (80 bits) of cryptographically secure randomness
	if _, err := rand.Read(id[6:]); err != nil {
		panic(err)
	}

	return id
}

func New() ID {
	return NewWithTime(time.Now())
}

func (id ID) String() string {
	return hex.EncodeToString(id[:])
}

func (id ID) MarshalBinary() (b []byte, err error) {
	return append([]byte(nil), id[:]...), nil
}

func (id *ID) UnmarshalBinary(b []byte) error {
	if len(b) != len(id) {
		return fmt.Errorf("invalid ID length: %d", len(b))
	}

	copy(id[:], b)
	return nil
}

func (id ID) MarshalText() (b []byte, err error) {
	return []byte(hex.EncodeToString(id[:])), nil
}

func (id *ID) UnmarshalText(b []byte) error {
	v, err := hex.DecodeString(string(b))

	if err != nil {
		return err
	}

	return id.UnmarshalBinary(v)
}

func (id ID) Value() (driver.Value, error) {
	return id[:], nil
}

func (id *ID) Scan(value any) error {
	switch v := value.(type) {
	// The UUID type is returned as a string
	case string:
		return id.UnmarshalText([]byte(strings.ReplaceAll(v, "-", "")))
	case []byte:
		return id.UnmarshalBinary(v)
	default:
		return fmt.Errorf("unsupported ID type: %T", value)
	}
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return id.UnmarshalText([]byte(s))
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(id[:]))
}

func Parse(s string) (ID, error) {
	var id ID
	v, err := hex.DecodeString(s)

	if err != nil {
		return id, err
	}

	if err := id.UnmarshalBinary(v); err != nil {
		return id, err
	}

	return id, nil
}

func MustParse(s string) ID {
	id, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return id
}
