package ec2

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"strconv"
	"strings"
)

// func memoryConv(m string) (f float64) {
// 	s := strings.Split(m, " ")
// 	f, _ = strconv.ParseFloat(s[0], 64)
// 	return f
// }

// UnmarshalJSON : comment
func (r *Ec2Attributes) UnmarshalJSON(data []byte) error {
	type Alias Ec2Attributes
	aux := &struct {
		Vcpu   string `json:"vcpu"`
		Memory string `json:"memory"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	// RAM
	a := strings.Split(aux.Memory, " ")
	if len(a) < 2 {
		r.Memory.Value = 0
		r.Memory.Unit = "GiB"
		return nil
	}
	val, err := strconv.ParseFloat(a[0], 64)
	if err != nil {
		return err
	}

	r.Memory.Value = val
	r.Memory.Unit = a[1]

	// vCPU
	r.Vcpu, err = strconv.Atoi(aux.Vcpu)
	if err != nil {
		return err
	}
	return nil
}
func (m RAM) MarshalJSON() ([]byte, error) {
	return json.Marshal(convertRAM(m))
}

func convertRAM(r RAM) string {
	s := fmt.Sprintf("%.2f", r.Value)
	return s + r.Unit
}
