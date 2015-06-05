/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : node_test.go

* Purpose :

* Creation Date : 03-27-2015

* Last Modified : Fri 05 Jun 2015 12:21:52 PM PDT

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package dnstk

import (
	"fmt"
	"github.com/kiyor/golib"
	"testing"
)

func TestParse(t *testing.T) {
	golib.Osexec("wget -N http://public-dns.tk/nameserver/cn.json")
	f, err := ParseFile("cn.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	f = f.RemoveAnycase([]string{"114.114.114.114"}).PingAble().UniqCity()
	fmt.Println(toJsonIndent(f))
}
