package main

import (
	"fmt"
	"testing"
)

var tt trit
var tr [6]trit
var s, carry trit

func TestSetTrit(t *testing.T) {

	fmt.Println("Test TRIT:")
	fmt.Printf("tt = NIL \n")
	tt = nil
	fmt.Printf("tt == nil : %v \n", tt == nil)
	fmt.Printf("tt == true : %v \n", tt == true)
	fmt.Printf("tt == false : %v \n", tt == false)

	fmt.Printf("tt = TRUE \n")
	tt = true
	fmt.Printf("tt == nil : %v \n", tt == nil)
	fmt.Printf("tt == true : %v \n", tt == true)
	fmt.Printf("tt == false : %v \n", tt == false)

	fmt.Printf("tt = FALSE \n")
	tt = false
	fmt.Printf("tt == nil : %v \n", tt == nil)
	fmt.Printf("tt == true : %v \n", tt == true)
	fmt.Printf("tt == false : %v \n", tt == false)
	fmt.Println("")

	t.Log("right")

	//if "342.7" == Add(a, b) {
	//  t.Log("right")
	//} else {
	//  t.Error("wrong")
	//}
}

func TestSetTryte(t *testing.T) {

	fmt.Println("Test TRYTE:")
	fmt.Printf("tryte = %v \n", tr)
	tr[0] = true
	tr[1] = false
	fmt.Printf("tryte = %v \n", tr)
	fmt.Println("")

	t.Log("right")
}

func TestTrit2Int(t *testing.T) {

	fmt.Println("Test TRIT to INT:")
	fmt.Printf("trit2int(true) = %v \n", trit2int(true))
	fmt.Printf("trit2int(nil) = %v \n", trit2int(nil))
	fmt.Printf("trit2int(false) = %v \n", trit2int(false))
	fmt.Println("")

	t.Log("right")
}

func TestInt2Trit(t *testing.T) {

	fmt.Println("Test INT to TRIT:")
	fmt.Printf("int2trit(1) = %v \n", int2trit(1))
	fmt.Printf("int2trit(0) = %v \n", int2trit(0))
	fmt.Printf("int2trit(-1) = %v \n", int2trit(-1))
	fmt.Println("")

	t.Log("right")
}

func TestAddHalfTrits(t *testing.T) {

	fmt.Println("Test ADD HALF TRITS  add_half_t(...):")
	s, carry = add_half_t(false, false)
	fmt.Printf("add_half_t( -1 + -1) => %v, %v \n", s, carry)
	s, carry = add_half_t(false, nil)
	fmt.Printf("add_half_t( -1 +  0) => %v, %v \n", s, carry)
	s, carry = add_half_t(false, true)
	fmt.Printf("add_half_t( -1 +  1) => %v, %v \n", s, carry)
	s, carry = add_half_t(nil, false)
	fmt.Printf("add_half_t(  0 + -1) => %v, %v \n", s, carry)
	s, carry = add_half_t(nil, nil)
	fmt.Printf("add_half_t(  0 +  0) => %v, %v \n", s, carry)
	s, carry = add_half_t(nil, true)
	fmt.Printf("add_half_t(  0 +  1) => %v, %v \n", s, carry)
	s, carry = add_half_t(true, false)
	fmt.Printf("add_half_t(  1 + -1) => %v, %v \n", s, carry)
	s, carry = add_half_t(true, nil)
	fmt.Printf("add_half_t(  1 +  0) => %v, %v \n", s, carry)
	s, carry = add_half_t(true, true)
	fmt.Printf("add_half_t(  1 +  1) => %v, %v \n", s, carry)
	fmt.Println("")

	t.Log("right")
}

func TestAddFullTrits(t *testing.T) {

	fmt.Println("Test ADD FULL TRITS  add_full_t(...):")
	s, carry = add_full_t(false, false, false)
	fmt.Printf("add_full_t( -1 + -1 + -1) => %v, %v \n", s, carry)
	s, carry = add_full_t(false, nil, false)
	fmt.Printf("add_full_t( -1 + 0 + -1) => %v, %v \n", s, carry)
	s, carry = add_full_t(false, true, false)
	fmt.Printf("add_full_t( -1 + 1 + -1) => %v, %v \n", s, carry)

	s, carry = add_full_t(false, true, true)
	fmt.Printf("add_full_t( -1 + 1 + 1) => %v, %v \n", s, carry)
	s, carry = add_full_t(nil, true, true)
	fmt.Printf("add_full_t(  0 + 1 +  1) => %v, %v \n", s, carry)
	s, carry = add_full_t(true, true, true)
	fmt.Printf("add_full_t(  1 + 1 +  1) => %v, %v \n", s, carry)

	fmt.Println("")

	t.Log("right")
}

func TestMulTrits(t *testing.T) {
	fmt.Println("Test MUL TRITS  mul_t(...):")
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
	fmt.Println("")

	t.Log("right")
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = mul_t(false, true)
	}
}
