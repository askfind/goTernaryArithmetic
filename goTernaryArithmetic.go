package main

import (
	"fmt"
)

// Таб.1 Алфавит троичной симметричной системы счисления
// +--------------------------+-------+-------+-------+
// | Символы 	              |   -   |	  0   |	  +   |
// +--------------------------+-------+-------+-------+
// | Числа                    |  -1   |   0   |   1   |
// +--------------------------+-------+-------+-------+
// | Логика                   | false |  nil  | true  |
// +--------------------------+-------+-------+-------+

//
// TRIT Arithmetic
//

// Объявление троичных типов
type trit interface{}

// Преобразование трита в целое число
func trit2int(t trit) int8 {
	if t == true {
		return 1
	}
	if t == false {
		return -1
	}
	return 0
}

// Преобразование целое число в трит
func int2trit(i int8) trit {
	if i > 0 {
		return true
	}
	if i < 0 {
		return false
	}
	return nil
}

//  Преобразовать трит в символ '-','0','+'
func trit2symb(t trit) string {
	if t == false {
		return "-"
	} else if t == nil {
		return "0"
	} else {
		return "+"
	}
}

// Таб.2 Полусумматор тритов
// .------------------------.
// |     | -1  |  0  |  1   |
// |------------------------|
// | -1  | -+  |  -  |  0   |
// |------------------------|
// |  0  | -1  |  0  |  1   |
// |------------------------|
// |  1  |  0  |  1  |  +-  |
// .------------------------.

// Полусумматор двух тритов с переносом
func add_half_t(a trit, b trit) (c trit, carry trit) {

	if a == false && b == false {
		return true, false
	} else if a == false && b == nil {
		return false, nil
	} else if a == false && b == true {
		return nil, nil
	} else if a == nil && b == false {
		return false, nil
	} else if a == nil && b == nil {
		return nil, nil
	} else if a == nil && b == true {
		return true, nil
	} else if a == true && b == false {
		return nil, nil
	} else if a == true && b == nil {
		return true, nil
	} else if a == true && b == true {
		return false, true
	}
	return nil, nil
}

//Перенос из n-1   -1  -1  -1   1   1   1
//1-е слагаемое	   -1  -1  -1   1   1   1
//2-е слагаемое	   -1 	0   1  -1   0   1
//Сумма   	    0   1  -1   1  -1   0
//Перенос в n+1	   -1  -1   0   0   1   1

// Полный сумматор двух тритов с переносом
func add_full_t(a trit, b trit, incarry trit) (c trit, outcarry trit) {
	s, sc := add_half_t(a, b)
	d, dc := add_half_t(s, incarry)
	ss, _ := add_half_t(sc, dc)
	return d, ss
}

// Таб.3 Троичное умножение
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
// Троичное сложение двух тритов с переносом
func mul_t(a trit, b trit) trit {

	if a == false && b == false {
		return true
	} else if a == false && b == true {
		return false
	} else if a == true && b == false {
		return false
	} else if a == true && b == true {
		return true
	}
	return nil
}

//
// TRYTE
//

// t5t4t3t2t1t0
type tryte [6]interface{}

//
// Операция сдвига тритов
//
// Параметр:
// if(d > 0) then "Вправо"
// if(d == 0) then "Нет сдвига"
// if(d < 0) then "Влево"
// Возврат: Троичное число
//
func shift_ts(x tryte, d int8) tryte {
	var tr tryte = x
	var n int8
	var s int8

	if d == 0 {
		return tr
	} else if d < 0 {
		n = -d
	} else {
		n = d
	}
	if d > 0 {
		for s = 0; s < n; s++ {
			for i := 0; i < len(tr)-1; i++ {
				tr[i] = tr[i+1]
			}
			id := len(tr) - 1
			tr[id] = nil
		}
	} else if d < 0 {
		for s = 0; s < n; s++ {
			for i := len(tr) - 1; i > 0; i-- {
				tr[i] = tr[i-1]
			}
			tr[0] = nil
		}
	}
	return tr
}

//
// Functions
//
// Изменить порядок тритов в трайте
func reverseTryte(input []interface{}) []interface{} {
	if len(input) == 0 {
		return input
	}
	return append(reverseTryte(input[1:]), input[0])
}

// Изменить порядок тритов в трайте
func printTryteInt(input []interface{}) []interface{} {
	if len(input) == 0 {
		return input
	}
	return append(printTryteInt(input[1:]), trit2int(input[0]))
}

// Изменить порядок тритов в трайте
func printTryteSymb(input []interface{}) []interface{} {
	if len(input) == 0 {
		return input
	}
	return append(printTryteSymb(input[1:]), trit2symb(input[0]))
}

// Возведение в степень по модулю 3
func pow3(x int8) int32 {
	var i int8
	var r int32 = 1
	for i = 0; i < x; i++ {
		r *= 3
	}
	return r
}

// ----------
// Main
// ----------
func main() {

	fmt.Printf("Run funcs ---------------------\n")

	fmt.Printf("- calculate trit  --------------\n")

	// Троичные переменные
	var a trit
	var b trit
	var c trit
	var carry trit

	a = nil
	b = true
	c = false
	carry = true

	fmt.Printf("a=%d, b=%d, c=%d, carry=%d   \n", trit2int(a), trit2int(b), trit2int(c), trit2int(carry))

	sf, sfc := add_full_t(a, b, carry)

	fmt.Printf("add_full_t( %d, %d, %d ) = %d,%d \n", trit2int(a), trit2int(b), trit2int(carry), trit2int(sf), trit2int(sfc))

	fmt.Printf("- calculate tryte  --------------\n")
	// Троичные переменные
	var x tryte
	var rx tryte
	var z tryte

	x[0] = true
	x[1] = false
	x[2] = nil
	x[3] = true
	x[4] = false
	x[5] = nil

	fmt.Printf("x = %v \n", printTryteInt(x[:]))
	fmt.Printf("x = %v \n", printTryteSymb(x[:]))
	rx = shift_ts(x, -5)
	fmt.Printf("shift_ts( x, %d ) = %v \n", -5, printTryteInt(rx[:]))
	fmt.Printf("shift_ts( x, %d ) = %v \n", -5, printTryteSymb(rx[:]))
	x = z
	x[5] = false
	fmt.Printf("x = %v \n", printTryteInt(x[:]))
	fmt.Printf("x = %v \n", printTryteSymb(x[:]))
	rx = shift_ts(x, 5)
	fmt.Printf("shift_ts( x, %d ) = %v \n", 5, printTryteInt(rx[:]))
	fmt.Printf("shift_ts( x, %d ) = %v \n", 5, printTryteSymb(rx[:]))

	fmt.Printf("--------------------------------\n")
}
