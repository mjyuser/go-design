package mediator

import (
	"os"
	"testing"
)


func Test_Mediator(t *testing.T) {
	NewEventDispatcher()
	repository := NewUserRepository()
	ur := repository.(*UserRepository)

	pwd, _ := os.Getwd()
	logger := NewLogger(pwd + "/app.log")
	eventDispatcher.Attach(ur, "register:user")
	eventDispatcher.Attach(Observer(logger), "*")

	ur.CreateUser(map[string]interface{}{}, false)
}