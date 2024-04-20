# Pointers in Go

Pointers are variables that store the address of a value in memory. They are used to pass values by reference. In Go, pointers are represented using the `*` symbol.

## Why Pointers

- avoid value copying
- directly modify the value

By default in Go when you pass a value to a function it is a copy of the value. If you want to modify the original value you need to pass a pointer to the value. For large and complex values in a program this could be a performance issue.
