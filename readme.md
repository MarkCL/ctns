# ConvertToNewStruct 
    converts a struct to another struct with same tag
    if tagName is not passed into this function, it will convert T1 to T2 with forced transformation
    tagName accepts json or other tag name which will be used to match fields
    it will only take first tagName for conversion
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

type Test3 string

type Test4 string

func main() {
	t1 := &Test{Name: "test", Age: 10, Address: "address", Mobile: "mobile"}
	t2 := ctns.ConvertToNewType[*Test, *Test2](t1, "json")
	var t3 Test3 = "ttt"
	t4 := ctns.ConvertToNewType[Test3, Test4](t3)
	fmt.Println(t2)
	fmt.Println(t4)
}
```