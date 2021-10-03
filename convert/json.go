package convert

import (
	"encoding/json"
)

func (v Value) Json() ([]byte, error) {
	return json.Marshal(v.r)
}
