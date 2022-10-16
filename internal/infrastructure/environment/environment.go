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
	LOG_LEVEL        string // nolint: golint
}

func init() {
	err := godotenv.Load("./env.local")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	env := GetInstance()
	env.Setup()
}

func (e *Single) Setup() {
	e.ENVIRONMENT = getenv("ENVIRONMENT", "development")
	e.OPSGENIE_API_KEY = getenv("OPSGENIE_API_KEY", "")
	e.OPSGENIE_API_URL = getenv("OPSGENIE_API_URL", "")
	e.LOG_LEVEL = getenv("LOG_LEVEL", "debug")

}

func (e *Single) IsDevelopment() bool {
	return e.ENVIRONMENT == "development"
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
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
