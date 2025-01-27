###
Strings:

String are represented by bytes(uint8) array.

A rune is a character in go. It is represented by int32.

A String in Go is UTF-8 encoded, instead of ASCII like in other programming languages.

```go
package main

import . "fmt"

func main() {
    s := élite
    
    // note here that the string is UTF-8 encoded and represented by bytes, hence the length of the string is 6
    Printf("%8T %[1]v %d\n", s, len(s)) // output: string élite 6  // %8T is type; %[1]v represents to use s as input for %v; %d is for integers, used here to print length of string
    Printf("%8T %[1]v %d\n", []rune(s)) // output: []int32 [233 108 105 116 101]
    
    b := []byte(s)
    Printf("%8T %[1]v %d\n", b, len(b)) // output: []uint8 [195 169 108 105 116 101] 6
}