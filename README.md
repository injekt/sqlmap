Install
-------

  go get github.com/injekt/sqlmap

Example
-------

```go
func main() {
  // initialize a db

  results, err := sqlmap.Query(db, "SELECT col1, col2 FROM mytable LIMIT 5")
  for _, r := range results {
    // r is a sqlmap.Result
    fmt.Println(r)
  }
}
```