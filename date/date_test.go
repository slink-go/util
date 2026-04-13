package date

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDateMarshalling(t *testing.T) {
	var v = struct {
		Title string `json:"title"`
		Date  Date   `json:"date"`
	}{
		Title: "test",
		Date:  Now(),
	}
	data, err := json.Marshal(v)
	fmt.Println(string(data), err)
}
