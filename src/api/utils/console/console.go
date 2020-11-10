package console

import (
	"fmt"
	"log"
	"auto"
)

func Pretty(data, interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}