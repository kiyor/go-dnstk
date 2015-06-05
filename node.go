/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : node.go

* Purpose :

* Creation Date : 03-27-2015

* Last Modified : Fri 05 Jun 2015 12:11:58 PM PDT

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package dnstk

import (
	"encoding/json"
	"io/ioutil"
	// 	"log"
	// 	"fmt"
	"github.com/kiyor/golib"
	"strings"
	"time"
)

// {"checked_at":"2015-03-27T13:07:51+01:00","city":"Nanjing","country_id":"CN","created_at":"2013-06-21T03:10:16+02:00","error":null,"ip":"114.114.114.114","name":"public1.114dns.com.","reliability":0.99,"version":""}
type DNS struct {
	CheckedAt   time.Time   `json:"checked_at"`
	City        string      `json:"city"`
	CountryId   string      `json:"country_id"`
	CreatedAt   time.Time   `json:"created_at"`
	Error       interface{} `json:"error"`
	Ip          string      `json:"ip"`
	Name        string      `json:"name"`
	Reliability float64     `json:"reliability"`
	Version     string      `json:"version"`
}

type DNSFile []*DNS

func ParseFile(file string) (*DNSFile, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var f DNSFile
	err = json.Unmarshal(b, &f)
	return &f, err
}

func (f *DNSFile) PingAble() *DNSFile {
	var s string
	m := make(map[string]DNS)
	for _, v := range *f {
		s += v.Ip + " "
		m[v.Ip] = *v
	}
	out, _ := golib.Osexec("fping -u " + s)
	rm := strings.Split(out, "\n")
	for _, v := range rm {
		delete(m, v)
	}
	var nf DNSFile
	for _, v := range m {
		d := v
		nf = append(nf, &d)
	}
	return &nf
}

func in(list []string, str string) (has bool) {
	for _, v := range list {
		if v == str {
			has = true
		}
	}
	return
}

func (f *DNSFile) RemoveAnycase(anycastlist []string) *DNSFile {

	var nf DNSFile
	for _, v := range *f {
		if !in(anycastlist, v.Ip) {
			nf = append(nf, v)
		}
	}
	return &nf
}

func (f *DNSFile) UniqCity() *DNSFile {
	m := make(map[string]DNS)
	for _, d := range *f {
		// if city already exist
		if len(d.City) > 0 {
			if v, ok := m[d.City]; ok {
				if d.Reliability > v.Reliability {
					m[d.City] = *d
				}
			} else {
				if d.Ip != "114.114.114.114" {
					m[d.City] = *d
				}
			}
		}
	}
	// 	log.Println(toJsonIndent(m))
	var nf DNSFile
	for _, v := range m {
		d := v
		nf = append(nf, &d)
	}
	return &nf
	// 	log.Println(toJsonIndent(*f))
}

func toJsonIndent(i interface{}) string {
	j, _ := json.MarshalIndent(i, "", "  ")
	return string(j)
}
