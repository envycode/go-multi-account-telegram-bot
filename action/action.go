package action

import (
	"encoding/json"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"log"
	"multi-account-telegram-bot/constant"
	"os"
	"os/exec"
)

type Action interface {
	Create(id, name string, action string) error
	Execute(id string, name string, message *telebot.Message, input string) (string, error)
}

type Impl struct {
}

func (i Impl) Create(id, name string, action string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	action = `var sender = JSON.parse(process.argv.slice(2)[0]); var payload = process.argv.slice(2)[1];` + action
	return ioutil.WriteFile(fmt.Sprintf("%s/%s/%s:%s.js", path, constant.InterpreterDir, id, name), []byte(action), 0777)
}

type MessageInformation struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

func (i Impl) Execute(id string, name string, message *telebot.Message, input string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	log.Println("execute custom action ", name)
	msg := MessageInformation{
		Username: message.Sender.Username,
		Fullname: fmt.Sprintf("%s %s", message.Sender.FirstName, message.Sender.LastName),
	}
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	cmd := exec.Command("node", fmt.Sprintf("%s/%s/%s:%s.js", path, constant.InterpreterDir, id, name), string(msgJson), input)
	cmd.Stderr = os.Stderr
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(result), nil
}
