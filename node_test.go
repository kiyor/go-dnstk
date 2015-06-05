/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : node_test.go

* Purpose :

* Creation Date : 03-27-2015

* Last Modified : Fri 05 Jun 2015 12:14:33 PM PDT

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package dnstk

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := ParseFile("cn.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	f = f.RemoveAnycase([]string{"114.114.114.114"}).PingAble().UniqCity()
	fmt.Println(toJsonIndent(f))
}
