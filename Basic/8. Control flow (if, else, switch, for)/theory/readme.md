- Counter
  - `i := 2` - инициализация счетчика
  - `i < 10` - условие, при котором будет выполняться цикл
  - `i++` - последействие, которое выполняется после тела цикла

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 2; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

- Без счетчика, только условие:

```go
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

- Бесконечный цикл
  В Go нет своего `while`, но его функционал можно реализовать с помощью `for`:

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```
