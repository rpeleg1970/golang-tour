package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  var _walk func(t *tree.Tree, ch chan int)
    _walk = func(t *tree.Tree, ch chan int) {
      if(t.Left!=nil) {_walk(t.Left,ch)}
      ch<-t.Value
      if(t.Right!=nil) {_walk(t.Right,ch)}
    }
  _walk(t,ch)
  close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
ch1 := make(chan int)
       ch2 := make(chan int)
       go Walk(t1,ch1)
       go Walk(t2,ch2)

       ok1,ok2 := true,true
       v1,v2 := 0,0
       for ok1 && ok2 {
         v1,ok1 = <-ch1
         v2,ok2 = <-ch2
         if(ok1!=ok2 || v1!=v2) {
           return false
         }
     }
     return true
}

func main() {
    // Test the Walk function
    ch:= make(chan int)
    go Walk(tree.New(1), ch)
    v,ok := <-ch
    for ok {
        fmt.Printf("%d,",v)
        v,ok = <-ch
        if(!ok) {fmt.Printf("done...\n")}
    }
    
    // Test the Same function:
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(3)))
}
