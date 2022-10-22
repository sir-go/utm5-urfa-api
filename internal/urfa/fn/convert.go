package fn

import (
	"encoding/base64"
	"net"
	"time"

	"github.com/pkg/errors"
)

const ErrFormat = "value of '%s' cast error (mustHave be a %s)"

var NotImplemented = errors.New("method is not implemented yet")

func mustHave(p Dict, name string) {
	_, ok := p[name]
	if !ok {
		panic(errors.Errorf("%s is required", name))
	}
}

func forEach(c conn, p Dict, name string, cb func(item interface{})) (length int) {
	if _, ok := p[name]; !ok {
		c.PutInt(0)
		return 0
	}

	a := p[name].([]interface{})
	length = len(a)
	c.PutInt(length)
	for _, ai := range a {
		cb(ai)
	}
	return
}

func forEachDict(c conn, p Dict, name string, cb func(arr Dict)) (length int) {
	return forEach(c, p, name, func(item interface{}) {
		cb(item.(map[string]interface{}))
	})
}

func getMapOf(c conn, dictGen func() Dict, amount ...int) (res []Dict) {
	var sz int
	if len(amount) != 0 {
		sz = amount[0]
	} else {
		sz = c.GetI()
	}
	res = make([]Dict, sz)

	for i := 0; i < sz; i++ {
		res[i] = dictGen()
	}
	return
}

func putI(c conn, p Dict, name string, def ...int) int {
	val, ok := p[name]
	if !ok {
		if len(def) < 1 {
			panic(errors.Errorf("%s is required", name))
		}
		val = def[0]
	}

	switch val.(type) {
	case int:
		c.PutInt(val.(int))
		return val.(int)
	case float64:
		c.PutInt(int(val.(float64)))
		return int(val.(float64))
	default:
		panic(errors.Errorf(ErrFormat, name, "number"))
	}
}

func putL(c conn, p Dict, name string, def ...int64) int64 {
	val, ok := p[name]
	if !ok {
		if len(def) < 1 {
			panic(errors.Errorf("%s is required", name))
		}
		val = def[0]
	}

	switch val.(type) {
	case int:
		c.PutLong(int64(val.(int)))
		return int64(val.(int))
	case float64:
		c.PutLong(int64(val.(float64)))
		return int64(val.(float64))
	default:
		panic(errors.Errorf(ErrFormat, name, "number"))
	}
}

func putD(c conn, p Dict, name string, def ...float64) float64 {
	val, ok := p[name]
	if !ok {
		if len(def) < 1 {
			panic(errors.Errorf("%s is required", name))
		}
		val = def[0]
	}

	switch val.(type) {
	case int:
		c.PutDbl(float64(val.(int)))
		return float64(val.(int))
	case float64:
		c.PutDbl(val.(float64))
		return val.(float64)
	default:
		panic(errors.Errorf(ErrFormat, name, "number"))
	}
}

func putS(c conn, p Dict, name string, def ...string) string {
	val, ok := p[name]
	if !ok {
		if len(def) < 1 {
			panic(errors.Errorf("%s is required", name))
		}
		val = def[0]
	}

	switch val.(type) {
	case string:
		c.PutStr(val.(string))
		return val.(string)
	default:
		panic(errors.Errorf(ErrFormat, name, "string"))
	}

}

func putA(c conn, p Dict, name string, def ...string) net.IP {
	val, ok := p[name]
	if !ok {
		if len(def) < 1 {
			panic(errors.Errorf("%s is required", name))
		}
		val = def[0]
	}

	switch val.(type) {
	case string:
		if _v := net.ParseIP(val.(string)); _v != nil {
			c.PutIP(_v)
			return _v
		} else {
			panic(errors.Errorf(ErrFormat, name, "string"))
		}
	default:
		panic(errors.Errorf(ErrFormat, name, "string"))
	}
}

func sendBin(c conn, p Dict, name string) {
	val, ok := p[name]
	if !ok {
		panic(errors.Errorf("%s is required", name))
	}

	switch val.(type) {
	case string:
		if buff, err := base64.StdEncoding.DecodeString(val.(string)); err == nil {
			c.SendBin(buff)
		} else {
			panic(errors.Errorf(ErrFormat, name, "Base64 string"))
		}
	case []byte:
		c.SendBin(val.([]byte))
	default:
		panic(errors.Errorf(ErrFormat, name, "Base64 string"))
	}
}

func TimeNow() int {
	return int(time.Now().Unix())
}

func TimeTodayMidnight() int {
	tNow := time.Now()
	return int(
		time.Date(tNow.Year(), tNow.Month(), tNow.Day(),
			0, 0, 0, 0, time.Local).Unix())
}

func TimeThisMonthBegin() int {
	tNow := time.Now()
	return int(
		time.Date(tNow.Year(), tNow.Month(), 1,
			0, 0, 0, 0, time.Local).Unix())
}
