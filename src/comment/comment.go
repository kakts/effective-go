/*
  comment パッケージ

  コメントに関する機能を提供する
*/
package comment

type Comment struct {
	Text string
}

func GetComment() Comment {
	return Comment{Text: "Hello comment."}

}
