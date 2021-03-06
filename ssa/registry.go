package ssa

import "fmt"

type registry map[string]Pass
var Registry registry

func GetPass(name string) Pass {
	var p Pass
	var ok bool
	if p, ok = Registry[name]; !ok {
		panic("no such pass: "+name)
	}
	return p
}

func RegisterPass(name string, p Pass) {
	fmt.Println("registered pass:", name)
	Registry[name] = p
}

func init() {
	Registry = make(registry)
}
