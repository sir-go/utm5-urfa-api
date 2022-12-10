package fn

import (
	"encoding/base64"
	"net"
	"time"

	"github.com/pkg/errors"
)

const ErrFormat = "value of '%s' cast error (mustHave be a %s)"

var NotImplemented = errors.New("method is not implemented yet")

// mustHave checks if a Dict contains a key, else - panic
func mustHave(p Dict, name string) {
	_, ok := p[name]
	if !ok {
		panic(errors.Errorf("%s is required", name))
	}
}

// forEach searches an array named by name in the Dict,
// applies a cb function to each element of the array (put it to the connection),
// puts the length of the array to the Connection,
// and returns this length.
// This function makes putting arrays to connection much simpler
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

// same forEach function but for processing the Dict elements
func forEachDict(c conn, p Dict, name string, cb func(arr Dict)) (length int) {
	return forEach(c, p, name, func(item interface{}) {
		cb(item.(map[string]interface{}))
	})
}

// getMapOf fetches arrays from the connection
// dictGen callback makes a map for each fetched element
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

// putI takes an Integer value from the Dict by key name (or def if nothing has found) and puts it to the connection
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

// putL takes a LongInt value from the Dict by key name (or def if nothing has found) and puts it to the connection
func putL(c conn, p Dict, name string, def ...int64) int64 {
	val, ok := p[name]
	if !ok {
		if len(def) < 1 {
			panic(errors.Errorf("%s is required", name))
		}
		val = def[0]
	}

	switch val.(type) {
	case int64:
		c.PutLong(val.(int64))
		return val.(int64)
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

// putD takes a Double value from the Dict by key name (or def if nothing has found) and puts it to the connection
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

// putS takes a String value from the Dict by key name (or def if nothing has found) and puts it to the connection
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

// putA takes an IP address (stored as a string) value from the Dict by key name
// (or def if nothing has found) and puts it to the connection
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

// sendBin takes a Bytes value from the Dict by key name (or def if nothing has found) and puts it to the connection
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

// TimeNow just returns the current timestamp
func TimeNow() int {
	return int(time.Now().Unix())
}

// TimeTodayMidnight returns the timestamp of the today's midnight
func TimeTodayMidnight() int {
	tNow := time.Now()
	return int(
		time.Date(tNow.Year(), tNow.Month(), tNow.Day(),
			0, 0, 0, 0, time.Local).Unix())
}

// TimeThisMonthBegin returns the timestamp of the current month's beginning
// (midnight of the first day of the current month)
func TimeThisMonthBegin() int {
	tNow := time.Now()
	return int(
		time.Date(tNow.Year(), tNow.Month(), 1,
			0, 0, 0, 0, time.Local).Unix())
}
