package types

import (
	"errors"
	"fmt"
	"testing"
)

func TestEnvs(t *testing.T) {
	jsonStr := `
    {
        "name":"liangyongxing",
        "age":"12"
    }
    `
	var envs Envs

	fmt.Println(envs)

	err := envs.Scan([]byte(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(envs)

	envs["abc"] = "def"

	value, err := envs.Value()
	if err != nil {
		t.Fatal(err)
	}

	bytes, ok := value.([]byte)
	if !ok {
		t.Fatal(errors.New("类型错误"))
	}
	fmt.Println(string(bytes))

}
