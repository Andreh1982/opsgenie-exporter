package environment

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type Single struct {
	ENVIRONMENT      string // nolint: golint
	OPSGENIE_API_URL string // nolint: golint
	OPSGENIE_API_KEY string // nolint: golint
}

func InitEnv() {
	err := godotenv.Load("./env.local")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	env := GetInstance()
	env.Setup()
}

func (e *Single) Setup() {
	e.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	e.OPSGENIE_API_KEY = os.Getenv("OPSGENIE_API_KEY")
	e.OPSGENIE_API_URL = os.Getenv("OPSGENIE_API_URL")

}

func (e *Single) IsDevelopment() bool {
	return e.ENVIRONMENT == "development"
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
			singleInstance.Setup()
		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return singleInstance
}
