package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type Token interface{}
type Symbol struct{Name string}
type String struct{Text string}
type Int struct{Num int64}
type StartList struct{}
type EndList struct{}


type TokenDecoder struct {
	lexer *lexer
}

func NewTokenDecoder(r io.Reader) *TokenDecoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	return &TokenDecoder{lexer: lex}
}

func (dec *TokenDecoder) Token() (Token, error) {
	switch dec.lexer.token {
	case scanner.Ident:
		name := dec.lexer.text()
		dec.lexer.next()
		return Symbol{Name: name}, nil
	case scanner.String:
		text, err := strconv.Unquote(dec.lexer.text())
		if err != nil {
			return nil, err
		}
		dec.lexer.next()
		return String{Text: text}, nil
	case scanner.Int:
		i, err := strconv.ParseInt(dec.lexer.text(), 10, 64)
		if err != nil {
			return nil, err
		}
		dec.lexer.next()
		return Int{Num: i}, nil
	case '(':
		dec.lexer.next()
		return StartList{}, nil
	case ')':
		dec.lexer.next()
		return EndList{}, nil
	}
	panic(fmt.Sprintf("unexpected token %q", dec.lexer.text()))
}

func (dec *TokenDecoder) Decode(v interface{}) (err error) {
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
	return NewTokenDecoder(bytes.NewReader(data)).Decode(out)
}