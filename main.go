// main package
package main

import (
	"fmt"

	"github.com/kakts/effective-go/src/comment"
	"github.com/kakts/effective-go/src/data"
)

func main() {
	comment := comment.GetComment()
	fmt.Println(comment.Text)

	data.InstanceTest()
}
