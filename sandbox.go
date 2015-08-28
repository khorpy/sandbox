package main

import (
	"encoding/json"
	"fmt"
)

type example struct {
	X int
	S string
}

func main() {
	v := example{10, "Hello!"}
	v_enc, _ := json.Marshal(v)
	fmt.Println(string(v_enc))
}
