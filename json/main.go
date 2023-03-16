package main

import (
	"encoding/json"
	"fmt"
)

type NFTModel struct {
	ID          string           `json:"token_id"`
	Description string           `json:"description"`
	Name        string           `json:"name"`
	Image       string           `json:"image"`
	Avatar      string           `json:"avatar"`
	Attributes  []*NFTAttributes `json:"properties"`
}

type NFTAttributes struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

func main() {
	Zhang3 := NFTModel{
		ID:          "张三",
		Description: "18",
		Name:        "18",
		Image:       "18",
		Avatar:      "18",
		Attributes: []*NFTAttributes{
			{TraitType: "renmin south road",
				Value: "123"},
		},
	}

	InfoOfZhang3, err := json.Marshal(Zhang3)
	if err == nil {
		fmt.Println(string(InfoOfZhang3))
	} else {
		fmt.Println(err)
		return
	}

}
