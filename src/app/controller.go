package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok {
			c.String(http.StatusNotFound, "404 page not found")
			return
		}

		c.HTML(http.StatusOK, "vacation-overview.html", map[string]interface{}{
			"TimesOff": timesOff,
		})
	})

	r.POST("/employees/:id/vacation/new", func(c *gin.Context) {
		var timeOff TimeOff
		err := c.BindJSON(&timeOff)

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		id := c.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok {
			TimesOff[id] = []TimeOff{}
		}

		TimesOff[id] = append(timesOff, timeOff)
		c.JSON(http.StatusCreated, &timeOff)
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", gin.H{
			"Employees": employees,
		})
	})
	admin.GET("/employee/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			c.HTML(http.StatusOK, "admin-employee-add.html", nil)
			return
		}

		employee, ok := employees[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - No such employee")
			return
		}

		c.HTML(http.StatusOK, "admin-employee-edit.html", gin.H{
			"Employee": employee,
		})
	})

	admin.POST("/employee/:id", func(c *gin.Context) {
		id := c.Param("id")

		if id == "Add" {
			pto, err := strconv.ParseFloat(c.PostForm("pto"), 32)
			if err != nil {
				c.String(http.StatusNotFound, "Error adding")
				return
			}
			log.Panicln(pto)
			c.HTML(http.StatusOK, "pto", nil)
			return
		}

		c.HTML(http.StatusOK, "EDIT", nil)
	})

	//fix css images path
	r.Static("/public", "./public")

	return r
}
