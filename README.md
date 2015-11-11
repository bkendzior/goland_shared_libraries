## Calling golang from ruby

Compile golang code using: 

```go build -buildmode=c-shared -o goCallHTTP.so goCallHTTP.go```

Now execute your ruby script:

```ruby goFromRuby.rb ```

Whala! (Note, requires Golang version 1.5)
