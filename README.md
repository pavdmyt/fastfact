# fastfact [![godoc][godoc-image]][godoc-url]

Few different implementations of the functions to calculate a [factorial](https://en.wikipedia.org/wiki/Factorial) of the given number. Utilizes [math/big](https://golang.org/pkg/math/big/) to calculate factorials of *really* big numbers.

There are also functions which are able to split calculations between [goroutines](https://tour.golang.org/concurrency/1) (workers) and do work concurrently (in parallel at multi-core machines) thus speeding up calculations.

Performance results can be viewed in [benchmarking.md](https://github.com/pavdmyt/fastfact/blob/master/benchmarking.md).


## Installation

```bash
go get -u github.com/pavdmyt/fastfact
```


## API

See [godoc reference](https://godoc.org/github.com/pavdmyt/fastfact) for detailed API documentation.


## License

MIT - Pavlo Dmytrenko; https://twitter.com/pavdmyt


[godoc-image]: https://godoc.org/github.com/pavdmyt/fastfact?status.svg
[godoc-url]: https://godoc.org/github.com/pavdmyt/fastfact
