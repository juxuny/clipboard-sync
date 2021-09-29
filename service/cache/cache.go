package cache

import (
	"encoding/json"
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/pkg/errors"
	"strings"
	"sync"
	"time"
)

type node struct {
	Data      string    `json:"data"`
	ExpiredAt time.Time `json:"expired_at"`
}

var memoryMap = sync.Map{}

func Set(key string, v interface{}, expireDuration ...time.Duration) error {
	n := node{
		Data: lib.ToJSON(v),
	}
	if len(expireDuration) <= 0 {
		n.ExpiredAt = time.Now().Add(time.Hour * 24 * 30)
	}
	data := lib.ToJSON(n)
	memoryMap.Store(key, data)
	return nil
}

func get(key string, out interface{}) error {
	v, b := memoryMap.Load(key)
	if !b {
		log.Debug("not found")
		return lib.ErrNotFound
	}
	var n node
	err := json.Unmarshal([]byte(v.(string)), &n)
	if err != nil {
		return errors.Wrap(err, "invalid cache data")
	}
	if n.ExpiredAt.Before(time.Now()) {
		return lib.ErrNotFound
	}
	return json.Unmarshal([]byte(n.Data), out)
}

func Get(key string, out interface{}) error {
	if err := get(key, out); err == nil {
		return nil
	} else {
		return err
	}
}

func HSet(value interface{}, key ...string) error {
	if len(key) == 0 {
		return errors.Errorf("key is empty")
	}
	k := strings.Join(key, ":")
	log.Debug(k)
	return Set(k, value)
}

func HGet(outValue interface{}, key ...string) error {
	if len(key) == 0 {
		return errors.Errorf("key is empty")
	}
	k := strings.Join(key, ":")
	log.Debug(k)
	return Get(k, outValue)
}
