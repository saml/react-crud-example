package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEnvInt(envVar string, defaultValue int) int {
	i, err := strconv.ParseInt(os.Getenv(envVar), 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(i)
}

func getEnvString(envVar string, defaultValue string) string {
	s := os.Getenv(envVar)
	if s == "" {
		return defaultValue
	}
	return s
}

func renderHomepage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
func newWebApp(templateDir string, staticDir string) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob(templateDir + "/*")
	r.GET("/", renderHomepage)
	r.Static("/public", staticDir)
	r.GET("/public")
	return r
}

func main() {
	docRoot := flag.String("htdocs", getEnvString("HTDOCS", "./build/public"), "Doc root")
	templateDir := flag.String("templates", getEnvString("TEMPLATE_DIR", "templates"), "Directory with templates")
	flag.Parse()

	log.Printf("htdocs = %s | templates = %s", *docRoot, *templateDir)
	app := newWebApp(*templateDir, *docRoot)
	app.Run()
}
