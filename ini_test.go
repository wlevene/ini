package ini

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestIni(t *testing.T) {
	doc := `
; 123
c=d
# 434

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
k =v
`
	i_doc := New().Load([]byte(doc)).Section("section")
	i_doc.Dump()

	v := i_doc.Get("k")
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
	ini.Dump()
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
	json_str := string(ini.Marshal2Json())
	fmt.Println(json_str)
	if json_str != `{"a":"b","s1":{"k":"v","k1":"v12"},"s2":{"k2":"v22"},"s3":{"a":"b","k":"v"}}` {
		t.Errorf("error %v", json_str)
	}

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

func TestIniDelete(t *testing.T) {
	doc := `
k=v
a=b
c=d
[section]

`
	ini := New().Load([]byte(doc))
	fmt.Println("--------------------------------")
	ini.Dump()

	fmt.Println("--------------------------------")
	ini.Del("a")
	ini.Dump()

	fmt.Println("--------------------------------")
	ini.Del("c")
	ini.Dump()

	fmt.Println("--------------------------------")
	ini.Del("k")
	ini.Dump()

}

func TestIniSet(t *testing.T) {
	doc := `
k =v
[section]
a=b
c=d
`
	ini := New().Load([]byte(doc)).Section("section")
	fmt.Println("--------------------------------")
	ini.Dump()

	fmt.Println("--------------------------------")
	ini.Set("a", 11).Set("c", 12.3).Section("").Set("k", "SET")
	ini.Dump()

	v := ini.Section("section").GetInt("a")

	if v != 11 {
		t.Errorf("Error: %d", v)
	}

	v1 := ini.GetFloat64("c")

	if v1 != 12.3 {
		t.Errorf("Error: %f", v1)
	}

	v2 := ini.Section("").Get("k")
	if v2 != "SET" {
		t.Errorf("Error: %s", v2)
	}

	fmt.Println("--------------------------------")
	ini.Set("a1", 1).Section("section").Set("k1", 11.11)
	ini.Dump()

}

func TestIniSave(t *testing.T) {
	doc := `
; 123
c11=d12312312
# 434

[section]
k=v
; dsfads 
;123
#3452345


[section1]
k1=v1

[section3]
k3=v3
`
	ini := New().Load([]byte(doc))
	ini.Dump()

	ini.Save("./save.ini")
	fmt.Println(ini.Err())
}

func TestIniSave2(t *testing.T) {

	filename := "./save.ini"
	ini := New().Set("a1", 1)
	ini.Dump()
	ini.Save(filename)

	bts, _ := ioutil.ReadFile(filename)

	if string(bts) != "a1 = 1\n" {
		t.Errorf("Error: %v", string(bts))
	}
}

func TestIniSave3(t *testing.T) {

	filename := "./save.ini"
	ini := New().Set("a1", 1).Section("s1").Set("a2", "v2")
	ini.Dump()
	ini.Save(filename)

	bts, _ := ioutil.ReadFile(filename)

	if string(bts) != "a1 = 1\n\n[s1]\na2 = v2\n" {
		t.Errorf("Error: %v", string(bts))
	}
}
