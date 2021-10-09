# golang_channel_example_with_equivalent_binary_tree

Exercise: Equivalent Binary Trees
There can be many different binary trees with the same sequence of values stored in it. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

![](https://i.imgur.com/WC06zzk.png)

A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

This example uses the tree package, which defines the type:

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

Exercise: Equivalent Binary Trees
1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

The documentation for Tree can be found here.

## my first design

in Walk function

```golang
func Walk(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  }
  Walk(t.Left, ch)
  ch <- t.Value
  Walk(t.Right, ch)
}
```

this could be work for go through the tree

but would cause deadlock due to if the receive over taken the channel

so this channel should be close when sender finish go through all tree

so need to modify with defer like following

```golang
func Walk(t *tree.Tree, ch chan int) {
  //close channel before leave
  close(ch)
  // declare inner func walk
  var walk func(t *tree.Tree)
  walk = func(t *tree.Tree) {
    if t == nil {
      return
    }
    walk(t.Left)
    ch <- t.Value
    walk(t.Right)
  }
  // call walk
  walk(t)
}
```
in such method walk channel always alive

and will trigger close(ch) when walk function process finished

thus the Same method could be implemented by Walk function 

Idea is that we go through t1, t2 with same manner concurrently, 

if we take the sample step, the node value of t1, t2 should be same if t1 is equals t2

```golang
func Same(t1 *tree.Tree, t2 *tree.Tree) {
  // create channel for receive the data 
  ch1, ch2 :=  make(chan int), make(chan int)
  // start transverse t1, t2
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  for {
    // receive data ,status from ch1
    v1, ok1 := <- ch1
    // receive data ,status from ch2
    v2, ok2 := <- ch2
    if !ok1 || !ok2 { // ch1 or ch2 is closed, means t1 or t2 has been transversed
      return ok1 == ok2
    }
    if v1 != v2 { // some node value are not equivalent for t1, t2 with same step
      return false
    }
  }
}
```
