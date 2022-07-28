package book

import "encoding/json"

// struct for store data post
// no casesensitive for  variable name when submit form JSON postman
// 		subTitle == SubTitle, etc
type BookInputUpdate struct {
	Title string `json:"title"`
	Price json.Number
	// SubTitle string `json:"sub_title"`  // when different var name and json form
	Description string
}
