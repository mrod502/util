package util

import (
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/shopspring/decimal"
)

func formatDecimal(d decimal.Decimal) string {
	return fmt.Sprintf("%v", d)
}

var (
	decoder = schema.NewDecoder()
	encoder = schema.NewEncoder()
)

//URIEncode - encode a struct as url values
func URIEncode(v interface{}) (form url.Values, err error) {
	form = url.Values{}
	err = encoder.Encode(v, form)

	return
}
