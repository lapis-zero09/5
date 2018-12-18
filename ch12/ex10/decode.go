package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

//!+lexer
type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

//!-lexer

//!+read
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
		if lex.text() == "t" {
			v.SetBool(true)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(f)
		lex.next()
		return
	case '-':
		lex.consume('-')
		switch lex.token {
		case scanner.Int:
			i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
			v.SetInt(int64(i))
		case scanner.Float:
			f, _ := strconv.ParseFloat(lex.text(), 64)
			v.SetFloat(f)
		default:
			panic(fmt.Sprintf("unexpected token %q", lex.text()))
		}
		lex.next()
		return
	case '#':
		lex.consume('#')
		lex.consume(scanner.Ident)
		lex.consume('(')
		r, _ := strconv.ParseFloat(lex.text(), 64)
		lex.next()
		i, _ := strconv.ParseFloat(lex.text(), 64)
		lex.next()
		lex.consume(')')
		v.SetComplex(complex(r, i))
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

//!-read

//!+readlist
func readList(lex *lexer, v reflect.Value) {
	log.Println(v.Kind())
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	case reflect.Interface:
		if lex.token != scanner.String {
			panic(fmt.Sprintf("got token %q", lex.text()))
		}
		t, _ := strconv.Unquote(lex.text())
		lex.next()
		val := reflect.New(getType(t).Elem())
		read(lex, val)
		v.Set(val)

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func getType(t string) reflect.Type {
	types := map[string]reflect.Type{
		"int":         reflect.ValueOf(int(0)).Type(),
		"int8":        reflect.ValueOf(int8(0)).Type(),
		"int16":       reflect.ValueOf(int16(0)).Type(),
		"int32":       reflect.ValueOf(int32(0)).Type(),
		"int64":       reflect.ValueOf(int64(0)).Type(),
		"uint":        reflect.ValueOf(uint(0)).Type(),
		"uint8":       reflect.ValueOf(uint8(0)).Type(),
		"uint16":      reflect.ValueOf(uint16(0)).Type(),
		"uint32":      reflect.ValueOf(uint32(0)).Type(),
		"uint64":      reflect.ValueOf(uint64(0)).Type(),
		"float32":     reflect.ValueOf(float32(0)).Type(),
		"float64":     reflect.ValueOf(float64(0)).Type(),
		"complex64":   reflect.ValueOf(complex64(0)).Type(),
		"complex128":  reflect.ValueOf(complex128(0)).Type(),
		"byte":        reflect.ValueOf(byte(0)).Type(),
		"rune":        reflect.ValueOf(rune(0)).Type(),
		"string":      reflect.ValueOf("").Type(),
		"interface{}": reflect.TypeOf([]interface{}{}).Elem(),
	}
	if val, ok := types[t]; ok {
		return val
	}

	if strings.HasPrefix(t, "[]") {
		return reflect.SliceOf(getType(t[2:]))
	}
	if strings.HasPrefix(t, "[") {
		i := strings.IndexByte(t, ']')
		if i < 0 {
			panic("invalid type: " + t)
		}
		size, err := strconv.ParseInt(t[1:i], 10, 64)
		if err != nil || size < 0 {
			panic("invalid array size: " + t[1:i])
		}
		return reflect.ArrayOf(int(size), getType(t[i+1:]))
	}
	if strings.HasPrefix(t, "map[") {
		cnt := 1
		p := 4
		for i, r := range t[4:] {
			if r == '[' {
				cnt++
			}
			if r == ']' {
				cnt--
				if cnt == 0 {
					p += i
					break
				}
			}
		}
		return reflect.MapOf(getType(t[4:p]), getType(t[p+1:]))
	}

	panic("unsupported type: " + t)
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

//!-readlist

type Decoder struct {
	lexer *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	return &Decoder{lexer: lex}
}

func (dec *Decoder) Decode(v interface{}) (err error) {
	dec.lexer.next()
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.lexer.scan.Position, x)
		}
	}()
	read(dec.lexer, reflect.ValueOf(v).Elem())
	return nil
}

func Unmarshal(data []byte, out interface{}) error {
	return NewDecoder(bytes.NewReader(data)).Decode(out)
}
