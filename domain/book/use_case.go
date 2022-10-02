package book

import (
	"books-api/domain/users"
	"books-api/infrastructure/authentication"
	"books-api/infrastructure/utils"
	"context"
	"errors"
	"math"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookUseCase interface {
	Create(c context.Context, model CreateModel) utils.BaseResponse
	GetAll(c context.Context, model GetAllModel) utils.BasePagination
	Get(c context.Context, model GetModel) utils.BaseResponse
}

type BookUseCaseImpl struct {
	Database *gorm.DB
	JwtAuth  authentication.IJwtAuth
}

// Get implements BookUseCase
func (u *BookUseCaseImpl) Get(c context.Context, model GetModel) utils.BaseResponse {
	db := u.Database.WithContext(c)

	var author users.Users

	db.First(&author, model.Email)

	var book Books

	result := db.First(&book, model.ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utils.WrapperReponse(http.StatusNotFound, "Book is not found", nil)
	}

	if book.Author != author.Name {
		return utils.WrapperReponse(http.StatusForbidden, "User not allow", nil)
	}

	response := BooksDto{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Tag:       book.Tag,
		CreatedAt: book.CreatedAt,
		UpdateAt:  book.UpdatedAt,
	}
	return utils.WrapperReponse(http.StatusOK, "Book", response)
}

func (u *BookUseCaseImpl) Create(c context.Context, model CreateModel) utils.BaseResponse {
	db := u.Database.WithContext(c)

	var author users.Users
	db.First(&author, model.Email)

	book := Books{
		Title:    model.Title,
		Author:   author.Name,
		Headline: model.Headline,
		Tag:      model.Tag,
	}
	result := db.Create(&book)

	if result.Error != nil {
		println(result.Error)
		return utils.WrapperReponse(http.StatusInternalServerError, "Failed to save book", nil)
	}

	response := BooksDto{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Tag:       book.Tag,
		CreatedAt: book.CreatedAt,
		UpdateAt:  book.UpdatedAt,
	}
	return utils.WrapperReponse(http.StatusCreated, "Success save book", response)
}

func (u *BookUseCaseImpl) GetAll(c context.Context, model GetAllModel) utils.BasePagination {
	db := u.Database.WithContext(c)

	var author users.Users

	db.First(&author, model.Email)

	var books []Books

	searchParams := db.Preload(clause.Associations)

	searchParams = searchParams.Scopes(utils.Pagination(model.Page, model.Limit))

	if len(model.Keyword) > 0 {
		searchParams = searchParams.Scopes(utils.Like([]string{"title", "headline", "tag"}, model.Keyword))
	}
	searchParams.Find(&books, "author=?", author.Name)

	var total int64

	countParams := db.Model(&Books{}).Where("author=?", author.Name)

	if len(model.Keyword) > 0 {

		countParams = searchParams.Scopes(utils.Like([]string{"title", "headline", "tag"}, model.Keyword))
	}

	countResult := countParams.Count(&total)

	if countResult.Error == nil {
		return utils.WrapperPaginate(http.StatusInternalServerError, "Can't get books", nil, nil)
	}

	rows := []BooksItemDto{}

	for _, book := range books {
		rows = append(rows, BooksItemDto{
			ID:       book.ID,
			Title:    book.Title,
			Headline: book.Headline,
			Author:   book.Author,
			Tag:      book.Tag,
		})
	}

	meta := map[string]interface{}{
		"currentPage":     model.Page,
		"totalPage":       math.Ceil(float64(total) / float64(model.Limit)),
		"totalData":       total,
		"totalDataOnPage": len(rows),
	}

	return utils.WrapperPaginate(http.StatusOK, "Ok", meta, rows)
}

func ConstructBookUseCase(db *gorm.DB, jwtAuth authentication.IJwtAuth) BookUseCase {
	return &BookUseCaseImpl{
		Database: db,
		JwtAuth:  jwtAuth,
	}
}
