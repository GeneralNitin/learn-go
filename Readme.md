## Go

### CLI Commands
```
go run <file_name1>, <file_name1>

go run .

go build 
```
### Section 1: Variables
<details>
<summary>Variables in Go</summary>

- Normal variable
    ```go
    var a int = 10
    ```
- Shorthand
    ```go
    a := 10
    ```
</details>


### Section 2: Loops
<details>
<summary>For Loop Variants</summary>

- Normal For Loop
    ```go
    for i := 0; i < 10; i++ {
        // do something
    }
    ```
- While Loop Style - For Loop
    ```go
    i := 0
    for  i < 10 {
        // do something
        i++
    }
    ```
-  Infinite Loop
    ```go
    for  {
        // do something
    }
   // OR 
    for true  {
        // do something
    }
    ```
</details>


### Section 3: Multiple Return
<details>
<summary>Return Multiple values from a Method</summary>

- Code example
    ```go
    func go() (int error) {
        return 0, nil
    }
  
    func main() {
        i, err := foo()
        if err !=nil {
            panic(err)
        } 
		println(i)
    }
    ```
</details>

### Section 4: Defer keyword

### Section 5: new vs make
    Make is only used with a slice, map and channel(chan)
    Make initializes the underlying datastuctures for above mentioned DS
    ```
        func main() {
            type TypeStruct struct {
                Name string
            }

            t := new(TestStruct) // type *TestStruct - Create an Object with Zero values of the variables
            var v TestStruct // type TestStruct
        }
    ```

### Section 6: interfaces ?
### Section 7: concurrency/waitGroups(making the main thread wait)/mutex(remove race condition)/channels(passing data to routines) ?


