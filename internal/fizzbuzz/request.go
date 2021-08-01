package fizzbuzz

type postRequest struct {
	Int1  int    `json:"int1" binding:"required,gte=1"`
	Int2  int    `json:"int2" binding:"required,gte=1"`
	Limit int    `json:"limit" binding:"required,gte=1,lt=100000"`
	Str1  string `json:"str1" binding:"required,min=1,max=100"`
	Str2  string `json:"str2" binding:"required,min=1,max=100"`
}
