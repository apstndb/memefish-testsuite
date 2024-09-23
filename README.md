Run spansql test cases in memefish

```
go test ./spanner/spansql > output/result.txt
```

update test cases
```
go install github.com/asty-org/asty@f7439dbcce4dffbb79f068009aed9f50b95039a8
go install github.com/segmentio/golines@latest

git submodule update

asty go2json -input ./google-cloud-go/spanner/spansql/sql_test.go | jq -f gen_tests.jq -r | asty json2go | golines -m 20 > ./spanner/spansql/testcases_test.go
```