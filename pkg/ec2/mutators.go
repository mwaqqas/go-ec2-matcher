package ec2

import (
	"encoding/json"
	// "fmt"
	"strconv"
	"strings"
)

// func memoryConv(m string) (f float64) {
// 	s := strings.Split(m, " ")
// 	f, _ = strconv.ParseFloat(s[0], 64)
// 	return f
// }

// func (r *ec2Attributes) UnmarshalJSON() (b []byte, err error) {
// 	err = json.Unmarshal(b, &r)
// 	if err != nil {
// 		return
// 	}
// 	strMemValue := r.Memory
// 	m := strings.Split(strMemValue, " ")
// 	// r.Memory, err = strconv.ParseFloat(m[0], 64)
// 	r.Memory = m[0]
// 	fmt.Println(r.Memory)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

func (r *ec2Attributes) UnmarshalJSON(data []byte) error {
	type Alias ec2Attributes
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
		r.Memory.value = 0
		r.Memory.unit = "MiB"
		return nil
	}
	val, err := strconv.ParseFloat(a[0], 64)
	if err != nil {
		return err
	}

	r.Memory.value = val
	r.Memory.unit = a[1]
	if err != nil {
		return err
	}

	// vCPU
	r.Vcpu, err = strconv.Atoi(aux.Vcpu)
	if err != nil {
		return err
	}
	return nil
}
