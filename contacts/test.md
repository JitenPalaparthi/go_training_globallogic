# to run tests

- to run all tests from the directory
  ```go test ./...```

- to run all tests from a package
  ```go test -timeout 30s models```

- to run all tests from a file
  ```go test models/contact_test.go```

- to run one specific test from a package
```go test -timeout 30s models/contact_test.go  -run ^TestValidateContactPass```

- to run few tests from a file
```go test -timeout 30s -run ^(TestValidateContactPass|TestValidateContactFail)```