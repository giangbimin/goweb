import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
)

// start application
func Start() {
	port := os.Getenv("GO_WEB_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	v := validator.New()

	e.GET("/", showHome)
	e.GET("/items", indexItems)
	e.GET("/items/:item_name", showItem)
	e.GET("/products", indexProducts)
	e.POST("/products", createProduct)
	e.GET("/products/:id", showProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.GET("/items", indexItems)

	e.Logger.Print("Start from port %s", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}