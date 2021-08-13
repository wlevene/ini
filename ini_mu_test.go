package ini

import (
	"fmt"
	"testing"
)

type TestConfig struct {
	K    string  `ini:"k" json:"k,omitempty"`
	K1   int     `ini:"k1" json:"k1,omitempty"`
	K2   float64 `json:"k2,omitempty" ini:"k2"`
	K3   int64   `ini:"k3"`
	User User    `ini:"user"`
}

type User struct {
	Name string `ini:"name"`
	Age  int    `ini:"age"`
}

func TestIniMu(t *testing.T) {
	doc := `
k=v
k1=2
k2=2.2
k3=3

[user]
name=tom
age=-23
`

	cfg := TestConfig{}

	Unmarshal([]byte(doc), &cfg)
	fmt.Println("cfg:", cfg)

}
