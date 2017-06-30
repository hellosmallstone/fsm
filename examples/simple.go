// +build ignore

package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "closing", Src: []string{"open"}, Dst: "closing"},
			{Name: "close", Src: []string{"open", "closing"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

    //closed ----> open ----> closing ----> closed
    fmt.Println("example 1: closed ----> open ----> closing ----> closed")
	fmt.Println(fsm.Current())

	err := fsm.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event("closing")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event("close")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	//closed ----> open ---->closed
	fmt.Println("example 2: closed ----> open ----> closed")
	fmt.Println(fsm.Current())

	err = fsm.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event("close")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())
}
