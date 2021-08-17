package ini

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/wlevene/ini/ast"
	"github.com/wlevene/ini/lexer"
	"github.com/wlevene/ini/parser"
	"github.com/wlevene/ini/token"

	"github.com/fsnotify/fsnotify"
)

// TODO: new section

type (
	Ini struct {
		currectSection *ast.SetcionNode
		src            []byte // source
		l              *lexer.Lexer
		p              *parser.Parser
		doc            *ast.Doc
		err            error

		watcher       *fsnotify.Watcher
		exitWatchChan chan bool
	}
)

func New() *Ini {

	in := &Ini{}
	return in
}

func (in *Ini) Err() error {
	return in.err
}

func (in *Ini) LoadFile(file string) *Ini {

	// read file content
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		in.err = err
		return in
	}

	in.Load(bts)

	return in
}

func (in *Ini) WatchFile(file string) *Ini {

	in.LoadFile(file)
	in.watch(file)

	return in
}

func (in *Ini) watch(file string) {
	if file == "" {
		return
	}

	in.watcher, in.err = fsnotify.NewWatcher()
	in.exitWatchChan = make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-in.watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					in.LoadFile(file)
				}
			case _, ok := <-in.watcher.Errors:
				if !ok {
					return
				}
			case <-in.exitWatchChan:
				return
			}

		}
	}()

	in.err = in.watcher.Add(file)
}

func (in *Ini) StopWatch() *Ini {

	in.watcher.Close()
	in.exitWatchChan <- true
	return in
}

func (in *Ini) Load(doc []byte) *Ini {

	if len(doc) <= 0 {
		return in
	}

	in.src = doc
	in.l = lexer.New(string(in.src))
	in.p = parser.New(in.l)

	in.doc, in.err = in.p.ParseDocument()
	return in
}

func (in *Ini) Dump() {

	if in.doc == nil {
		return
	}

	in.doc.DumpV2()
}

func (this *Ini) Marshal2Map() map[string]interface{} {

	if this.doc == nil {
		return nil
	}

	if this.err != nil {
		return nil
	}

	kvMaps := make(map[string]interface{})

	for c := this.doc.FirstChild(); c != nil; c = c.NextSibling() {
		if kv_node, ok := c.(*ast.KVNode); ok {
			kvMaps[kv_node.Key.Literal] = kv_node.Value.Literal
		}

		if sect_node, ok := c.(*ast.SetcionNode); ok {

			secMap := make(map[string]interface{})

			for kv := sect_node.FirstChild(); kv != nil; kv = kv.NextSibling() {
				if kvnode, ok := kv.(*ast.KVNode); ok {
					secMap[kvnode.Key.Literal] = kvnode.Value.Literal
				}
			}

			kvMaps[sect_node.Name.Literal] = secMap
			continue
		}
	}

	return kvMaps
}

func (this *Ini) Marshal2Json() []byte {

	kvMaps := this.Marshal2Map()

	if kvMaps == nil {
		return nil
	}

	result, err := json.Marshal(kvMaps)
	this.err = err

	return result
}

func (this *Ini) Section(section string) *Ini {

	if this.doc == nil {
		return this
	}

	if this.err != nil {
		return this
	}

	this.sectionForAstDoc(section)
	return this
}

func (this *Ini) Get(key string) string {
	return this.GetDef(key, "")
}

func (this *Ini) GetDef(key string, def string) string {

	if this.doc == nil ||
		this.err != nil {
		return def
	}

	if key == "" {
		return def
	}

	tok := this.getToken(key)
	if tok.Type != token.TokenType_VALUE {
		return def
	}

	return tok.Literal
}

func (this *Ini) GetInt(key string) int {

	return this.GetIntDef(key, 0)
}

func (this *Ini) GetIntDef(key string, def int) int {

	val := this.Get(key)
	if val == "" {
		return def
	}

	ival, err := strconv.Atoi(val)
	if err != nil {
		return def
	}

	return ival
}

func (this *Ini) GetInt64(key string) int64 {

	return this.GetInt64Def(key, 0)
}

func (this *Ini) GetInt64Def(key string, def int64) int64 {

	val := this.Get(key)
	if val == "" {
		return def
	}

	ival, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return def
	}

	return ival
}

func (this *Ini) GetFloat64(key string) float64 {

	return this.GetFloat64Def(key, 0)
}

func (this *Ini) GetFloat64Def(key string, def float64) float64 {

	val := this.Get(key)
	if val == "" {
		return def
	}

	fval, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return def
	}
	return fval
}

func (this *Ini) Set(key string, val interface{}) *Ini {

	if this.doc == nil ||
		this.err != nil {
		return this
	}

	if key == "" || val == nil {
		return this
	}

	var val_str string
	switch val.(type) {
	case int:
		val_str = fmt.Sprintf("%d", val.(int))
	case int32:
		val_str = fmt.Sprintf("%d", val.(int32))
	case int64:
		val_str = fmt.Sprintf("%d", val.(int64))
	case float32:
		val_str = strconv.FormatFloat(float64(val.(float32)), 'f', -1, 32)
	case float64:
		val_str = strconv.FormatFloat(float64(val.(float64)), 'f', -1, 64)
	case string:
		val_str = val.(string)
	default:
		return this
	}

	val_str = strings.Replace(val_str, "\n", "", -1)
	val_str = strings.Replace(val_str, "\t", "", -1)
	val_str = strings.Trim(val_str, " ")

	this.setKVNode(key, val_str)

	return this

}

