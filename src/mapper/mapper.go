package mapper

import "encoding/json"

//generics came with go 1.18 version
func MarshalMap[To interface{}](from interface{}) *To {

	jsonM, err := json.Marshal(from)
	if err != nil {
		panic("error while serialization")
	}

	to := new(To)
	err = json.Unmarshal(jsonM, &to)
	if err != nil {
		panic("error while deserialization")
	}
	return to
}
