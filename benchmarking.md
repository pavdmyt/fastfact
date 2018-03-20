Methodology
-----------

For benchmarking the following code snippet is used (calculating factorial of 1 million):

```go
// (10^6)! = 8.263931688×10^5565708
package main

import (
	"fmt"
	"math/big"
	"runtime"
	"time"

	"github.com/pavdmyt/fastfact"
)

const value = 1e6

func main() {
	fmt.Println("Logical CPUs: ", runtime.NumCPU())
	fmt.Println("Value to calculate: ", value)
	fmt.Printf("---\n\n")

	funcsMap := map[string]func(n uint64) *big.Int{
		"SimpleFactIter": fastfact.SimpleFactIter,
		"SimpleFactFast": fastfact.SimpleFactFast,
		"HalfIterFact":   fastfact.HalfIterFact,
	}

	for fnName, fn := range funcsMap {
		fmt.Printf("%v -- ", fnName)
		start := time.Now()
		fn(value)
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("elapsed time: %v \n", elapsed)
	}

	// Concurrent funcs
	cfuncsMap := map[string]func(n, w uint64) *big.Int{
		"ConcFactIter": fastfact.ConcFactIter,
		"ConcFactFast": fastfact.ConcFactFast,
	}

	for fnName, fn := range cfuncsMap {
		fmt.Println()
		for w := 1; w <= 20; w++ {
			fmt.Printf("%v :: workers %v -- ", fnName, w)
			start := time.Now()
			fn(value, uint64(w))
			end := time.Now()
			elapsed := end.Sub(start)
			fmt.Printf("elapsed time: %v \n", elapsed)
		}
	}

}
```


Results
-------

2.5 GHz Intel Core i7 (MacBook Pro Mid 2015)

Mac OS High Sierra 10.13.3

```
Logical CPUs:  8
Value to calculate:  1e+06
---

SimpleFactIter -- elapsed time: 1m46.891471321s
SimpleFactFast -- elapsed time: 3.298304602s
HalfIterFact -- elapsed time: 58.44551817s

ConcFactIter :: workers 1 -- elapsed time: 1m44.651212752s
ConcFactIter :: workers 2 -- elapsed time: 32.643371805s
ConcFactIter :: workers 3 -- elapsed time: 17.349941887s
ConcFactIter :: workers 4 -- elapsed time: 12.212981866s
ConcFactIter :: workers 5 -- elapsed time: 10.13304031s
ConcFactIter :: workers 6 -- elapsed time: 8.777945971s
ConcFactIter :: workers 7 -- elapsed time: 7.994894315s
ConcFactIter :: workers 8 -- elapsed time: 7.319136241s
ConcFactIter :: workers 9 -- elapsed time: 7.06090162s
ConcFactIter :: workers 10 -- elapsed time: 6.693791224s
ConcFactIter :: workers 11 -- elapsed time: 6.710252179s
ConcFactIter :: workers 12 -- elapsed time: 6.290332724s
ConcFactIter :: workers 13 -- elapsed time: 6.370237167s
ConcFactIter :: workers 14 -- elapsed time: 6.440249234s
ConcFactIter :: workers 15 -- elapsed time: 6.104821685s
ConcFactIter :: workers 16 -- elapsed time: 6.078660806s
ConcFactIter :: workers 17 -- elapsed time: 6.206644528s
ConcFactIter :: workers 18 -- elapsed time: 6.020823198s
ConcFactIter :: workers 19 -- elapsed time: 5.918030852s
ConcFactIter :: workers 20 -- elapsed time: 6.098986628s

ConcFactFast :: workers 1 -- elapsed time: 3.085179362s
ConcFactFast :: workers 2 -- elapsed time: 2.072858126s
ConcFactFast :: workers 3 -- elapsed time: 2.230860195s
ConcFactFast :: workers 4 -- elapsed time: 2.347110989s
ConcFactFast :: workers 5 -- elapsed time: 2.536107528s
ConcFactFast :: workers 6 -- elapsed time: 2.709673992s
ConcFactFast :: workers 7 -- elapsed time: 2.930657447s
ConcFactFast :: workers 8 -- elapsed time: 3.135572134s
ConcFactFast :: workers 9 -- elapsed time: 3.253768263s
ConcFactFast :: workers 10 -- elapsed time: 3.285120717s
ConcFactFast :: workers 11 -- elapsed time: 3.581594355s
ConcFactFast :: workers 12 -- elapsed time: 3.737157149s
ConcFactFast :: workers 13 -- elapsed time: 3.818365675s
ConcFactFast :: workers 14 -- elapsed time: 4.004657817s
ConcFactFast :: workers 15 -- elapsed time: 3.998096669s
ConcFactFast :: workers 16 -- elapsed time: 4.190567967s
ConcFactFast :: workers 17 -- elapsed time: 4.232954667s
ConcFactFast :: workers 18 -- elapsed time: 4.337561912s
ConcFactFast :: workers 19 -- elapsed time: 4.392707296s
ConcFactFast :: workers 20 -- elapsed time: 4.49191385s
```


