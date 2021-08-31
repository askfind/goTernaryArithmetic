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

// Объявление троичных типов
type trit interface{}
type tryte [6]interface{}

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
//2-е слагаемое	   -1 	0	1  -1   0   1
//Сумма   	        0   1  -1   1  -1   0
//Перенос в n+1	   -1  -1	0   0   1   1

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

// ----------
// Main
// ----------
func main() {

	fmt.Printf("Run funcs ---------------------\n")

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

	fmt.Printf("--------------------------------\n")
}
