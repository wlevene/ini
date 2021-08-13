package ini

import (
	"fmt"
	"testing"
	"time"
)

func TestIni(t *testing.T) {
	doc := `
; 123
[section]
k=v
; dsfads 
;123
#3452345


[section1]
k1=v1
`
	ini := New().Load([]byte(doc))
	ini.Dump()
}

func TestIni0(t *testing.T) {
	doc := `
[section]
k=v
`
	v := New().Load([]byte(doc)).Section("section").Get("k")
	fmt.Println("v: ", v)
	if v != "v" {
		t.Errorf("error %s", v)
	}
}

func TestIni2(t *testing.T) {
	doc := `
[section]
k=v
k1 = 1
`
	ini := New().Load([]byte(doc))

	v := ini.Section("section").Get("k")
	if v != "v" {
		t.Errorf("error %s", v)
	}

	iv := ini.GetInt("k1")
	if iv != 1 {
		t.Errorf("error %d", iv)
	}

	iv = ini.GetInt("k2")
	if iv != 0 {
		t.Errorf("error %d", iv)
	}

	iv = ini.GetIntDef("k2", 12)
	if iv != 12 {
		t.Errorf("error %d", iv)
	}

}

func TestIni3(t *testing.T) {
	doc := `
a =b
c=d
[section]
k =v
`
	ini := New().Load([]byte(doc))
	ini.Dump()

	v := ini.Section("section").Get("k")
	if v != "v" {
		t.Errorf("error %s", v)
	}
	v = ini.Section("").Get("a")
	if v != "b" {
		t.Errorf("error %s", v)
	}

}

func TestIni4(t *testing.T) {
	doc := `
a =b
c= d

a1 = 2.1 
`
	ini := New().Section("").Load([]byte(doc))
	// ini.Dump()
	v := ini.Get("a")
	if v != "b" {
		t.Errorf("error %s:%d", v, len(v))
	}

	iv := ini.GetInt("a")
	if iv != 0 {
		t.Errorf("error %d", iv)
	}

	iv = ini.GetIntDef("a", 10)
	if iv != 10 {
		t.Errorf("error %d", iv)
	}

	iv = ini.GetIntDef("a1", 10)
	if iv != 10 {
		t.Errorf("error %d", iv)
	}

	fv := ini.GetFloat64Def("a1", 10)
	if fv != 2.1 {
		t.Errorf("error %d", iv)
	}

}

func TestIni5(t *testing.T) {
	doc := `
a =b
[s1]
k=v
k1 = v12

[s2]
k2=v2
k2= v22

[s3]
k =v
a= b
`
	ini := New().Load([]byte(doc))
	fmt.Println(string(ini.Marshal2Json()))

}

func TestIniFile(t *testing.T) {
	file := "./test.ini"
	ini := New().LoadFile(file)
	ini.Dump()

	fmt.Println(string(ini.Marshal2Json()))
	fmt.Println(ini.Err())

}

func TestIniWatchFile(t *testing.T) {
	file := "./test.ini"

	ini := New().WatchFile(file)
	fmt.Println(string(ini.Marshal2Json()))

	time.Sleep(10 * time.Second)
	fmt.Println(string(ini.Marshal2Json()))
	ini.StopWatch()

}
