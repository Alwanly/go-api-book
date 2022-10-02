package book

type CreateModel struct {
	Email    string
	Title    string `json:"title" binding:"required"`
	Headline string `json:"head_line" binding:"required"`
	Tag      string `json:"tag" binding:"required"`
}

type GetAllModel struct {
	Email   string
	Keyword string `form:"keyword"`
	Page    int    `form:"page" binding:"omitempty,numeric,min=1"`
	Limit   int    `form:"limit" binding:"omitempty,numeric,min=1"`
}

type GetModel struct {
	Email string
	ID    int `uri:"id" binding:"required,numeric"`
}
