package main
import "fmt"
type data struct{
    x, y int
}

func main(){
    var p *data = new(data)
    p.x, p.y = 1, 2
    fmt.Println(p)
    pp := new(data)
    pp.x, pp.y = 3, 4
    fmt.Println(pp)
}
