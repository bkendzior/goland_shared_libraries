// Requires golang 1.5
// go build -buildmode=c-shared -o goCallHTTP.so goCallHTTP.go

package main

import "C"
//import "fmt"
import "net/http"

//export GoHTTP
func GoHTTP(url *C.char) *C.char {
  resp, err := http.Get(C.GoString(url))

  if err != nil {
    return C.CString("Invalid response")
  }

//  fmt.Println("%s", resp.Status)

  return C.CString(resp.Status)
}

func main() {} // Required but ignored
