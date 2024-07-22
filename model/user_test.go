package model

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestUserJson(t *testing.T) {
	jsonString := `{
	   "id": 1,
	   "name": "John Doe",
	   "hero": {
	       "heroId": 1,
	       "heroName": "Superman",
	       "heroAttr": [
	           {
	               "attrId": 1,
	               "value": 100
	           },
	           {
	               "attrId": 2,
	               "value": 200
	           }
	       ]
	   },
	   "props": {
	       "1": {
	           "propId": 101,
	           "propNum": 2
	       },
	       "2": {
	           "propId": 102,
	           "propNum": 5
	       }
	   },
	   "normalInt": [1, 2, 3, 4, 5]
	}`
	user := NewUser()
	err := json.Unmarshal([]byte(jsonString), user)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	// 打印解析后的结构体
	fmt.Printf("Parsed User: %+v\n", user)
}
