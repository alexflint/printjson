package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

import "io/ioutil"

func fail(msg interface{}, parts ...interface{}) {
	fmt.Printf(fmt.Sprintf("%v\n", msg), parts...)
	os.Exit(1)
}

func main() {
	if len(os.Args) > 2 {
		fail("usage: printjson [FILE]")
	}
	if len(os.Args) > 1 && os.Args[1] == "--help" || os.Args[1] == "-help" {
		fail("usage: printjson [FILE]")
	}

	var err error
	var buf []byte
	if len(os.Args) < 2 || os.Args[1] == "-" {
		buf, err = ioutil.ReadAll(os.Stdin)
	} else {
		buf, err = ioutil.ReadFile(os.Args[1])
	}
	if err != nil {
		fail(err)
	}

	var b bytes.Buffer
	err = json.Indent(&b, buf, "", "  ")
	if err != nil {
		fail(err)
	}
	b.WriteTo(os.Stdout)
}
