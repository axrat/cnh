package main
import (
	"github.com/TransAssist/goz"
)
import (
	_ "github.com/mattn/go-sqlite3"
)
import (
	"fmt"
	"github.com/andelf/go-curl"
)

func main() {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	url:="https://api.bitflyer.jp/v1/getmarkets"
	easy.Setopt(curl.OPT_URL, url)
	fooTest := func (buf []byte, userdata interface{}) bool {
		println("DEBUG: size=>", len(buf))
		println("DEBUG: content=>", string(buf))
		return true
	}
	easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)
	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	goz.Complete()
}