3.7GHz Intel Core i3-4170

Debian GNU Linux (9 "stretch")

```
Logical CPUs:  4
Value to calculate:  1e+06
---

SimpleFactFast -- elapsed time: 2.909772092s
HalfIterFact -- elapsed time: 1m8.50245915s
SimpleFactIter -- elapsed time: 1m57.486627904s

ConcFactIter :: workers 1 -- elapsed time: 1m54.401843303s
ConcFactIter :: workers 2 -- elapsed time: 42.335188303s
ConcFactIter :: workers 3 -- elapsed time: 22.119973197s
ConcFactIter :: workers 4 -- elapsed time: 16.075372869s
ConcFactIter :: workers 5 -- elapsed time: 12.927340156s
ConcFactIter :: workers 6 -- elapsed time: 11.350166285s
ConcFactIter :: workers 7 -- elapsed time: 10.099888116s
ConcFactIter :: workers 8 -- elapsed time: 9.378428897s
ConcFactIter :: workers 9 -- elapsed time: 8.872771431s
ConcFactIter :: workers 10 -- elapsed time: 8.683862973s
ConcFactIter :: workers 11 -- elapsed time: 8.007421478s
ConcFactIter :: workers 12 -- elapsed time: 7.754752209s
ConcFactIter :: workers 13 -- elapsed time: 7.499564317s
ConcFactIter :: workers 14 -- elapsed time: 7.589316523s
ConcFactIter :: workers 15 -- elapsed time: 7.143799068s
ConcFactIter :: workers 16 -- elapsed time: 7.239254828s
ConcFactIter :: workers 17 -- elapsed time: 7.042676556s
ConcFactIter :: workers 18 -- elapsed time: 6.982440868s
ConcFactIter :: workers 19 -- elapsed time: 6.888201532s
ConcFactIter :: workers 20 -- elapsed time: 6.774055605s

ConcFactFast :: workers 1 -- elapsed time: 2.907954989s
ConcFactFast :: workers 2 -- elapsed time: 2.119570606s
ConcFactFast :: workers 3 -- elapsed time: 2.282450764s
ConcFactFast :: workers 4 -- elapsed time: 2.466834714s
ConcFactFast :: workers 5 -- elapsed time: 2.615551396s
ConcFactFast :: workers 6 -- elapsed time: 2.780724177s
ConcFactFast :: workers 7 -- elapsed time: 2.916450005s
ConcFactFast :: workers 8 -- elapsed time: 3.117211745s
ConcFactFast :: workers 9 -- elapsed time: 3.26846721s
ConcFactFast :: workers 10 -- elapsed time: 3.306619585s
ConcFactFast :: workers 11 -- elapsed time: 3.501615455s
ConcFactFast :: workers 12 -- elapsed time: 3.586054541s
ConcFactFast :: workers 13 -- elapsed time: 3.732140728s
ConcFactFast :: workers 14 -- elapsed time: 3.840676903s
ConcFactFast :: workers 15 -- elapsed time: 3.903011014s
ConcFactFast :: workers 16 -- elapsed time: 4.143965956s
ConcFactFast :: workers 17 -- elapsed time: 4.158132895s
ConcFactFast :: workers 18 -- elapsed time: 4.283548579s
ConcFactFast :: workers 19 -- elapsed time: 4.284276562s
ConcFactFast :: workers 20 -- elapsed time: 4.385495799s
```
