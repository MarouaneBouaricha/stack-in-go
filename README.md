# stack-in-go
A simple stack implementation in the go language

func main() {
    var s Stack
    s.Push(34)
    s.Push(17)
    fmt.Printf("stack %v\n", s)
}

output:

stack [0:34][1:17][2:0]
