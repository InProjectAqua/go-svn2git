package main

import (
	"runtime"
	"fmt"
)

func default_routina_qty() int64 {
	var rootina_count int64=60000

	if runtime.GOOS=="darwin" {
		rootina_count =5000
	}

	if runtime.GOARCH=="mips"{
		rootina_count =100
	}

	fmt.Println("[default_routina_qty] QTY", rootina_count)
	return rootina_count

}
