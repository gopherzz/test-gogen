package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

const (
	HANDLERS_DIR = "handlers/"
)

type config struct {
	Action         string          `json:"url"`
	HandlerConfigs []handlerConfig `json:"handlers"`
}

type handlerConfig struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Placeholder string `json:"ph"`
	Required    bool   `json:"required"`
}

func main() {
	var cfgFilename string
	var cfg config
	flag.StringVar(&cfgFilename, "config", "config.json", "Handlers configs filename.")
	flag.Parse()

	file, err := ioutil.ReadFile(cfgFilename)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := json.Unmarshal(file, &cfg); err != nil {
		log.Fatal(err.Error())
	}

	err = os.Mkdir(HANDLERS_DIR, os.ModeDir)
	if err != nil {
		log.Fatal(err.Error())
	}

	handlersGoFile, err := os.OpenFile(HANDLERS_DIR+"handlers.go", os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err.Error())
	}

	handlersGoTemplate, err := template.ParseGlob("templates/handlers.template")
	if err != nil {
		log.Fatal(err.Error())
	}
	handlersGoTemplate.Execute(handlersGoFile, cfg)

	handlersHtmlFile, err := os.OpenFile(HANDLERS_DIR+"form.html", os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err.Error())
	}

	handlersHtmlTemplate, err := template.ParseGlob("templates/form.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	handlersHtmlTemplate.Execute(handlersHtmlFile, cfg)
}
