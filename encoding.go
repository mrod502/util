package util

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/shopspring/decimal"
)

//Encoder and Decoder for url QueryStrings
var (
	QueryEncoder = schema.NewEncoder()
	QueryDecoder = schema.NewDecoder()
)

func init() {
	QueryEncoder.RegisterEncoder(decimal.Decimal{}, EncodeDecimal)
	QueryEncoder.RegisterEncoder(time.Time{}, EncodeTime)

	QueryDecoder.RegisterConverter(decimal.Decimal{}, ConvertDecimal)
	QueryDecoder.RegisterConverter(time.Time{}, ConvertTime)
}

//EncodeDecimal - convert to URL queryString
func EncodeDecimal(v reflect.Value) string {
	return fmt.Sprint(v)
}

//ConvertDecimal - convert from URL queryString
func ConvertDecimal(s string) reflect.Value {
	d, _ := decimal.NewFromString(s)
	return reflect.ValueOf(d)
}

//EncodeTime - convert to URL queryString (ms time)
func EncodeTime(v reflect.Value) string {
	t := v.Interface().(time.Time).UnixNano() / Million
	return fmt.Sprintf("%d", t)
}

//ConvertTime - convert from URL queryString (ms time)
func ConvertTime(s string) reflect.Value {
	v, _ := strconv.ParseInt(s, 10, 64)
	t := time.Unix(0, v*Million)
	return reflect.ValueOf(t)
}
