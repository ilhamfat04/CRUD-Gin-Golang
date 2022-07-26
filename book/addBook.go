package book

// struct for store data post
// no casesensitive for  variable name when submit form JSON postman
// 		subTitle == SubTitle, etc
type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `binding:"required,number"`
	// SubTitle string `json:"sub_title"`  // when different var name and json form
	SubTitle string
}
