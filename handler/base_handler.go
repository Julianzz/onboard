package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// RestHandlerInterface restful hander interface
type RestHandlerInterface interface {
	Get(w http.ResponseWriter, req *http.Request, params map[string]string)
	Put(res http.ResponseWriter, req *http.Request, params map[string]string, body []byte)
	Post(res http.ResponseWriter, req *http.Request, params map[string]string, body []byte)
	Error(res http.ResponseWriter, req *http.Request, params map[string]string)
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

// RestfulHandler base restfule handler, will delegate method to other interface
type RestfulHandler struct {
	H RestHandlerInterface
}

func (handler *RestfulHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	switch req.Method {
	case "GET":
		handler.H.Get(res, req, vars)
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		// TODO, how to deal with read body error
		if err != nil {
			fmt.Printf("error in read body %v\n", err)
		}
		handler.H.Post(res, req, vars, body)
	case "PUT":
		body, err := ioutil.ReadAll(req.Body)

		// TODO, how to deal with read body error
		if err != nil {
			fmt.Printf("error in read body %v\n", err)
		}
		handler.H.Put(res, req, vars, body)
	default:
		handler.H.Error(res, req, vars)
	}
}

// DefaultRestHandler default handler, will return 404
type DefaultRestHandler struct {
}

// Get handle get request
func (h *DefaultRestHandler) Get(w http.ResponseWriter, req *http.Request, params map[string]string) {
	w.WriteHeader(http.StatusBadRequest)
}

// Put handle put request
func (h *DefaultRestHandler) Put(w http.ResponseWriter, req *http.Request, params map[string]string, body []byte) {
	w.WriteHeader(http.StatusBadRequest)
}

// Post handler post request
func (h *DefaultRestHandler) Post(w http.ResponseWriter, req *http.Request, params map[string]string, body []byte) {
	w.WriteHeader(http.StatusBadRequest)
}

// handler error
func (h *DefaultRestHandler) Error(w http.ResponseWriter, req *http.Request, params map[string]string) {
	w.WriteHeader(http.StatusBadRequest)
}
