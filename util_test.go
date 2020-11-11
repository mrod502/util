package util

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestEncoders(t *testing.T) {
	type s struct {
		D decimal.Decimal
		T time.Time
	}
	x := s{D: decimal.New(123, -2), T: time.Now()}
	fmt.Println("value", reflect.ValueOf(x.D))

	var form = url.Values{}
	QueryEncoder.Encode(x, form)
	//encodedString := fmt.Sprintf("%+v", form)
	fmt.Printf("%+v\n", form)

}

func TestDB(t *testing.T) {
	for i := 0; i < 64; i++ {
		fmt.Println(strconv.FormatInt(int64(i), 36))
	}
}
