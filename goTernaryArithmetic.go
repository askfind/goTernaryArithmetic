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

// Таб.2 Троичное сложение
// .------------------------.
// |     | -1  |  0  |  1   |
// |------------------------|
// | -1  | -+  |  -  |  0   |
// |------------------------|
// |  0  | -1  |  0  |  1   |
// |------------------------|
// |  1  |  0  |  1  |  +-  |
// .------------------------.

// Троичное сложение двух тритов с переносом
func add_t(a trit, b trit) (c trit, carry trit) {

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

	// Троичные переменные
	var t trit
	var tr tryte
	var s, carry trit

	// Тесты
	fmt.Println("\nt1: trit = NIL ")
	t = nil
	fmt.Printf("t == nil : %v \n", t == nil)
	fmt.Printf("t == true : %v \n", t == true)
	fmt.Printf("t == false : %v \n", t == false)

	fmt.Println("\nt2: t = TRUE ")
	t = true
	fmt.Printf("t == nil : %v \n", t == nil)
	fmt.Printf("t == true : %v \n", t == true)
	fmt.Printf("t == false : %v \n", t == false)

	fmt.Println("\nt3: t = FALSE ")
	t = false
	fmt.Printf("t == nil : %v \n", t == nil)
	fmt.Printf("t == true : %v \n", t == true)
	fmt.Printf("t == false : %v \n", t == false)

	fmt.Println("\nt4: tryte ")
	fmt.Printf("tryte = %v \n", tr)

	fmt.Println("\nt5: tryte]")
	tr[0] = true
	tr[1] = false
	fmt.Printf("tryte = %v \n", tr)

	fmt.Println("\nt6: trit2int()")
	fmt.Printf("trit2int(true) = %v \n", trit2int(true))
	fmt.Printf("trit2int(nil) = %v \n", trit2int(nil))
	fmt.Printf("trit2int(false) = %v \n", trit2int(false))

	fmt.Println("\nt7: int2trit(...)")
	fmt.Printf("int2trit(1) = %v \n", int2trit(1))
	fmt.Printf("int2trit(0) = %v \n", int2trit(0))
	fmt.Printf("int2trit(-1) = %v \n", int2trit(-1))

	fmt.Println("\nt8: add_t(...) --------")
	s, carry = add_t(false, false)
	fmt.Printf("add_t( -1 + -1) => %v, %v \n", s, carry)
	s, carry = add_t(false, nil)
	fmt.Printf("add_t( -1 +  0) => %v, %v \n", s, carry)
	s, carry = add_t(false, true)
	fmt.Printf("add_t( -1 +  1) => %v, %v \n", s, carry)
	s, carry = add_t(nil, false)
	fmt.Printf("add_t(  0 + -1) => %v, %v \n", s, carry)
	s, carry = add_t(nil, nil)
	fmt.Printf("add_t(  0 +  0) => %v, %v \n", s, carry)
	s, carry = add_t(nil, true)
	fmt.Printf("add_t(  0 +  1) => %v, %v \n", s, carry)
	s, carry = add_t(true, false)
	fmt.Printf("add_t(  1 + -1) => %v, %v \n", s, carry)
	s, carry = add_t(true, nil)
	fmt.Printf("add_t(  1 +  0) => %v, %v \n", s, carry)
	s, carry = add_t(true, true)
	fmt.Printf("add_t(  1 +  1) => %v, %v \n", s, carry)

	fmt.Println("\nt9: mul_t(...) --------")
	s = mul_t(false, false)
	fmt.Printf("mul_t( -1 * -1) => %v \n", s)
	s = mul_t(false, nil)
	fmt.Printf("mul_t( -1 *  0) => %v \n", s)
	s = mul_t(false, true)
	fmt.Printf("mul_t( -1 *  1) => %v \n", s)
	s = mul_t(nil, false)
	fmt.Printf("mul_t(  0 * -1) => %v \n", s)
	s = mul_t(nil, nil)
	fmt.Printf("mul_t(  0 *  0) => %v \n", s)
	s = mul_t(nil, true)
	fmt.Printf("mul_t(  0 *  1) => %v \n", s)
	s = mul_t(true, false)
	fmt.Printf("mul_t(  1 * -1) => %v \n", s)
	s = mul_t(true, nil)
	fmt.Printf("mul_t(  1 *  0) => %v \n", s)
	s = mul_t(true, true)
	fmt.Printf("mul_t(  1 *  1) => %v \n", s)

	fmt.Printf("---------------------\n")
}
