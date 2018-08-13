package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type Dog struct {
	Name string
	Age  int
}

type Beast struct {
	Name string
}

func main() {
	dog := Dog{}
	beast := Beast{}
	payload := strings.NewReader(`{"name":"cricket","age":20}`)

	var buf bytes.Buffer
	body := io.TeeReader(payload, &buf)

	err := json.NewDecoder(body).Decode(&dog)
	check(err)
	err = json.NewDecoder(&buf).Decode(&beast)
	check(err)

	fmt.Printf("%#v\n", dog)
	fmt.Printf("%#v\n", beast)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
