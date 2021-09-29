package router

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/pkg/errors"
	routing "github.com/qiangxue/fasthttp-routing"
	"net/url"
	"reflect"
	"strings"
	"sync"
)

var decoder = schema.NewDecoder() //把get, post 请求转成 struct

func init() {
	decoder.IgnoreUnknownKeys(true)
}

type Handler struct {
	method reflect.Method
}

var contextPool = &sync.Pool{New: func() interface{} {
	return &Context{}
}}

var validatorPool = &sync.Pool{
	New: func() interface{} {
		return validator.New()
	},
}

func (t *Handler) ServeHTTP(c *routing.Context) error {
	context := contextPool.Get().(*Context)
	defer contextPool.Put(context)
	defer lib.CollectRecover()
	context.Context = c
	tt := t.method.Func.Type()
	var in = make([]reflect.Value, tt.NumIn())
	foundContext := false
	for i := 0; i < tt.NumIn(); i++ {
		paramType := tt.In(i)
		paramName := ""
		var param reflect.Value
		if paramType.Kind() == reflect.Ptr {
			paramName = paramType.Elem().Name()
		} else {
			paramName = paramType.Name()
		}
		if paramName == "Context" {
			foundContext = true
			if paramType.Kind() == reflect.Ptr {
				param = reflect.ValueOf(context)
			} else {
				param = reflect.ValueOf(*context)
			}
			in[i] = param
			continue
		}
		if !foundContext {
			param = reflect.Zero(paramType)
		} else {
			// decode request param
			if req, err := t.parseRequest(context, paramType); err != nil {
				context.ERROR(err)
				context.Abort()
				return nil
			} else {
				param = req
			}
		}
		in[i] = param
	}
	out := t.method.Func.Call(in)
	if len(out) > 0 {
		for i := 0; i < len(out); i++ {
			if out[i].Type().Name() == "error" {
				if err := out[i].Interface(); err != nil {
					context.ERROR(err)
					return nil
				}
			}
		}
	}
	return nil
}

func (t *Handler) parseRequest(c *Context, inType reflect.Type) (ret reflect.Value, err error) {
	if inType.Kind() == reflect.Ptr {
		ret = reflect.New(inType.Elem())
	} else {
		ret = reflect.New(inType)
	}
	o := ret.Interface()
	if c.IsPost() {
		ct := string(c.Request.Header.Peek("Content-Type"))
		if strings.Contains(ct, "multipart/form-data") {
			m, err := c.MultipartForm()
			if err != nil {
				log.Error(err)
				return ret, errors.Wrap(err, "parse form-data failed")
			}
			if err := decoder.Decode(o, m.Value); err != nil {
				log.Error(err)
				return ret, errors.Wrap(err, "decode failed")
			}
		} else if strings.Contains(ct, "application/json") {
			if err := json.Unmarshal(c.PostBody(), o); err != nil {
				log.Error(err)
				return ret, errors.Wrap(err, "decode json data failed")
			}
		} else {
			u, err := url.ParseQuery(string(c.PostArgs().QueryString()))
			if err != nil {
				log.Error(err)
				return ret, errors.Wrap(err, "parse request failed")
			}
			if err := decoder.Decode(o, u); err != nil {
				log.Error(err)
				return ret, errors.Wrap(err, "decode failed")
			}
		}
	} else {
		u, err := url.ParseQuery(string(c.QueryArgs().QueryString()))
		if err != nil {
			log.Error(err)
			return ret, errors.Wrap(err, "parse request failed")
		}
		if err := decoder.Decode(o, u); err != nil {
			log.Error(err)
			return ret, errors.Wrap(err, "decode failed")
		}
	}
	if inType.Kind() == reflect.Ptr {
		ret = reflect.ValueOf(o)
	} else {
		ret = reflect.ValueOf(o).Elem()
	}
	return ret, runValidate(inType, ret, o)
}

func runValidate(inType reflect.Type, value reflect.Value, data interface{}) error {
	validate := validatorPool.Get().(*validator.Validate)
	defer func() {
		validatorPool.Put(validate)
	}()
	defer lib.CollectRecover()
	if err := validate.Struct(data); err != nil {
		return err
	}
	if inType.NumMethod() > 0 {
		for i := 0; i < inType.NumMethod(); i++ {
			m := inType.Method(i)
			if m.Name == "Validate" {
				validateResult := m.Func.Call([]reflect.Value{value})
				if len(validateResult) == 0 {
					return nil
				}
				var validateError interface{}
				for i := 0; i < len(validateResult); i++ {
					if validateResult[i].Type().Name() == "error" {
						validateError = validateResult[i].Interface()
					}
				}
				if validateError != nil {
					return validateError.(error)
				}
			}
		}
	}
	return nil
}
