package main

import (
	"fmt"
	"syscall"

	"github.com/cilium/ebpf"
)

func main() {
	fmt.Println("hell world")

	err := syscall.Setuid(syscall.Getuid())
	if err != nil {
		panic(err)
	}

	spec, err := ebpf.LoadCollectionSpec("ebpf/mybpf.o")
	if err != nil {
		panic(err)
	}

	var objs struct {
		XCProg *ebpf.Program `ebpf:"xdp_xconnect"`
		XCMap  *ebpf.Map     `ebpf:"xconnect_map"`
	}
	if err := spec.LoadAndAssign(&objs, nil); err != nil {
		panic(err)
	}
	defer objs.XCProg.Close()
	defer objs.XCMap.Close()
}