func (this *Ini) Del(key string) *Ini {

	if this.doc == nil ||
		this.err != nil {
		return this
	}

	if key == "" {
		return this
	}

	if this.currectSection == nil {

		for c := this.doc.FirstChild(); c != nil; c = c.NextSibling() {
			if kv_node, ok := c.(*ast.KVNode); ok {
				if kv_node.Key.Literal == key {
					this.doc.RemoveChild(this.doc, kv_node)
					break
				}
				continue
			}

			// if sect_node, ok := c.(*ast.SetcionNode); ok {
			// 	continue
			// }
		}

	} else {
		for c := this.currectSection.FirstChild(); c != nil; c = c.NextSibling() {
			kvnodev2 := c.(*ast.KVNode)
			if kvnodev2.Key.Literal == key {
				this.currectSection.RemoveChild(this.currectSection, c)
				break
			}
		}
	}

	return this
}

func (this *Ini) DelSection(section string) *Ini {
	if this.doc == nil ||
		this.err != nil {
		return this
	}

	if section == "" {
		return this
	}

	if this.currectSection == nil {

		for c := this.doc.FirstChild(); c != nil; c = c.NextSibling() {

			if sect_node, ok := c.(*ast.SetcionNode); ok {
				if sect_node.Name.Literal == section {
					this.doc.RemoveChild(this.doc, sect_node)
					break
				}
			}
		}
	}

	return this
}

// TODO: implement Save
func (this *Ini) Save(filename string) *Ini {

	if filename == "" {
		return this
	}

	if this.doc == nil {
		return this
	}

	var result string

	is_last_type_comment := false
	for c := this.doc.FirstChild(); c != nil; c = c.NextSibling() {
		if kv_node, ok := c.(*ast.KVNode); ok {
			is_last_type_comment = false
			result = fmt.Sprintf("%s%s = %v\n", result, kv_node.Key.Literal, kv_node.Value.Literal)
			continue
		}

		if sect_node, ok := c.(*ast.SetcionNode); ok {
			is_last_type_comment = false
			result = fmt.Sprintf("%s[%s]\n", result, sect_node.Name.Literal)

			for c := sect_node.FirstChild(); c != nil; c = c.NextSibling() {
				if kv_node, ok := c.(*ast.KVNode); ok {
					result = fmt.Sprintf("%s%s = %v\n", result, kv_node.Key.Literal, kv_node.Value.Literal)
					continue
				}
			}

			continue
		}

		if comm_node, ok := c.(*ast.CommentNode); ok {
			if !is_last_type_comment {
				result = fmt.Sprintf("%s\n", result)
			}

			is_last_type_comment = true
			result = fmt.Sprintf("%s%s\n", result, comm_node.Comment.Literal)
		}
	}

	os.Remove(filename)
	var data = []byte(result)
	this.err = ioutil.WriteFile(filename, data, 0666)

	return this
}

// ----------------------------------------------------------------

func (this *Ini) sectionForAstDoc(section string) {

	if this.doc == nil ||
		this.err != nil {
		return
	}

	this.currectSection = nil
	if section == "" {
		return
	}

	for c := this.doc.FirstChild(); c != nil; c = c.NextSibling() {

		if sect_node, ok := c.(*ast.SetcionNode); ok {

			if sect_node.Name.Literal == section {
				this.currectSection = sect_node
				return
			}
		}
	}
}

func (this *Ini) getToken(key string) token.Token {

	var tok token.Token

	if key == "" {
		return tok
	}

	if this.currectSection == nil {

		for c := this.doc.FirstChild(); c != nil; c = c.NextSibling() {
			if kv_node, ok := c.(*ast.KVNode); ok {
				if kv_node.Key.Literal == key {
					tok = kv_node.Value
					return tok
				}
				continue
			}
		}

	} else {
		for c := this.currectSection.FirstChild(); c != nil; c = c.NextSibling() {

			kvnodev2 := c.(*ast.KVNode)
			if kvnodev2.Key.Literal == key {
				tok = kvnodev2.Value
				return tok
			}
		}
	}

	return tok
}

func (this *Ini) setKVNode(key string, val string) *Ini {

	if key == "" || val == "" {
		return this
	}

	line := 1
	found := false

	if this.currectSection == nil {

		var last_kv_node *ast.KVNode
		for c := this.doc.FirstChild(); c != nil; c = c.NextSibling() {
			if kv_node, ok := c.(*ast.KVNode); ok {
				last_kv_node = kv_node
				if kv_node.Key.Literal == key {
					kv_node.Value.Literal = val
					return this
				}
				continue
			}

			// if sect_node, ok := c.(*ast.SetcionNode); ok {
			// 	continue
			// }
		}

		kvnode := &ast.KVNode{
			Key: token.Token{
				Type:    token.TokenType_KEY,
				Literal: key,
				Line:    line,
			},
			Value: token.Token{
				Type:    token.TokenType_VALUE,
				Literal: val,
				Line:    line,
			},
		}

		this.doc.InsertAfter(this.doc, last_kv_node, kvnode)
		this.re_adjust_node_line()

	} else {

		for c := this.currectSection.FirstChild(); c != nil; c = c.NextSibling() {

			kvnodev2 := c.(*ast.KVNode)
			line = kvnodev2.Key.Line + 1
			if kvnodev2.Key.Literal == key {
				kvnodev2.Value.Literal = val
				return this

			}
		}

		if found == false {

			kvnode := &ast.KVNode{
				Key: token.Token{
					Type:    token.TokenType_KEY,
					Literal: key,
					Line:    line,
				},
				Value: token.Token{
					Type:    token.TokenType_VALUE,
					Literal: val,
					Line:    line,
				},
			}

			this.currectSection.AppendChild(this.currectSection, kvnode)
			this.re_adjust_node_line()

		}
	}

	return this
}

// TODO: support
func (this *Ini) re_adjust_node_line() {

}
