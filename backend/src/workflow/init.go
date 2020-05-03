package workflow

import (
	"log"

	"github.com/BurntSushi/toml"
)

type workflowConfig struct {
	JwtSecret string `toml:"jwtsecret"`
}

var (
	aWorkflowConfig workflowConfig
)

func init() {
	_, err := toml.DecodeFile("../config/WorkflowRelated.toml", &aWorkflowConfig)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Fail to load Database related configurations.")
	}
	log.Println(aWorkflowConfig)
}
