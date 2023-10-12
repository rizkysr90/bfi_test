package main

import (
	"myapp/database"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)
type Ticket struct {
	Id int32 `json:"id"`
	Masalah string `json:"masalah"`
	UserId int32 `json:"userId"`
	CreatedAt time.Time
	DeletedAt time.Time
	UpdatedAt time.Time

}
type User struct {
	Id            int32          `json:"id"`
	Username      string         `json:"username"`
	Name 			string		`json:"name"`
	Ticket []Ticket  `json:"ticket"`
}

func main() {
	// init database
	db := database.Connection()

	// create table
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	e.POST("/tickets", func(c echo.Context) error {
		// get payload from request body
		// {masalah : string, user_id : number}
		u := new(Ticket)
		if err := c.Bind(u); err != nil {
			return err
		}
		// Create database
		ticket := Ticket{Masalah: u.Masalah, UserId: u.UserId, CreatedAt: time.Now()}
		db.Create(&ticket)
			
		return c.JSON(http.StatusCreated, u)
	})
	e.GET("/tickets", func(c echo.Context) error {
		var users []User
		db.Model(&User{}).Preload("Ticket").Find(&users)
		return c.JSON(http.StatusOK, users)

	})
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		// Create database
		user := User{Username: u.Username, Name: u.Name}
		db.Create(&user)
		return c.JSON(http.StatusCreated, u)

	})
    e.Logger.Fatal(e.Start(":8000"))

}
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "host=localhost user=postgres password=adarizki123 dbname=bfi_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// dsn := "root:ad\arizki123@tcp(127.0.0.1:5432)/bfi_Test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})