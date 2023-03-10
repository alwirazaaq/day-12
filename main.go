package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

type Blog struct {
	Title    string
	Content  string
	Author   string
	PostDate string
}

var dataBlog = []Blog{
	{
		Title:    "tak tak dung dung tak",
		Content:  "itulah suara gendang cilacap",
		Author:   "sam smith",
		PostDate: "10 Mei 1888",
	},
	{
		Title:    "cebong vs kadrun",
		Content:  "kericuhan saat mendekati pemilu sangat menyenangkan",
		Author:   "bapak saya",
		PostDate: "10 Mei 2023",
	},
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	//route static untuk mengakses folder public
	e.Static("/ASSETS", "ASSETS")

	//rendeer
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	//Routing
	e.GET("/hello", helloworld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/project", project)
	e.GET("/testimonial", testimonial)
	e.GET("/project-detail/:id", projetDetail)
	e.POST("/project", postproject)
	e.POST("/delete-project/:id", deleteProject)

	fmt.Println("yahahahaha hayuukkk 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloworld(c echo.Context) error {
	return c.String(http.StatusOK, "hellojing")
}

func home(c echo.Context) error {
	blogs := map[string]interface{}{
		"Blogs": dataBlog,
	}
	return c.Render(http.StatusOK, "index.html", blogs)
}

func contact(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func project(c echo.Context) error {
	return c.Render(http.StatusOK, "project.html", nil)
}

func testimonial(c echo.Context) error {
	return c.Render(http.StatusOK, "testimonial.html", nil)
}

func projetDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":      id,
		"Title":   "project 1",
		"Content": "Bagaimana pemberontakan tersebut berjalan? Pemberontakan PKI Madiun diawali dengan melancarkan propaganda anti pemerintah dan pemogokan kerja oleh kaum buruh. Selain itu pemberontakan juga dilakukan dengan menculik dan membunuh beberapa tokoh negara.",
	}

	return c.Render(http.StatusOK, "project-detail.html", data)
}

func postproject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("description")

	println("Project Name : " + projectName)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Description : " + description)

	var newBlog = Blog{
		Title:    projectName,
		Content:  description,
		Author:   "yang bersangkutan",
		PostDate: time.Now().String(),
	}

	dataBlog = append(dataBlog, newBlog)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
