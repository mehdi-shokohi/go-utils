package etc

import (
	"context"
	"fmt"
	"testing"
)

type PostBody struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func TestCallHttpClient(t *testing.T) {

	rep, err := SendHttpPost(context.Background(),nil, "https://jsonplaceholder.typicode.com/posts", PostBody{Title: "foo", Body: "bar", UserID: 1})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep.Body()))
}
