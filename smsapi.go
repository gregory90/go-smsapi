package smsapi

import (
	"errors"
	"strings"

	"github.com/gregory90/go-webutils"
	"github.com/gregory90/go-webutils/request"

	. "github.com/gregory90/go-webutils/logger"
)

var (
	apiHost  string
	user     string
	password string
	hostname string
	path     string
)

func Init(apiHostArg string, userArg string, passwordArg string, hostnameArg string, pathArg string) {
	apiHost = apiHostArg
	user = userArg
	password = passwordArg
	hostname = hostnameArg
	path = pathArg
}

func Send(message string, from string, to string, uid string, test string) (bool, string, error) {
	passwordHash := utils.GetMD5Hash(password)

	_, body, errs := request.Client.Get(apiHost).Query("username=" + user).Query("password=" + passwordHash).Query("to=" + to).Query("message=" + message).Query("from=" + from).Query("encoding=" + "utf-8").Query("idx=" + uid).Query("notify_url=" + hostname + path).Query("test=" + test).End()
	if len(errs) > 0 {
		for _, err := range errs {
			Log.Error(err.Error())
		}
		return false, "", errors.New("SMS not sent, errors logged")
	}

	result := strings.Split(body, ":")

	if result[0] == "OK" {
		return true, "", nil
	} else {
		return false, result[1], nil
	}
}
