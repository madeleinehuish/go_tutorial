package main
import "os"
import "fmt"

func main() {
// os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
// You can get individual args with normal indexing.

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(len(argsWithProg))
}