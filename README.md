# httplus

http extension for golang

## Usage

### Delete

reference

```go
Delete(url string) (*http.Response, error)
```

example

```go
resp, err := httplus.Delete("http://api.tekito/images/1")
```

### PostFormWithFile

reference

```go
PostFormWithFile(targetUrl string,
                 valueParams map[string]string,
                 keyName string,
                 filename string) (*http.Response, error)
```

example

```go
params := map[string]string{
    "key1": "value1",
    "key2": "value2",
}
resp, err := httplus.PostFormWithFile("http://api.tekito/images", params, "file", "./tmp/sample.jpg")
```

&copy; funnythingz
