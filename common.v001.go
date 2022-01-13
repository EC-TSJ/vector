/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package vector

import (
	"strconv"
)

// engarce para los eventos
var (
	__enum order
)

type (
	//enumeraciones
	pop   int
	order int
	// Interface Sort
	data []Element
)

const (
	POP_LEFT pop = iota
	POP_RIGHT
)

const (
	ASCENDANT order = iota
	DESCENDANT
)

//!+

// Size
func (p data) Len() int { return len(p) }

// Less
func (p data) Less(i, j int) bool {
	var a, b string

	//Convierte valores a string
	if v, ok := p[i].D.(string); ok {
		a = v
	}
	if v, ok := p[i].D.(int); ok {
		a = strconv.Itoa(v)
	}
	if v, ok := p[i].D.(float64); ok {
		a = strconv.FormatFloat(v, 'E', -1, 64)
	}
	if v, ok := p[j].D.(string); ok {
		b = v
	}
	if v, ok := p[j].D.(int); ok {
		b = strconv.Itoa(v)
	}
	if v, ok := p[j].D.(float64); ok {
		b = strconv.FormatFloat(v, 'E', -1, 64)
	}

	// clasifica
	if __enum == ASCENDANT {
		return a < b
	} else if __enum == DESCENDANT {
		return a > b
	}

	return false
}

// Swap
func (p data) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//!-
