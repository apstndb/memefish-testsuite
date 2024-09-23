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
	"testing"
)

func reparseDDL(s string) (interface{}, error) {
	p := &memefish.Parser{Lexer: &memefish.Lexer{File: &token.File{
		Buffer: s,
	}}}
	ddl, err := p.ParseDDL()
	if err != nil {
		return nil, err
	}
	return ddl, nil
}

func reparseDML(s string) (interface{}, error) {
	p := &memefish.Parser{Lexer: &memefish.Lexer{File: &token.File{
		Buffer: s,
	}}}
	dml, err := p.ParseDML()
	if err != nil {
		return nil, err
	}
	return dml, nil
}

func reparseQuery(s string) (interface{}, error) {
	p := &memefish.Parser{Lexer: &memefish.Lexer{File: &token.File{
		Buffer: s,
	}}}
	q, err := p.ParseQuery()
	return q, err
}

func reparseExpr(s string) (interface{}, error) {
	p := &memefish.Parser{Lexer: &memefish.Lexer{File: &token.File{
		Buffer: s,
	}}}
	e, pe := p.ParseExpr()
	if pe != nil {
		return nil, pe
	}
	return e, nil
}

func TestSQL(t *testing.T) {

	/*
		latz, err := time.LoadLocation("America/Los_Angeles")
		if err != nil {
			t.Fatalf("Loading Los Angeles time zone info: %v", err)
		}

		line := func(n int) Position { return Position{Line: n} }
	*/

	for _, test := range tests {
		// As a confidence check, confirm that parsing the SQL produces the original input.
		_, err := test.reparse(test.sql)
		if err != nil {
			t.Errorf("Reparsing %q: %v", test.sql, err)
			continue
		}
	}
}
