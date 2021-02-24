// Code generated by "enumer -trimprefix=Level -transform screaming-snake -type=Level"; DO NOT EDIT.

//
package logging

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _LevelName = "DEBUGINFOWARNERRORFATAL"

var _LevelIndex = [...]uint8{0, 5, 9, 13, 18, 23}

func (i Level) String() string {
	i -= 1
	if i < 0 || i >= Level(len(_LevelIndex)-1) {
		return fmt.Sprintf("Level(%d)", i+1)
	}
	return _LevelName[_LevelIndex[i]:_LevelIndex[i+1]]
}

var _LevelValues = []Level{1, 2, 3, 4, 5}

var _LevelNames = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

var _LevelNameToValueMap = map[string]Level{
	_LevelName[0:5]:   1,
	_LevelName[5:9]:   2,
	_LevelName[9:13]:  3,
	_LevelName[13:18]: 4,
	_LevelName[18:23]: 5,
}

// LevelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func LevelString(s string) (Level, error) {

	if val, ok := _LevelNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Level values", s)
}

func ParseLevel(s string) (Level, error) {
	return LevelString(s)
}

// LevelValues returns all values of the enum
func LevelValues() []Level {
	return _LevelValues
}

func LevelNames() []string {
	return _LevelNames
}

// IsALevel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Level) IsALevel() bool {
	for _, v := range _LevelValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Level
func (i Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Level
func (i *Level) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Level should be a string, got %s", data)
	}

	var err error
	*i, err = LevelString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for Level
func (i Level) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Level
func (i *Level) UnmarshalText(text []byte) error {
	var err error
	*i, err = LevelString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for Level
func (i Level) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Level
func (i *Level) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = LevelString(s)
	return err
}

func (i Level) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Level) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := LevelString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
