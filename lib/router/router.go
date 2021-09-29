package router

import (
	"fmt"
	"github.com/fatih/camelcase"
	"github.com/juxuny/clipboard-sync/lib/log"
	routing "github.com/qiangxue/fasthttp-routing"
	"path"
	"reflect"
	"runtime"
	"strings"
)

var noAuthPath = map[string]bool{}

func IsNoAuthPath(path string) bool {
	return noAuthPath[path]
}

type builder struct {
	m               map[interface{}]struct{}
	handlers        []routing.Handler
	prefix          string
	noAuthMethodSet map[uintptr]bool
}

type InterfaceCreator func() (group interface{}, noAuthMethodList []interface{})

func NewBuilder(urlPrefix string) *builder {
	r := &builder{
		prefix:          urlPrefix,
		m:               make(map[interface{}]struct{}),
		handlers:        make([]routing.Handler, 0),
		noAuthMethodSet: map[uintptr]bool{},
	}
	return r
}

func (t *builder) Use(handler ...routing.Handler) {
	t.handlers = append(t.handlers, handler...)
}

func (t *builder) Register(creator InterfaceCreator) {
	v, methods := creator()
	t.m[v] = struct{}{}
	for _, m := range methods {
		//log.Debug(reflect.ValueOf(m).Pointer())
		//log.Debug(runtime.FuncForPC(reflect.ValueOf(m).Pointer()).Name())
		t.noAuthMethodSet[reflect.ValueOf(m).Pointer()] = true
	}
}

func (t *builder) Build(r *routing.Router) *routing.Router {
	if r == nil {
		r = routing.New()
	}
	if len(t.handlers) > 0 {
		r.Use(t.handlers...)
	}
	for v := range t.m {
		if err := t.bind(r, v); err != nil {
			panic(err)
		}
	}
	return r
}

func (t *builder) bind(r *routing.Router, v interface{}) error {
	tt := reflect.TypeOf(v)
	pathSlice := make([]string, 0)
	if t.prefix != "" {
		pathSlice = append(pathSlice, t.prefix)
	}
	structName := tt.Name()
	if tt.Kind() == reflect.Ptr {
		structName = tt.Elem().Name()
	}
	l := camelcase.Split(structName)
	if len(l) > 0 {
		pathSlice = append(pathSlice, l...)
	}
	methodNum := tt.NumMethod()
	for i := 0; i < methodNum; i++ {
		method := tt.Method(i)
		l := camelcase.Split(method.Name)
		methodPath := append(pathSlice, l...)
		h := &Handler{method: method}
		apiPath := strings.ToLower(path.Join(methodPath...))
		var noAuth bool
		//log.Debug(runtime.FuncForPC(method.Func.Pointer()).Name())
		if t.noAuthMethodSet[method.Func.Pointer()] {
			noAuthPath[apiPath] = true
		}
		noAuth = noAuthPath[apiPath]
		r.Get(apiPath, h.ServeHTTP)
		r.Post(apiPath, h.ServeHTTP)
		log.Debug("register path: "+fmt.Sprintf("(auth=%v)", !noAuth), strings.ToLower(path.Join(methodPath...)), "caller:", runtime.FuncForPC(method.Func.Pointer()).Name())
	}
	return nil
}
