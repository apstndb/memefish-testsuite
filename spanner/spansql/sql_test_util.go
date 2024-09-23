/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spansql

import (
	"github.com/cloudspannerecosystem/memefish"
	"github.com/cloudspannerecosystem/memefish/token"
)

func boolAddr(b bool) *bool {
	return &b
}

func stringAddr(s string) *string {
	return &s
}

func intAddr(i int) *int {
	return &i
}

var reparseDDL = func(s string) (interface{}, error) {
	p := &memefish.Parser{Lexer: &memefish.Lexer{File: &token.File{
		Buffer: s,
	}}}
	ddl, err := p.ParseDDL()
	if err != nil {
		return nil, err
	}
	return ddl, nil
}
var reparseDML = func(s string) (interface{}, error) {
	p := &memefish.Parser{Lexer: &memefish.Lexer{File: &token.File{
		Buffer: s,
	}}}
	dml, err := p.ParseDML()
	if err != nil {
		return nil, err
	}
	return dml, nil
}
var reparseQuery = func(s string) (interface{}, error) {
	p := &memefish.Parser{Lexer: &memefish.Lexer{File: &token.File{
		Buffer: s,
	}}}
	q, err := p.ParseQuery()
	return q, err
}
var reparseExpr = func(s string) (interface{}, error) {
	p := &memefish.Parser{Lexer: &memefish.Lexer{File: &token.File{
		Buffer: s,
	}}}
	e, pe := p.ParseExpr()
	if pe != nil {
		return nil, pe
	}
	return e, nil
}
