package binding_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Alwanly/go-codebase/pkg/binding"
	"github.com/Alwanly/go-codebase/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type TestModel struct {
	AuthUserData *middleware.AuthUserData
	Field1       string `json:"field1"`
	Field2       string `query:"field2"`
	Field3       bool   `query:"field3" validate:"required,min=1"`
}

func TestBindModel_Success(t *testing.T) {
	app := fiber.New()
	log, _ := zap.NewDevelopment()

	app.Post("/test", func(c *fiber.Ctx) error {
		model := new(TestModel)
		err := binding.BindModel(log, c, model, binding.BindFromBody(), binding.BindFromQuery())
		assert.NoError(t, err)
		assert.Equal(t, "value1", model.Field1)
		assert.Equal(t, "value2", model.Field2)
		return c.SendStatus(http.StatusOK)
	})

	req := httptest.NewRequest("POST", "/test?field2=value2", strings.NewReader(`{"field1":"value1"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestBindModel_BindFromBodyError(t *testing.T) {
	app := fiber.New()
	log, _ := zap.NewDevelopment()

	app.Post("/test", func(c *fiber.Ctx) error {
		model := new(TestModel)
		err := binding.BindModel(log, c, model, binding.BindFromBody())
		assert.Error(t, err)
		modelBindingErr, ok := err.(*binding.ModelBindingError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, modelBindingErr.Code)
		return c.SendStatus(http.StatusBadRequest)
	})

	req := httptest.NewRequest("POST", "/test", strings.NewReader(`invalid json`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestBindModel_BindFromQueryError(t *testing.T) {
	app := fiber.New()
	log, _ := zap.NewDevelopment()

	app.Post("/test", func(c *fiber.Ctx) error {
		model := new(TestModel)
		err := binding.BindModel(log, c, model, binding.BindFromQuery())
		if err != nil {
			modelBindingErr, ok := err.(*binding.ModelBindingError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusBadRequest, modelBindingErr.Code)
			return c.Status(http.StatusBadRequest).JSON(modelBindingErr.ResponseBody)
		}
		return c.SendStatus(http.StatusOK)
	})

	req := httptest.NewRequest("POST", "/test?field3=a", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestBindModel_AuthUserData(t *testing.T) {
	app := fiber.New()
	log, _ := zap.NewDevelopment()

	app.Post("/test", func(c *fiber.Ctx) error {
		authUser := &middleware.AuthUserData{UserID: "123"}
		c.Locals(middleware.LocalTokenKey, authUser)

		model := new(TestModel)
		err := binding.BindModel(log, c, model)
		assert.NoError(t, err)
		assert.Equal(t, authUser, model.AuthUserData)
		return c.SendStatus(http.StatusOK)
	})

	req := httptest.NewRequest("POST", "/test", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestBindModel_BindFromHeaders(t *testing.T) {
	app := fiber.New()
	log, _ := zap.NewDevelopment()

	app.Post("/test", func(c *fiber.Ctx) error {
		model := new(TestModel)
		err := binding.BindModel(log, c, model, binding.BindFromHeaders())
		assert.NoError(t, err)
		assert.Equal(t, "value1", model.Field1)
		return c.SendStatus(http.StatusOK)
	})

	req := httptest.NewRequest("POST", "/test", nil)
	req.Header.Set("Field1", "value1")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestBindModel_BindFromParams(t *testing.T) {
	app := fiber.New()
	log, _ := zap.NewDevelopment()

	app.Get("/test/:field2", func(c *fiber.Ctx) error {
		model := new(TestModel)
		err := binding.BindModel(log, c, model, binding.BindFromParams())
		assert.NoError(t, err)
		assert.Equal(t, "value2", model.Field2)
		return c.SendStatus(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test/value2", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
