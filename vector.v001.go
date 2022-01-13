package vector

import (
	"ec-tsj/core"
	"ec-tsj/event"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type IVector interface {
	Insert(int, ...T)
	Add(...T)
	Remove(int) T
	Clone() IVector
	ToSlice(...int) []Element
	Set(int, T)
	Get(int) T
	String() string
	Reverse()
	Pop(pop) T
	Count(T) int
	Index(T) int
	Sort(order)
	IsEmpty() bool
	Size() int
	Max() T
	Min() T
}

func init() {
	Events = event.EventEmitter("Tuple", false)
}

var (
	Events event.IEventEmitter
)

type (
	// Dato
	T = core.T
	// Vector
	Element struct {
		D T
	}
	// Slice de Vectors
	Vector []Element
)

func (t Element) String() string {
	return fmt.Sprintf("%v", t.D)
}

// Crea un Objeto Vector
//
// @param {...T} t
// return {*Vector}
func NewVector(Ts ...T) IVector {
	st := &Vector{}
	s := len(Ts)

	if s != 0 {
		st.Add(Ts...)
	}

	return st
}

// Añade uno o más elementos al Vector
// Puede ser un Vector, slice de interfaces ó elementos sueltos
//
// @param {...T}
func (this *Vector) Add(Ts ...T) {
	if v, ok := Ts[0].([]T); ok {
		for _, f := range v {
			*this = append(*this, Element{f})
			Events.Emit("Add", f)
		}
	} else if v, ok := Ts[0].([]Element); ok {
		for _, f := range v {
			*this = append(*this, f)
			Events.Emit("Add", f)
		}
	} else {
		for _, v := range Ts {
			*this = append(*this, Element{v})
			Events.Emit("Add", v)
		}
	}
}

// !+

func (this *Vector) _parms(i int) int {
	if i < 0 {
		i = 0
	}
	if i >= this.Size() {
		i = this.Size()
	}

	return i
}

// Inserta uno o más elementos al Vector, empezando en un elemento dado
// Puede ser un Vector, slice de interfaces ó elementos sueltos
//
// @param {int}
// @param {...T} t
func (this *Vector) Insert(start int, Ts ...T) {
	euler := *this
	start = this._parms(start)
	rhs := euler[start:]
	if v, ok := Ts[0].([]T); ok {
		for _, f := range v {
			euler = append(euler, Element{f})
			Events.Emit("Add", f)
		}
	} else if v, ok := Ts[0].([]Element); ok {
		for _, f := range v {
			euler = append(euler, f)
			Events.Emit("Add", f)
		}
	} else {
		for _, v := range Ts {
			euler = append(euler, Element{v})
			Events.Emit("Add", v)
		}
	}
	euler = append(euler, rhs...)
	*this = euler
}

// Borra un elemento dado del Vector
//
// @param {int} t
// return {T}
func (this *Vector) Remove(T int) T {
	euler := *this
	T = this._parms(T)
	aser := euler[T]
	euler = append(euler[:T], euler[T+1:]...)
	*this = euler
	Events.Emit("Remove", T)

	return aser.D
}

//!-

// Hace una copia del Vector subyacente
//
// return {*Vector}
func (this *Vector) Clone() IVector {
	t := NewVector()
	v := t.(*Vector)
	*v = append(*v, *this...)

	return v
}

// Retorna un slice de los datos del Vector. Base 0.
//
// @param {blanco|int,int} first, last ó en blanco
// return {[]T}
func (this *Vector) ToSlice(parms ...int) []Element {
	copy := this.Clone()
	fcopy := copy.(*Vector)

	parms = append(parms, -1, -1)
	parms[0] = map[bool]int{true: parms[0], false: 0}[parms[0] != -1]
	parms[1] = map[bool]int{true: parms[1], false: len(*fcopy)}[parms[1] != -1]

	euler := *fcopy
	euler = euler[parms[0]:parms[1]]

	return euler
}

// Pone el dato n en el Vector. Base 0
//
// @param {int}
// @param {T} t
func (this Vector) Set(idx int, T T) {
	if idx < 0 {
		idx = 0
	} else if idx >= this.Size() {
		this = append(this, Element{T}) // Añade otro T
		return
	}
	this[idx] = Element{T} // modifica un T
}

// Obtiene el dato n del Vector. Base 0
//
// @param {int} t
// return {T}
func (this Vector) Get(idx int) T {
	if idx < 0 {
		idx = 0
	} else if idx >= this.Size() {
		idx = this.Size() - 1 // idx es el máximo
	}
	return this[idx].D
}

// Retorna representación en cadena del Vector
//
// return {string}
func (this *Vector) String() string {
	Ts := make([]string, 0, len(*this))
	for _, val := range *this {
		if fmt.Sprintf("%T", val) == "tuple.Tuple" {
			Ts = append(Ts, fmt.Sprintf("%v", val.D))
		} else {
			Ts = append(Ts, fmt.Sprintf("%v", val))
		}
	}

	return fmt.Sprintf("(%s)", strings.Join(Ts, ", "))
}

// Reversa el Vector subyacente
func (this Vector) Reverse() {
	for i, j := 0, this.Size()-1; i < j; i, j = i+1, j-1 {
		this[i], this[j] = this[j], this[i]
	}
}

// Pop el T más a la izquierda/derecha del Vector y lo retorna
//
// @param {pop} t
// return {T}
func (this *Vector) Pop(enum pop) (ret T) {
	euler := *this
	if enum == POP_LEFT {
		ret = euler[0]
		*this = euler[1:] // recorta el primero
	} else if enum == POP_RIGHT {
		ret = euler[this.Size()-1]
		*this = euler[:this.Size()-1] // recorta el ultimo
	}

	return
}

// Cuenta el número de veces que aparece un objeto en el Vector
// Típicode Python
//
// @param {T} t
// return {int}
func (this *Vector) Count(value T) (count int) {
	for _, m := range *this {
		if m == value {
			count++
		}
	}

	return
}

// Nos dice el sitio, más bajo, en que aparece. Típico de Python
//
// @param {T} t
// return {int}
func (this *Vector) Index(value T) int {
	var index int = -1

	for n, v := range *this {
		if v == value {
			index = n
			break
		}
	}

	return index
}

// IsEmpty retorna true si el Vector está vacío
//
// return {bool}
func (this *Vector) IsEmpty() bool {
	return this.Size() == 0
}

// Size retorna el tamaño del Vector
//
// return {int}
func (this *Vector) Size() int {
	return len(*this)
}

// Devuelve el elemento máximo de un tuplePtr
//
// @param {Vector}
// return {T}
func (this *Vector) Max() (max T) {
	var maxA, maxB, maxC = 0, nullString, 0.0

	for _, m := range *this {
		switch g := m.D.(type) {
		case string:
			if g > maxB {
				max, maxB = g, g
				maxA, _ = strconv.Atoi(g)
				maxC, _ = strconv.ParseFloat(g, 64)
			}
		case int:
			if g > maxA {
				max, maxA = g, g
				maxB = strconv.Itoa(g)
				maxC = float64(g)
			}
		case float64:
			if g > maxC {
				max, maxC = g, g
				maxA = int(g)
				maxB = fmt.Sprintf("%f", g)
			}
		}
	}

	return
}

// Devuelve el elemento mínimo de un tuplePtr
//
// @param {Vector} t
// return {T}
func (this *Vector) Min() (min T) {
	var minA, minB, minC = math.MaxInt64, "�", math.Inf(1)

	for _, m := range *this {
		switch g := m.D.(type) {
		case string:
			if g < minB {
				min, minB = g, g
				minA, _ = strconv.Atoi(g)
				minC, _ = strconv.ParseFloat(g, 64)
			}
		case int:
			if g < minA {
				min, minA = g, g
				minB = strconv.Itoa(g)
				minC = float64(g)
			}
		case float64:
			if g < minC {
				min, minC = g, g
				minA = int(g)
				minB = fmt.Sprintf("%f", g)
			}
		}
	}

	return
}

// Clasifica los elementos por orden
//
// @param {order} t
func (this *Vector) Sort(enum order) {
	__enum = enum
	sort.Sort(data(*this))
}
