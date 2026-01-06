package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`        //fungsi json tag untuk mengubah nama field pada JSON
	Name     string `json:"name"`      //fungsi json tag untuk mengubah nama field pada JSON
	ImageURL string `json:"image_url"` //fungsi json tag untuk mengubah nama field pada JSON
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0014",
		Name:     "Apple Mac Book Pro",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id": "P0014", "name": "Apple Mac Book Pro", "image_url": "http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
