package mediator

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type EventDispatcher struct {
	Observers map[string][]Observer
}

var eventDispatcher *EventDispatcher
var once sync.Once

func NewEventDispatcher() *EventDispatcher {
	once.Do(func() {
		eventDispatcher = &EventDispatcher{Observers: map[string][]Observer{}}
		eventDispatcher.Observers["*"] = make([]Observer, 0)
	})

	return eventDispatcher
}

func (dispatcher *EventDispatcher) InitEventGroup(event string) {
	if event == "" {
		event = "*"
	}

	if _, ok := dispatcher.Observers[event]; !ok {
		dispatcher.Observers[event] = make([]Observer, 0)
	}
}

func (dispatcher *EventDispatcher) GetEventObservers(event string) []Observer {
	dispatcher.InitEventGroup(event)
	return append(dispatcher.Observers[event], dispatcher.Observers["*"]...)
}

func (dispatcher *EventDispatcher) Attach(observer Observer, event string) {
	dispatcher.InitEventGroup(event)
	dispatcher.Observers[event] = append(dispatcher.Observers[event], observer)
}

func (dispatcher *EventDispatcher) Detach(observer Observer, event string) {
	observers := dispatcher.GetEventObservers(event)
	for i, o := range observers {
		if o == observer {
			dispatcher.Observers[event] = append(dispatcher.Observers[event][0:i-1], dispatcher.Observers[event][i:]...)
		}
	}
}

func (dispatcher *EventDispatcher) Trigger(event string, emitter interface{}, data interface{}) {
	for _, observer := range dispatcher.GetEventObservers(event) {
		observer.Update(event, emitter, data)
	}
}

type Observer interface {
	Update(event string, emitter interface{}, data interface{})
}


type UserRepository struct {
	Users  map[int64]*User
}

func NewUserRepository() Observer{
	ur := &UserRepository{
		Users: make(map[int64]*User),
	}
	eventDispatcher.Attach(ur, "users:deleted")

	return ur
}


func (u *UserRepository) Update(event string, emitter interface{}, data interface{}) {
	switch event {
	case "users:deleted":
		if emitter == u {
			return
		}
	}
}

func (u *UserRepository) Initialize(filename string) {
	eventDispatcher.Trigger("users:init", u, filename)
}

func (u *UserRepository) CreateUser(data map[string]interface{}, silent bool) *User {
	fmt.Println("UserRepository: Creating a user.")
	user := &User{}
	user.Update(data)

	id := time.Now().Unix()

	user.Update(map[string]interface{}{"id": id})
	u.Users[id] = user
	if !silent {
		eventDispatcher.Trigger("users:created", u, user)
	}

	return user
}

func (u *UserRepository) DeleteUser(user User, silent bool) {
	fmt.Println("UserRepository: Deleting a user")
	if _, ok := u.Users[user.Id]; !ok {
		return
	}

	delete(u.Users, user.Id)

	if !silent {
		eventDispatcher.Trigger("users:deleted", u, user)
	}
}

func (u *UserRepository) UpdateUser(user User, data map[string]interface{}, silent bool) *User {
	fmt.Println("UserRepository: Updating a user")
	if _, ok := u.Users[user.Id]; !ok {
		return nil
	}

	user.Update(data)
	if !silent {
		eventDispatcher.Trigger("users:updated", u, user)
	}

	return &user
}

type User struct {
	Id int64 `json:"id"`
}

func (u *User) Update(data map[string]interface{}) {}

type Logger struct {
	filename string
}

func NewLogger(filename string) *Logger {
	logger := &Logger{filename: filename}
	_, err := os.Stat(logger.filename)

	if err != nil {
		if os.IsExist(err) {
			return logger
		}

		_, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
	}

	return logger
}

func (l *Logger) Update(event string, emitter interface{}, data interface{}) {
	content, _ := json.Marshal(data)
	format := fmt.Sprintf("%s %s with data %s", time.Now().Format("2006-01-02 15:04:05"), event, content)
	file, err := os.OpenFile(l.filename, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(format)
	if err != nil {
		fmt.Println("write err:", err)
	}
	defer file.Close()
}

