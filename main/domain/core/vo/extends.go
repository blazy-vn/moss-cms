package vo

import (
	"database/sql/driver"
	"encoding/json"
)

// Extends additional extension value object, used for article details expansion
type Extends []ExtendsItem

func (ext Extends) Get(key string) any {
	for _, item := range ext {
		if item.Key == key {
			return item.Value
		}
	}
	return nil
}

func (ext *Extends) Scan(value interface{}) error {
	s, _ := value.(string)
	if len(s) == 0 {
		*ext = Extends{}
		return nil
	}
	_ = json.Unmarshal([]byte(s), ext)
	return nil
}

func (ext Extends) Value() (driver.Value, error) {
	b, err := json.Marshal(&ext)
	if err != nil || len(b) == 0 {
		return "", err
	}
	return string(b), nil
}

type ExtendsItem struct {
	Key   string           `json:"key"`
	Value ExtendsItemValue `json:"value"`
}

type ExtendsItemValue any
