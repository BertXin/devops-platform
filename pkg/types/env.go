package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Envs map[string]string

func (Envs) GormDataType() string {
	return "json"
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (envs *Envs) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	err := json.Unmarshal(bytes, &envs)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (envs *Envs) Value() (driver.Value, error) {

	jsonStr, err := json.Marshal(envs)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(jsonStr).MarshalJSON()
}
