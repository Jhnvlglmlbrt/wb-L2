Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Интерфейс равен `nil`, только если и тип, и значение равны `nil`.
Переменная `err` имеет тип `error` (интерфейс), но хранит значение типа `*customError` (указатель на customError). 
Даже если указатель указывает на nil, сама переменная err все равно не является nil, потому что у нее есть тип (интерфейс error), 
и она не содержит nil значения этого типа.

Функция `test()` возвращает `nil`, который удовлетворяет интерфейсу `error`. В данном случае, тип возвращаемого значения функции `test()` - `*customError`, являющийся указателем на структуру `customError`, реализующую интерфейс `error`. Даже если внутренний указатель равен nil, переменная `err` не будет равна nil, потому что она содержит значение интерфейса `error` (в данном случае, указатель на customError). В результате условие if err != nil будет истинным, и программа выведет "error".