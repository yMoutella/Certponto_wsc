package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type President struct {
	InicioMandato string `json:"inicioMandato"`
	FimMandato    string `json:"fimMandato"`
	Nome          string `json:"nome"`
	Partido       string `json:"partido"`
	Vice          string `json:"vice"`
	Eleicao       string `json:"eleicao"`
}

func getPresidentList() []President {
	presidentList := make([]President, 0)

	collector := colly.NewCollector()

	collector.OnHTML("table", func(e *colly.HTMLElement) {

		temp := President{}
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			el.ForEach("td[nowrap]", func(i int, el *colly.HTMLElement) {
				switch i {
				case 0:
					temp.InicioMandato = strings.TrimSpace(el.Text)
				case 1:
					temp.FimMandato = strings.TrimSpace(el.Text)
				case 2:
					temp.Nome = strings.TrimSpace(el.Text)
				case 3:
					temp.Partido = strings.TrimSpace(el.Text)
				case 4:
					temp.Vice = strings.TrimSpace(el.Text)
				case 5:
					temp.Eleicao = strings.TrimSpace(el.Text)
				}
			})
			presidentList = append(presidentList, temp)
		})

	})

	collector.OnRequest(func(f *colly.Request) {
		fmt.Println("Visiting", f.URL.String())
	})

	collector.Visit("https://www.gov.br/cti/pt-br/trajetoria-historica/presidentes-da-republica-desde-a-criacao-do-cti")
	return presidentList
}

func getPresidents(c *gin.Context) {

	pList := getPresidentList()

	c.JSON(200, pList)
}

func main() {

	router := gin.Default()
	router.GET("/presidents", getPresidents)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Presidents API",
		})
	})
	router.Run(":8080")

}
