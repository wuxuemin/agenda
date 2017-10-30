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

type modelconfig struct {
	model    model
	filename string
}

var (
	logger = utils.NewLogger()
	models []modelconfig
)

func addModel(model model, filename string) {
	models = append(models, modelconfig{
		model:    model,
		filename: filename,
	})
}

// Init initializes registered models
func Init() {
	var err interface{}
	finished := make(chan bool)
	// initialize all models concurrently
	for _, m := range models {
		go func(m modelconfig) {
			defer func() {
				if e := recover(); e != nil {
					err = e
				}
				finished <- true
			}()
			path := viper.GetString(m.filename)
			if len(path) == 0 {
				os.Mkdir("accountdata", 0755)
				path = "accountdata/" + m.filename + ".json"
			}
			m.model.Init(path)
		}(m)
	}
	// wait for all models to finish initialization
	for _ = range models {
		<-finished
	}
	if err != nil {
		logger.Fatalln("[init model]", err)
	}
}
