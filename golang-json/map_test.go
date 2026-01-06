package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"928", "name":"Rafi Fajar Nanda", "age":24}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{} // gunakan map untuk menampung data JSON yang tidak diketahui strukturnya
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["age"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "928",
		"name":  "Laptop RoG Strix",
		"price": 10000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
