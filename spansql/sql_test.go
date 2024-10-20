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
	"testing"
)

func toReparse[T any](f func(filepath, s string) (T, error)) func(s string) (interface{}, error) {
	return func(s string) (interface{}, error) {
		return f("", s)
	}
}

var (
	reparseExpr  = toReparse(memefish.ParseExpr)
	reparseQuery = toReparse(memefish.ParseQuery)
	reparseDML   = toReparse(memefish.ParseDML)
	reparseDDL   = toReparse(memefish.ParseDDL)
)

func TestSQL(t *testing.T) {
	for _, test := range tests {
		// As a confidence check, confirm that parsing the SQL produces the original input.
		_, err := test.reparse(test.sql)
		if err != nil {
			t.Errorf("Reparsing %q: %v", test.sql, err)
		}
	}
}
