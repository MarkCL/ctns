# ConvertToNewStruct 
    converts a struct to another struct with same tag
    tagName accepts json or other tag name which will be used to match fields
    a simple main.go example for using this function:
```
package main

import (
	"fmt"
	"github.com/MarkCL/ctns"
)

type Test struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Mobile  string `json:"mobile"`
}

type Test2 struct {
	Ss string `json:"name"`
	Pp int    `json:"age"`
	Ad string `json:"address"`
}

func main() {
	t1 := Test{Name: "test", Age: 10, Address: "address", Mobile: "mobile"}
	t2 := ctns.ConvertToNewStruct[Test, *Test2](t1, "json")
	fmt.Println(t2)
}
```