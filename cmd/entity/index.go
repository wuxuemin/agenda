package entity

import (
	"os"
	"utils"

	"github.com/spf13/viper"
)

type model interface {
	Init(string)
	load()
	dump()
}

type modelConfig struct {
	model    model
	filename string
}

var (
	logger = utils.NewLogger()
	models []modelConfig
)

func addModel(model model, filename string) {
	models = append(models, modelConfig{
		model:    model,
		filename: filename,
	})
}

// Init initializes registered models
func Init() {
	var err interface{}
	finished := make(chan bool)
	// initialize all models concurrently
	for _, i := range models {
		go func(i modelConfig) {
			defer func() {
				if e := recover(); e != nil {
					err = e
				}
				finished <- true
			}()
			path := viper.GetString(i.filename)
			if len(path) == 0 {
				os.Mkdir("data", 0755)
				path = "data/" + i.filename + ".json"
			}
			i.model.Init(path)
		}(i)
	}
	// wait for all models to finish initialization
	for _ = range models {
		<-finished
	}
	if err != nil {
		logger.Fatalln("[init model]", err)
	}
}
