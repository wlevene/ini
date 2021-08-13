package ini

import (
	"fmt"
	"testing"

	"github.com/wlevene/ini/scanner"
)

func TestIniParse(t *testing.T) {
	doc := `
[section]
k=v
k2=1.1
k3=0.0.0.0
k4=a-1.0/23
k5=-12. 45
k6=~/fe/34lod093@   ##$:94983-0*&^%$
`
	s := scanner.New([]byte(doc))
	s.Scan()
}

func TestIniParse2(t *testing.T) {
	doc := `
a=b
c=d

[section]
k=v
kk =vv
kkk= vvv

[section1]
k =v1

[ section2 ]
k2=v2

[ section3 ]
k2 = 3213
`

	s := scanner.New([]byte(doc))

	s.Scan()

}

func TestIniParse3(t *testing.T) {
	doc := `
a=b
`
	s := scanner.New([]byte(doc))
	s.Scan()
}

func TestIniParse4(t *testing.T) {
	doc := `
a=b
[s1]

k=v



ip=0.0.0.0
[s2]
k=v
ip=127.0.0.1
`
	s := scanner.New([]byte(doc))
	s.Scan()
}

func TestIniParse5(t *testing.T) {
	doc := `
[s1]
k=v
`
	s := scanner.New([]byte(doc))
	s.Scan()
}

func TestIniParse6(t *testing.T) {
	str := "123"
	s := &str
	s = nil

	fmt.Println(str, s)
}

func TestIniParseV2(t *testing.T) {
	doc := `
[section]
k=v
`
	s := scanner.New([]byte(doc))
	s.Scan()
}

func TestIniParseV22(t *testing.T) {
	doc := `
a=b
c=d

[section]
k=v
kk =1.22
kkk= vvv

[section1]
k =v1

[ section2 ]
k2=v2

[ section3 ]
k2 = 3213
`

	s := scanner.New([]byte(doc))

	s.Scan()

}
