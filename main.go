package main

import (
	"fmt"

	"github.com/kakts/effective-go/src/comment"
)

func main() {
	comment := comment.GetComment()
	fmt.Println(comment.Text)
}
