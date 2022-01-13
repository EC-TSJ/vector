# Set

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

Tuple es una librería para manipular la DataStructure tuple.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Funciones Principales
--- 


Tiene los objetos siguientes:
- Type `struct` ***Tuple***, para hacer las definiciones.

- Type `[]Tuple` ***STuple*** con los métodos de ITuple.

- Type `ìnterface` ***ITuple***, con métodos:
  - *`Insert(int, ...Item)`*	
  - *`IsEmpty() bool`*
  - *`Size() int`*
  - *`Add(...Item)`*
  - *`Remove(int) Item`*
  - *`Clone() ITuple`*
  - *`ToSlice(...int) []Tuple`*
  - *`Set(int, Item)`*
  - *`Get(int) Item`*
  - *`String() string`*
  - *`Reverse()`*
  - *`Pop(pop) Item`*
  - *`Count(Item) int`*
  - *`Index(Item) int`*
  - *`Sort(order)`*
  - *`Max() Item`*
  - *`Min() Item`*	


- Funciones:

	- *`NewTuple(...Item) ITuple`*
	- *`Len(t STuple) int `*
	- *`Max(t STuple) (max Item)`* 
	- *`Min(t STuple) (min Item) `*


-Enumeraciones:

- ***pop***:

  - *`POP_LEFT`*
  -	*`POP_RIGHT`*

- ***order***:

	- *`ASCENDANT`*
	- *`DESCENDANT`*



## Ejemplos
```go

 	-ejemplo 1.-

	teb := tpl.New([]interface{}{9, 8, 61, 7, "jesus", 5, 56, "lopez", 32, 15})
	teb.Sort(tpl.ASCENDANT)
	zm := teb.Copy()
	teb.Set(teb.Len()-1, 25)
	teb.Set(125, 135)
	teb.Set(3, 25)
	teb.Remove(5)
	teb.InsertItems(5, 53, "venticinco", 79)
	teb.AddItems(52, "venticuatro", 78)
	teb.Add(teb)
	scb := teb.Get(131)
	sal := teb.String()
	teb.Reverse()
	zl := teb.ToSlice(3, 8)
	zk := teb.ToSlice()
	left := teb.Pop(tpl.POP_LEFT)
	right := teb.Pop(tpl.POP_RIGHT)
	za := teb.Count(25)
	zs := teb.Index(11)
	zd := teb.Index(25)
	tuple := tpl.Tuple([]interface{}{9, 8, 7, 6, 5, 56, 21, 32, 15})
	size := tpl.Len(tuple)
	max := tpl.Max(tuple)
	min := tpl.Min(tuple)

	fmt.Println(teb, scb, sal, zl, zk, zm, left, right, za, zs, zd)
	fmt.Println(tuple, size, max, min)

	-ejemplo 2.-

	var listA tuple.STuple
	listA = []Tuple{{D: "unp"}, {D: 4}, {D: 56}, {D: "kkk"}}
	lis := STuple([]Tuple{{D: "unp"}, {D: 4}, {D: 56}, {D: "kkk"}})
	list := NewTuple([]Tuple{{D: "unp"}, {D: 4}, {D: 56}, {D: "kkk"}})
	list.Add("dos", "tres", 9393, 450459403)
	salva := listA.Get(2)
	listA.Set(2, "gilipo")
	calling := list.String()
	list.Reverse()
	pop := list.Pop(tuple.POP_LEFT)
	popA := list.Pop(tuple.POP_RIGHT)
	list.Sort(tuple.ASCENDANT)
	listo := listA.Clone()
	golo := listA.ToSlice()
	golo2 := lis.ToSlice(2, 3)
	sqs := listA.Remove(3)
	list.Insert(2, []tuple.Tuple{{"uno"}, {"dos"}, {"tres"}, {"cuatro"}, {"cinco"}, {"seis"}, {"siete"}})
	vario := tuple.Len([]tuple.Tuple{{"uno"}, {"dos"}, {"tres"}, {"cuatro"}, {"cinco"}, {"seis"}, {"siete"}})
	max := tuple.Max([]tuple.Tuple{{"uno"}, {"dos"}, {"tres"}, {"cuatro"}, {"cinco"}, {"seis"}, {"siete"}})
	min := tuple.Min([]tuple.Tuple{{"uno"}, {"dos"}, {"tres"}, {"cuatro"}, {"cinco"}, {"seis"}, {"siete"}})
	fmt.Println(vario, list, listA, salva, calling, lis, pop, popA, listo, golo, golo2, sqs, max, min)


```
## Notas





<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->
## LICENSE

**[MIT](LICENSE)**
