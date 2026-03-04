package main

import (
	"encoding/json"
	"fmt"
)

// OptionalPointerFields demonstrate pattern:
// pointers (*string, *int, *bool) for fields,
// which can be nil (or not passed) or could have a value

// Why do we need it?
// in Golang string always "", int alwayer 0
// We can't use "not passed" of use "0 / empty string"
// Pointer gives us two options nil | &value

type UserUpdateRequest struct {
	Name *string `json:"name,omitempty`
	Age *int `json:"age,omitempty`
	Email *string `json:"email,omitempty`
}

type PersonExample struct {
	Name *string `json:"name"`
}

// Helpers for comfortable creating porinters (good to have in tests of API)
func ptrString(s string) *string {return &s}
func ptrInt(i int) *int { return &i}

func main() {
	// Exmaple 1: PATCH - try to update only Name, other things don't change
	req1 := UserUpdateRequest{
		Name: ptrString("Alice"),
	}
	applyUpdate(req1)

	// Example 2: Set Age to 0, with nil it means that we don't need to change it
	req2 := UserUpdateRequest{
		Age: ptrInt(0),
	}
	applyUpdate(req2)

	// Example 3: JSON with null and no fiels
	jsonStr := `{"name": "Bob", "age": null}`

	var req3 UserUpdateRequest
	_ = json.Unmarshal([]byte(jsonStr), &req3)

	fmt.Printf("Name: %v, Age: %v, Email: %v\n",
	req3.Name != nil, req3.Age != nil, req3.Email != nil,
	// Name: true, Age: false (null → nil), Email: false (absent → nil)
	)
	if req3.Name != nil {
		fmt.Println("Name =", *req3.Name) // Bob
	}

	// Example 4: Marshal - pointer while marshalling use value, not address
	// encodin/json deferences pointers automatically and use value by address

	name := "Alice"
	p := PersonExample{Name: &name}
	data, _ := json.Marshal(p)
	fmt.Println(string(data)) // {"name": "Alice"} -> value not address


}



func applyUpdate(req UserUpdateRequest) {
	if req.Name != nil {
		fmt.Println("Updating name to", *req.Name)
	}
	if req.Age != nil {
		fmt.Println("Updating age to", *req.Age)
	}
	if req.Email != nil {
		fmt.Println("Updating email to", *req.Email)
	}
}
