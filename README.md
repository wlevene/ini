

![logo](./logo.png)



#  INI Parser Library

ini parser library for Golang,  easy-use、fast、 use ast parse content


[![Build Status](https://travis-ci.org/meolu/walden.svg?branch=master)](https://github.com/wlevene/ini)
![version](https://img.shields.io/badge/version-0.1.1-blue)

# Features

* can be read by []byte
* can be read by file
* Supports file monitoring and takes effect in real time without reloading
* Unmarshal to struct
* Marshal to Json



# Installation

```shell
go get github.com/wlevene/ini
```



# Example

```go
import (
	"fmt"

	"github.com/wlevene/ini"
)
```

### GetValue

```go
doc := `
[section]
k=v

[section1]
k1=v1
k2=1
k3=3.5
k4=0.0.0.0
`
v1 := ini.New().Load([]byte(doc)).Section("section1").Get("k1")
fmt.Println(v1)
```

Output

```
v1
```



```go
i := ini.New().Load([]byte(doc))
v1 := i.Section("section1").Get("k1")
v2 := i.GetInt("k2")
v3 := i.GetFloat64("k3")
v4 := i.Get("k4")
v5 := i.GetIntDef("keyint", 10)
v6 := i.GetDef("keys", "defualt")

fmt.Printf("v1:%v v2:%v v3:%v v4:%v v5:%v v6:%v\n", v1, v2, v3, v4, v5, v6)
```



Output

```
v1:v1 v2:1 v3:3.5 v4:0.0.0.0 v5:10 v6:defualt
```



### Marshal2Json

```go
fmt.Println(string(i.Marshal2Json()))
```

Output

```json
{"section":{"k":"v"},"section1":{"k1":"v1","k2":"1","k3":"3.5","k4":"0.0.0.0"}}
```



### Unmarshal Struct

```go
type TestConfig struct {
	K    string  `ini:"k" json:"k,omitempty"`
	K1   int     `ini:"k1" json:"k1,omitempty"`
	K2   float64 `ini:"k2"`
	K3   int64   `ini:"k3"`
	User User    `ini:"user"`
}

type User struct {
	Name string `ini:"name"`
	Age  int    `ini:"age"`
}


```

```go
doc := `
k=v
k1=2
k2=2.2
k3=3

[user]
name=tom
age=-23
`

cfg := TestConfig{}

ini.Unmarshal([]byte(doc), &cfg)
fmt.Println("cfg:", cfg)
```

Output

```
{v 2 2.2 3 {tom -23}}
```



### Parse File

ini file

```ini
; this is comment
; author levene 
; date 2021-8-1


a='23'34?::'<>,.'
c=d

[s1]
k=67676
k1 =fdasf 
k2= sdafj3490&@)34 34w2

# comment 
# 12.0.0.1
[s2]

k=3


k2=945
k3=-435
k4=0.0.0.0

k5=127.0.0.1
k6=levene@github.com

k7=~/.path.txt
k8=./34/34/uh.txt

k9=234@!@#$%^&*()324
k10='23'34?::'<>,.'

```

```go
file := "./test.ini"
ini.New().LoadFile(file).Section("s2").Get("k2")

fmt.Println(string(ini.Marshal2Json()))
```

Output

```
945
```



### Watch File

```go
file := "./test.ini"

idoc := idoc.New().WatchFile(file)
v := idoc.Section("s2").Get("k1")
fmt.Println("v:", v1)

// modify k1=v1   ==> k1=v2
time.Sleep(10 * time.Second)

v = idoc.Section("s2").Get("k1")
fmt.Println("v:", v1)
```

Output

```
v: v1
v: v2
```





Print file json

```go
file := "./test.ini"
fmt.Println(string(ini.New().LoadFile(file).Marshal2Json()))
```

Output

```json
{
  "a": "'23'34?::'<>,.'",
  "c": "d",
  "s1": {
    "k": "67676",
    "k2": "34w2"
  },
  "s2": {
    "k": "3",
    "k10": "'23'34?::'<>,.'",
    "k2": "945",
    "k3": "-435",
    "k4": "0.0.0.0",
    "k5": "127.0.0.1",
    "k6": "levene@github.com",
    "k7": "~/.path.txt",
    "k8": "./34/34/uh.txt",
    "k9": "234@!@#$%^&*()324"
  }
}
```


### Dump AST struct

```json
KVNode {
    Key: a
    Value: '23'34?::'<>,.'
}
KVNode {
    Key: c
    Value: d
}
Section {
    Section: [s1]
    KVNode {
        Key: k
        Value: 67676
    }
    KVNode {
        Key: k2
        Value: 34w2
    }
}
Section {
    Section: [s2]
    KVNode {
        Value: 3
        Key: k
    }
    KVNode {
        Key: k2
        Value: 945
    }
    KVNode {
        Key: k3
        Value: -435
    }
    KVNode {
        Key: k4
        Value: 0.0.0.0
    }
    KVNode {
        Value: 127.0.0.1
        Key: k5
    }
    KVNode {
        Key: k6
        Value: levene@github.com
    }
    KVNode {
        Key: k7
        Value: ~/.path.txt
    }
    KVNode {
        Key: k8
        Value: ./34/34/uh.txt
    }
    KVNode {
        Key: k9
        Value: 234@!@#$%^&*()324
    }
    KVNode {
        Key: k10
        Value: '23'34?::'<>,.'
    }
}
```


## Contributors





## License

[MIT](https://github.com/RichardLitt/standard-readme/blob/master/LICENSE)  
