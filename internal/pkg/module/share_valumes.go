package module

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ShareVolume struct {
	Name      string `json:"name"`
	Type      string `json:"type"` //nfs,pvc
	NfsPath   string `json:"nfs_path"`
	NfsServer string `json:"nfs_server"`
	PvcName   string `json:"pvc_name"`
}

func (ShareVolume) GormDataType() string {
	return "json"
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (v *ShareVolume) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	err := json.Unmarshal(bytes, &v)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (v ShareVolume) Value() (driver.Value, error) {

	jsonStr, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(jsonStr).MarshalJSON()
}
