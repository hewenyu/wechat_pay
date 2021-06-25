package login

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var BufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 4<<10)) // 4KB
	},
}

type PayData struct {
	data map[string]interface{}
}

func NewPayData() *PayData {

	return &PayData{
		data: make(map[string]interface{}),
	}
}

func (pd *PayData) IsSet(key string) bool {

	_, ok := pd.data[key]
	return ok
}

func (pd *PayData) Set(key string, val interface{}) {

	vKind := reflect.ValueOf(val).Kind()
	switch vKind {
	case reflect.String:
		pd.data[key] = val.(string)
	case reflect.Int:
		pd.data[key] = strconv.Itoa(val.(int))
	case reflect.Int64:
		pd.data[key] = strconv.FormatInt(val.(int64), 10)
	case reflect.Float32:
		pd.data[key] = strconv.FormatFloat(float64(val.(float32)), 'f', -1, 32)
	case reflect.Float64:
		pd.data[key] = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case reflect.Ptr:
		pd.data[key] = val
	case reflect.Struct:
		pd.data[key] = val
	case reflect.Map:
		pd.data[key] = val
	case reflect.Slice:
		pd.data[key] = val
	default:
		pd.data[key] = ""
	}
}

func (pd *PayData) Get(key string) string {

	val, ok := pd.data[key]
	if !ok {
		return ""
	}

	_, oks := val.(string)
	if oks {
		return val.(string)
	} else {
		b, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		str := string(b)
		if str == "null" {
			return ""
		}
		return str
	}
}

func (pd *PayData) ToUrl() string {

	url := ""
	for key := range pd.data {

		v := pd.Get(key)
		if key != "sign" && v != "" {

			url += key + "=" + v + "&"
		}
	}

	return strings.Trim(url, "&")
}

func (pd *PayData) ToJson() string {

	b, err := json.Marshal(pd.data)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (pd *PayData) FromJson(r io.Reader) error {

	buf := BufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer BufPool.Put(buf)

	_, err := buf.ReadFrom(r)
	if err != nil {
		return err
	}

	var jsonTemplate interface{}
	err = json.Unmarshal(buf.Bytes(), &jsonTemplate)
	if err != nil {
		return err
	}

	pd.data = jsonTemplate.(map[string]interface{})

	return nil
}

func (pd *PayData) FromJsonStr(s string) error {

	var jsonTemplate interface{}
	err := json.Unmarshal([]byte(s), &jsonTemplate)
	if err != nil {
		return err
	}

	pd.data = jsonTemplate.(map[string]interface{})

	return nil
}
