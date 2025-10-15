package types

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//
// Example of dependency injection.
// 'Learning Go': 7. Types, Methods, and Interfaces - P.174
//
// see also: https://github.com/google/wire

//
// data store
//

type DataStore interface {
	UserNameForID(userId string) (string, bool)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

//
// log
//

func LogOutput(message string) {
	fmt.Println(message)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (la LoggerAdapter) Log(message string) {
	la(message)
}

//
// logic
//

type Logic interface {
	SayHello(userID string) (string, error)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in Saygoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown err")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

//
// controller
//

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("in SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func TestDI(t *testing.T) {
	// assembly
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)

	// http.HandleFunc("/hello", c.SayHello)
	// http.ListenAndServe(":8080", nil)

	// use httptest
	r := httptest.NewRequest(http.MethodGet, "/hello?user_id=1", nil)
	rr := httptest.NewRecorder()

	c.SayHello(rr, r)
	resp, ok := io.ReadAll(rr.Body)
	fmt.Println(rr.Code, string(resp), ok)
}
