bars
====

A Go library to generate simple character based graphs inspired by 
[Spark](https://github.com/holman/spark/blob/master/spark).


Use
---

The ```MakeBar``` function is the foundation of the package and is the function
that generates the graph. ```MakeBar``` will take two arguments. The first one 
is the set of values to generate the graph from. The second argument must be an 
instance of the ```BarSet``` struct. Each ```BarSet``` instance holds the runes 
that will be used to render the graph.

```go
type BarSet struct {
    // Near zero values
    Zero rune
    // Values from 1/16 to 3/16
    Eighth rune
    // Values from 3/16 to 5/16
    Quarter rune
    // Values from 5/16 to 7/16
    ThreeEiths rune
    // Values from 7/16 to 9/16
    Half rune
    // Values from 9/16 to 11/16
    FiveEights rune
    // Values from 11/16 to 13/16
    ThreeQuarters rune
    // Values from 13/16 to 15/16
    SevenEights rune
    // Values near One
    One rune
}
```
Two instances of ```BarSet``` are provided. ```NiceBarSet``` will display the 
graph with ``` ▁▂▃▄▅▆▇█```, while ```BraileBarSet``` will display the graph
with ``` ⣀⣤⣶⣿```.

You can, of course, create your own graphs with other characters by creating
a new instance of ```BarSet``` with the characters you desire.


Example
-------

Here is a very simple example of how to use the library. It creates a program
very similar to Spark.

```go
package main

import (
    "fmt"
    "github.com/maurodec/bars"
    "os"
    "strconv"
)

func main() {
    args := os.Args[1:]

    numbers := make([]float64, len(args))
    for i := 0; i < len(numbers); i++ {
        if n, err := strconv.ParseFloat(args[i], 64); err == nil {
            numbers[i] = n
        } else {
            fmt.Printf("Invalid argument: %s\n", args[i])
            os.Exit(1)
        }
    }

    lines := bars.MakeBar(numbers, bars.NiceBarSet)
    fmt.Println(string(lines))
}
```

Run it:

```
go run fauxspark.go 3 2 2 1 4 8 3 3 9 4 4 7 6 2

> "▂▁▁ ▃▇▂▂█▃▃▆▅▁"
```

NOTE: Your browser may display some of the characters in a slightly different
manner than you browser.
