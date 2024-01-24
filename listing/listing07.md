Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
3
4
5
6
7
8
0
0
...
```

Проблема здесь в том, что при закрытии каналов `a` и `b`, горутина в функции merge продолжает читать из этих каналов, и она получает значения по умолчанию для типа `int` после их закрытия, что является 0.

После закрытия канала в него можно продолжать писать, и значения, записанные после закрытия, будут считываться как нули.

В go нет способа проверить закрыт ли канал не вычитывая из него значения, поэтому нужно применить другой подход для решения этой задачи:

```go
func merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			out <- v
		}
		wg.Done()
	}(a)
	go func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			out <- v
		}
	}(b)
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
```

Либо сделать так - 

```go 
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c) // Закрываем канал c при выходе из горутины
		for {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil // Устанавливаем a в nil, чтобы игнорировать его в дальнейших чтениях
				} else {
					c <- v
				}
			case v, ok := <-b:
				if !ok {
					b = nil // Устанавливаем b в nil, чтобы игнорировать его в дальнейших чтениях
				} else {
					c <- v
				}
			}

			// Завершаем цикл, если оба канала закрыты
			if a == nil && b == nil {
				return
			}
		}
	}()
	return c
}

```
