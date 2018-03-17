Methodology
-----------

For benchmarking the following code snippet is used:

```
package main

import (
	"github.com/pavdmyt/fastfact"
)

func main() {
	factFunc := fastfact.IterFact
	factFunc(1e6)
}

```


Results
-------

## MulRangeFact implementation

```
time go run main.go
go run main.go  3.30s user 0.27s system 104% cpu 3.416 total
```


## IterFact

```
time go run main.go
go run main.go  108.76s user 3.61s system 108% cpu 1:44.01 total
```
