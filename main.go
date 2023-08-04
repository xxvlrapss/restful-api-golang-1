package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type MenuItem struct { 
	Name string
	OrderCode string
	Price int
}

func getFoodMenu(c echo.Context) error {
	 foodMenu := []MenuItem{
		{
			Name: "Indomie",
			OrderCode: "indomie",
			Price: 30000,
		},
		{
			Name: "Ayam lada hitam",
			OrderCode: "ayam_lada_hitam",
			Price: 40000,
		},
		{
			Name: "Nasi Goreng",
			OrderCode: "nasi_goreng",
			Price: 15000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

type DrinkItem struct { 
	Name string
	OrderCode string
	Price int
}

func getDrinkMenu(c echo.Context) error {
	 DrinksMenu := []DrinkItem{
		{
			Name: "Teh Manis",
			OrderCode: "teh_manis",
			Price: 12000,
		},
		{
			Name: "Jus Jeruk",
			OrderCode: "jus_jeruk",
			Price: 15000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": DrinksMenu,
	})
}


func main() {
	e := echo.New()
	// localhost:14040/menu/food
 e.GET("/menu/food", getFoodMenu)
 e.GET("/menu/drink", getDrinkMenu)
 e.Logger.Fatal(e.Start(":14040"))
}

