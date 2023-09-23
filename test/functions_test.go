package just_a_test

import (
	// "fmt"
	"fmt"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"testing"
)

func TestPrint(t *testing.T) {
	// slice
	var variable5 = []string{"AAAA", "BBBB"}
	// var variable5 interface{} = "a"
	messageResult := globalFunction.GetMessage("sys002", variable5)

	fmt.Println(messageResult)
}


/*
bikin function dengan parameter (code string, replacement string)
return-nya :
	map[string]interface{}{
		"code": "conn001",
		"id": "Koneksi Error",
		"en": "Connection Error",
	}
*/

