package giangweb

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// productParams validator
type productParams struct {
	Name            string `json:"product_name" validate:"required,min=4"`
	Vendor          string `json:"vendor" validate:"min=5,max=10"`
	Email           string `json:"email" validate:"required_with=Vendor,email"`
	Website         string `json:"website" validate:"url"`
	Country         string `json:"country" validate:"len=2"`
	DefaultDeviceIP string `json:"default_device_ip" validate:"ip"`
}

// PostValidator validator
type PostValidator struct {
	validator *validator.Validate
}

// Validate PostValidator
func (p *PostValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = map[int]string{1: "mobile", 2: "pc", 3: "laptop"}

func indexProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func createProduct(c echo.Context) error {
	var reqBody productParams
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := v.Struct(reqBody); err != nil {
		errs := err.(validator.ValidationErrors)
		e.Logger.Print("Start from port %err", errs)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errs.Error()})
	}
	if reqBody.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "format incorrect"})
	}
	products[len(products)+1] = reqBody.Name
	return c.JSON(http.StatusOK, products)
}

func showProduct(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var product = products[productID]
	if product == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not Found"})
	}
	return c.JSON(http.StatusOK, map[int]string{productID: product})
}

func updateProduct(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var product = products[productID]
	if product == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not Found"})
	}
	var reqBody productParams
	e.Validator = &PostValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	products[productID] = reqBody.Name
	return c.JSON(http.StatusOK, map[int]string{productID: products[productID]})
}

func deleteProduct(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var product = products[productID]
	if product == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not Found"})
	}
	delete(products, productID)
	return c.JSON(http.StatusOK, products)
}
