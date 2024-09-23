Run spansql test cases in memefish

```
go test ./spanner/spansql > output/result.txt
```

update test cases
```
go install github.com/asty-org/asty@f7439dbcce4dffbb79f068009aed9f50b95039a8
asty go2json -input ./google-cloud-go/spanner/spansql/sql_test.go | gojq -f gen_tests.jq -r  > ./spanner/spansql/tests.go
```