package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RestHandlerInterface restful hander interface
type RestHandlerInterface interface {
	Get(w http.ResponseWriter, req *http.Request, params map[string]string) (interface{}, error)
	Put(res http.ResponseWriter, req *http.Request, params map[string]string, body []byte) (interface{}, error)
	Post(res http.ResponseWriter, req *http.Request, params map[string]string, body []byte) (interface{}, error)
	Error(res http.ResponseWriter, req *http.Request, params map[string]string) (interface{}, error)
}

// RestfulError struct for restful error, including http code and info for return
type RestfulError struct {
	err  error
	info string
	code int
}

// NewRestfulError create new restful error
func NewRestfulError(err error, code int, info string) *RestfulError {
	return &RestfulError{
		err:  err,
		code: code,
		info: info,
	}
}

func (cerr *RestfulError) Error() string {
	errorinfo := fmt.Sprintf("%v", cerr.err)
	return errorinfo
}

func (cerr *RestfulError) String() string {
	errorinfo := fmt.Sprintf("error:%v info:%v", cerr.err, cerr.info)
	return errorinfo
}

// RestfulHandler base restfule handler, will delegate method to other interface
type RestfulHandler struct {
	H RestHandlerInterface
}

func (handler *RestfulHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	var result interface{}
	var err error

	switch req.Method {
	case "GET":
		result, err = handler.H.Get(w, req, vars)
	case "POST":
		var body []byte
		body, err = ioutil.ReadAll(req.Body)
		// TODO, how to deal with read body error
		if err != nil {
			log.Printf("error in read body %v\n", err)
		}
		result, err = handler.H.Post(w, req, vars, body)
	case "PUT":
		var body []byte
		body, err = ioutil.ReadAll(req.Body)

		// TODO, how to deal with read body error, to ignore?
		if err != nil {
			fmt.Printf("error in read body %v\n", err)
		}
		result, err = handler.H.Put(w, req, vars, body)
	default:
		result, err = handler.H.Error(w, req, vars)
	}

	if err != nil {
		e := err.(*RestfulError)
		w.WriteHeader(e.code)
		w.Write([]byte(e.String()))
	} else {
		values, _ := json.Marshal(result)
		w.Write(values)
	}
}

// DefaultRestHandler default handler, will return 404
type DefaultRestHandler struct {
}

// Get handle get request
func (h *DefaultRestHandler) Get(w http.ResponseWriter, req *http.Request, params map[string]string) (interface{}, error) {
	return nil, NewRestfulError(errors.New("not implement this method"), http.StatusBadRequest, "")
}

// Put handle put request
func (h *DefaultRestHandler) Put(w http.ResponseWriter, req *http.Request, params map[string]string, body []byte) (interface{}, error) {
	return nil, NewRestfulError(errors.New("not implement this method"), http.StatusBadRequest, "")
}

// Post handler post request
func (h *DefaultRestHandler) Post(w http.ResponseWriter, req *http.Request, params map[string]string, body []byte) (interface{}, error) {
	return nil, NewRestfulError(errors.New("not implement this method"), http.StatusBadRequest, "")
}

// handler error
func (h *DefaultRestHandler) Error(w http.ResponseWriter, req *http.Request, params map[string]string) (interface{}, error) {
	return nil, NewRestfulError(errors.New("not implement this method"), http.StatusBadRequest, "")
}
