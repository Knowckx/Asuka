package encoding

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// go binary encoder
func ToGOB64(in Person) string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(in)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

// go binary decoder
func FromGOB64(str string) Person {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(`failed base64 Decode`, err)
	}
	b := bytes.NewBuffer(by)
	d := gob.NewDecoder(b)
	out := Person{}

	err = d.Decode(&out)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
	return out
}
