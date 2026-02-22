package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func getHTTPRoutes(projectType string) map[string][]string {
	switch projectType {
	case "E-commerce":
		return map[string][]string{
			"user-service":    {"/users", "/users/{id}", "/users/login", "/users/register"},
			"product-service": {"/products", "/products/{id}", "/products/category/{category}"},
			"order-service":   {"/orders", "/orders/{id}", "/orders/user/{userId}"},
			"payment-service": {"/payments", "/payments/{id}", "/payments/status/{status}"},
			"cart-service":    {"/cart", "/cart/{id}", "/cart/user/{userId}"},
		}
	case "Video-Streaming":
		return map[string][]string{
			"user-service":           {"/users", "/users/{id}", "/users/login", "/users/register"},
			"video-service":          {"/videos", "/videos/{id}", "/videos/stream/{id}"},
			"subscription-service":   {"/subscriptions", "/subscriptions/{id}", "/subscriptions/user/{userId}"},
			"recommendation-service": {"/recommendations", "/recommendations/{userId}"},
			"analytics-service":      {"/analytics", "/analytics/{videoId}"},
		}
	case "Food-Delivery":
		return map[string][]string{
			"user-service":       {"/users", "/users/{id}", "/users/login", "/users/register"},
			"restaurant-service": {"/restaurants", "/restaurants/{id}", "/restaurants/cuisine/{cuisine}"},
			"order-service":      {"/orders", "/orders/{id}", "/orders/user/{userId}"},
			"delivery-service":   {"/deliveries", "/deliveries/{id}", "/deliveries/status/{status}"},
			"payment-service":    {"/payments", "/payments/{id}", "/payments/status/{status}"},
		}
	default:
		return map[string][]string{
			"api-service": {"/api", "/api/{id}"},
		}
	}
}

func generateRoutes(projectType, targetPath, httpRouter string) {
	routes := getHTTPRoutes(projectType)
	for service, endpoints := range routes {
		servicePath := filepath.Join(targetPath, service, "internal", "routes")
		routesFile := filepath.Join(servicePath, "routes.go")
		os.MkdirAll(servicePath, os.ModePerm)

		var routeContent strings.Builder
		routeContent.WriteString(fmt.Sprintf("package %s\n\n", strings.ReplaceAll(service, "-", "")))
		routeContent.WriteString(fmt.Sprintf("import (\n\t\"%s\"\n)\n\n", httpRouterImport(httpRouter)))
		routeContent.WriteString("func RegisterRoutes(router " + httpRouterType(httpRouter) + ") {\n")

		for _, endpoint := range endpoints {
			routeContent.WriteString(fmt.Sprintf("\trouter.%s(\"%s\", handler)\n", httpRouterMethod(httpRouter), endpoint))
		}

		routeContent.WriteString("}\n")

		err := os.WriteFile(routesFile, []byte(routeContent.String()), 0644)
		if err != nil {
			color.Red("Failed to create routes for %s: %v", service, err)
			return
		}
	}
}

func httpRouterImport(router string) string {
	switch router {
	case "Gin":
		return "github.com/gin-gonic/gin"
	case "Fiber":
		return "github.com/gofiber/fiber/v2"
	case "Chi":
		return "github.com/go-chi/chi/v5"
	case "Echo":
		return "github.com/labstack/echo/v4"
	case "Mux":
		return "github.com/gorilla/mux"
	default:
		return "net/http"
	}
}

func httpRouterType(router string) string {
	switch router {
	case "Gin":
		return "*gin.Engine"
	case "Fiber":
		return "*fiber.App"
	case "Chi":
		return "*chi.Mux"
	case "Echo":
		return "*echo.Echo"
	case "Mux":
		return "*mux.Router"
	default:
		return "http.ServeMux"
	}
}

func httpRouterMethod(router string) string {
	switch router {
	case "Gin", "Chi":
		return "GET"
	case "Fiber", "Echo":
		return "Get"
	case "Mux":
		return "HandleFunc"
	default:
		return "Handle"
	}
}

func getRouterInit(httpRouter string) (string, string, string) {
	switch httpRouter {
	case "Gin":
		return `"github.com/gin-gonic/gin"`, `
			r := gin.Default()
			r.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Welcome to {{SERVICE_NAME}}"})
			})`, `r.Run(":" + port)`
	case "Fiber":
		return `"github.com/gofiber/fiber/v2"`, `
			app := fiber.New()
			app.Get("/", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{"message": "Welcome to {{SERVICE_NAME}}"})
			})`, `app.Listen(":" + port)`
	case "Echo":
		return `"github.com/labstack/echo/v4"`, `
			e := echo.New()
			e.GET("/", func(c echo.Context) error {
				return c.JSON(200, map[string]string{"message": "Welcome to {{SERVICE_NAME}}"})
			})`, `e.Start(":" + port)`
	case "Mux":
		return `"github.com/gorilla/mux"`, `
			r := mux.NewRouter()
			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Welcome to {{SERVICE_NAME}}"))
			})`, `http.ListenAndServe(":" + port, r)`
	default:
		return `"net/http"`, `
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Welcome to {{SERVICE_NAME}}"))
			})`, `http.ListenAndServe(":" + port, nil)`
	}
}

func PrintRoutes(projectType string) {
	routes := getHTTPRoutes(projectType)

	for service, endpoints := range routes {
		fmt.Printf("Service: %s\n", service)
		for _, endpoint := range endpoints {
			fmt.Printf("  Route: %s\n", endpoint)
		}
		fmt.Println()
	}
}
