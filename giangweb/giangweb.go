package giangweb

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		e.Logger.Fatal("Unable Load Config")
	}
}

//Start start application
func Start() {
	e.GET("/", showHome)
	e.GET("/items", indexItems)
	e.GET("/items/:item_name", showItem)
	e.GET("/products", indexProducts)
	e.POST("/products", createProduct)
	e.GET("/products/:id", showProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.GET("/items", indexItems)

	e.Logger.Print("Start from port %s", cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
