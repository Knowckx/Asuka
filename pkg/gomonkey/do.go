package gomonkey

import "fmt"

type Tom struct {
	name string
}

// GetName 就是希望被打桩的目标
func (this *Tom) GetName() (string, error) {
	output := fmt.Sprintf("My Name Is %s", this.name)
	return output, nil
}

type Peopel interface {
	GetName() (string, error)
}

func NewPeopel(name string) Peopel {
	out := &Tom{
		name: "tom1",
	}
	return out
}
