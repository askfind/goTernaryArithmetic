package main

import (
	"fmt"
)

// Таб.1 Алфавит троичной симметричной системы счисления
// +--------------------------+-------+-------+-------+
// | Числа                    |  -1   |   0   |   1   |
// +--------------------------+-------+-------+-------+
// | Символы (вар.1)          |   -   |	  0   |	  +   |
// +--------------------------+-------+-------+-------+
// | Символы (вар.2)          |   v   |	  0   |	  ^   |
// +--------------------------+-------+-------+-------+
// | Символы (вар.3)          |   N   |	  O   |	  P   |
// +--------------------------+-------+-------+-------+
// | Символы (вар.4)          |   N   |	  Z   |	  P   |
// +--------------------------+-------+-------+-------+
// | Символы (вар.5)          |   0   |	  i   |	  1   |
// +--------------------------+-------+-------+-------+
// | Логика                   | false |  nil  | true  |
// +--------------------------+-------+-------+-------+

// В троичной логике только 729 коммутативных двухоперандных
// функций из 19,683, т.е. таких что F(A,B) = F(B,A)

// ----------------------------------------------------
// TRIT Arithmetic
// ----------------------------------------------------

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

// Таб.3 Таблица полного сумматора тритов
// .-----------------------------------------.
// | Перенос из n-1   -1  -1  -1   1   1   1 |
// | 1-е слагаемое    -1  -1  -1   1   1   1 |
// | 2-е слагаемое    -1   0   1  -1   0   1 |
// | Сумма   	       0   1  -1   1  -1   0 |
// | Перенос в n+1    -1  -1   0   0   1   1 |
// .-----------------------------------------.
// Полный сумматор двух тритов с переносом
func add_full_t(a trit, b trit, incarry trit) (c trit, outcarry trit) {
	s, sc := add_half_t(a, b)
	d, dc := add_half_t(s, incarry)
	ss, _ := add_half_t(sc, dc)
	return d, ss
}

// Таб.4 Троичное умножение
//       MUL
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
// Троичное умножение двух тритов с переносом
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

// Таб.5 Троичное отрицание
//       NOT
// .-----------.
// |  -  |  +  |
// |-----------|
// |  0  |  0  |
// |-----------|
// |  +  |  -  |
// .-----------.
func not_t(a trit) trit {
	if a == false {
		return true
	} else if a == true {
		return false
	}
	return nil
}

// Таб.6 Троичное умножение
//       AND
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  -  |  -  |
// |-----------------------|
// |  0  |  -  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
//  X AND Y = MIN(X,Y)
// Троичное умножение двух тритов с переносом
func and_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return false
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return false
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return false
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

// Таб.7 Троичное или
//       OR
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  0  |  +  |
// |-----------------------|
// |  0  |  0  |  0  |  +  |
// |-----------------------|
// |  1  |  +  |  +  |  +  |
// .-----------------------.
//   X OR Y = MAX(X,Y)
func or_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return true
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return true
	} else if a == true && b == false {
		return true
	} else if a == true && b == nil {
		return true
	} else if a == true && b == true {
		return true
	}
	return nil
}

// Таб.8 Троичное исключающее или
//       XOR
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  0  |  +  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  +  |  0  |  -  |
// .-----------------------.
func xor_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return true
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return true
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return false
	}
	return nil
}

// Таб.9 Троичное
//       EQV
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
// EQV(X,Y) = NOT (XOR(X,Y))
func eqv_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

// Таб.10 Троичное
//       NAND
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  +  |  +  |
// |-----------------------|
// |  0  |  +  |  0  |  0  |
// |-----------------------|
// |  1  |  +  |  0  |  -  |
// .-----------------------.
// NAND(X,Y) = NOT (AND(X,Y))
func nand_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return true
	} else if a == false && b == true {
		return true
	} else if a == nil && b == false {
		return true
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return true
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return false
	}
	return nil
}

// Таб.11 Троичное
//       NOR
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  -  |
// |-----------------------|
// |  1  |  -  |  -  |  -  |
// .-----------------------.
// NOR(X,Y) = NOT((OR(X,Y))
func nor_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return false
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return false
	} else if a == true && b == true {
		return false
	}
	return nil
}

// Таб.12 Троичное
//       IMP
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  +  |  +  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
func imp_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return true
	} else if a == false && b == true {
		return true
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

// Таб.13 Троичное исключающее максимального
//        XMAX
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  0  |  +  |
// |-----------------------|
// |  0  |  0  |  -  |  +  |
// |-----------------------|
// |  1  |  +  |  +  |  -  |
// .-----------------------.
// XMAX:
//	F = MAX(A,B), если A != B
//	    0	    , если A == B
//	(имейте в виду - это не XOR).
func xmax_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return true
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return false
	} else if a == nil && b == true {
		return true
	} else if a == true && b == false {
		return true
	} else if a == true && b == nil {
		return true
	} else if a == true && b == true {
		return false
	}
	return nil
}

// Таб.14 Троичное Инверсно Исключающий минимального
//        IXMAX
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  -  |  -  |
// |-----------------------|
// |  0  |  -  |  +  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
func ixmax_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return false
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return false
	} else if a == nil && b == nil {
		return true
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

// Таб.15 Троичное Инверсно Исключающий минимального
//        MEAN
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  -  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  +  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  -  |
// .-----------------------.
// Mean:
//	Смотрит насколько "средние" операнды
//	Если ii, то возращает 1
//	Если iX или Xi, то возращает i
//	Иначе 0
func mean_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return true
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return false
	}
	return nil
}

// Таб.16 Троичное Инверсно Исключающий минимального
//        Magnitude
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  0  |  -  |  -  |
// |-----------------------|
// |  0  |  +  |  0  |  -  |
// |-----------------------|
// |  1  |  +  |  +  |  0  |
// .-----------------------.
// Magnitude:
//	Сравнение
//	(Функция нессимитричная)
//	Возращает 0, если A < B
//		  i, если A = B
//		  1, если A > B
func magnitude_t(a trit, b trit) trit {
	if a == false && b == false {
		return nil
	} else if a == false && b == nil {
		return false
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return true
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return false
	} else if a == true && b == false {
		return true
	} else if a == true && b == nil {
		return true
	} else if a == true && b == true {
		return nil
	}
	return nil
}

// Таб.17 Троичное дополнительный код
//        NEG
// .-----------.
// |  -  |  -  |
// |-----------|
// |  0  |  1  |
// |-----------|
// |  +  |  0  |
// .-----------.
func neg_t(a trit) trit {
	if a == false {
		return false
	} else if a == nil {
		return true
	}
	return nil
}

// Расмотрим еще некоторые функции
// { -, 0, + }
//
//  Сложение по модулю
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  -  0
//	0 | -  0  +
//	+ | 0  +  -
//	--+----------
func add_mod_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return false
	} else if a == false && b == true {
		return nil
	} else if a == nil && b == false {
		return false
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return true
	} else if a == true && b == false {
		return nil
	} else if a == true && b == nil {
		return true
	} else if a == true && b == true {
		return false
	}
	return nil
}

//  Перенос в сложении по модулю
//	--+----------
//	  | -  0  +
//	--+----------
//	- | -  0  0
//	0 | 0  0  0
//	+ | 0  0  +
//	--+----------
func carry_add_mod_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return nil
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return nil
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

//  Сложение с насышением
//	--+----------
//	  | -  0  +
//	--+----------
//	- | -  -  0
//	0 | -  0  +
//	+ | 0  +  +
//	--+----------
func add_satiation_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return false
	} else if a == false && b == true {
		return nil
	} else if a == nil && b == false {
		return false
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return true
	} else if a == true && b == false {
		return nil
	} else if a == true && b == nil {
		return true
	} else if a == true && b == true {
		return true
	}
	return nil
}

//
//  Функция Webb
//	--+----------
//	  | -  0  +
//	--+----------
//	- | 0  +  -
//	0 | +  +  -
//	+ | -  -  -
//	--+----------
func webb_t(a trit, b trit) trit {
	if a == false && b == false {
		return nil
	} else if a == false && b == nil {
		return true
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return true
	} else if a == nil && b == nil {
		return true
	} else if a == nil && b == true {
		return false
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return false
	} else if a == true && b == true {
		return false
	}
	return nil
}

//   Тождество (строгоe)
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  -  -
//	0 | -  +  -
//	+ | -  -  +
//	--+----------
func identity_strict_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return false
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return false
	} else if a == nil && b == nil {
		return true
	} else if a == nil && b == true {
		return false
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return false
	} else if a == true && b == true {
		return true
	}
	return nil
}

//   Тождество (weak)
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  0  -
//	0 | 0  0  0      в общем как умножение
//	+ | -  0  +
//	--+----------
func weak_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

//  Коньюнкция Лукашевича (сильная)
//	--+----------
//	  | -  0  +
//	--+----------
//	- | -  -  -
//	0 | -  -  0
//	+ | -  0  +
//	--+----------
func conjunction_lukashevich_strong_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return false
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return false
	} else if a == nil && b == nil {
		return false
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

//  Импликация Лукашевича
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  0  -
//	0 | +  +  0
//	+ | +  +  +
//	--+----------
func lukashevich_implication_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return true
	} else if a == nil && b == nil {
		return true
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return true
	} else if a == true && b == nil {
		return true
	} else if a == true && b == true {
		return true
	}
	return nil
}

//  Коньюнкция Клини
//	--+----------
//	  | -  0  +
//	--+----------
//	- | -  0  -
//	0 | 0  0  0
//	+ | -  0  +
//	--+----------
func klini_conjunction_t(a trit, b trit) trit {
	if a == false && b == false {
		return false
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

//  Импликация Клини
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  +  +
//	0 | 0  0  0
//	+ | -  0  +
//	--+----------
func implication_clinic_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return true
	} else if a == false && b == true {
		return true
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return false
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

//  Интуиционистская импликация Геделя
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  -  -
//	0 | +  +  0
//	+ | +  +  +
//	--+----------
func goedel_intuitionistic_implication_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return false
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return true
	} else if a == nil && b == nil {
		return true
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return true
	} else if a == true && b == nil {
		return true
	} else if a == true && b == true {
		return true
	}
	return nil
}

//  Материальная импликация
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  0  -
//	0 | +  0  0
//	+ | +  +  +
//	--+----------
func material_implication_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return true
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return true
	} else if a == true && b == nil {
		return true
	} else if a == true && b == true {
		return true
	}
	return nil
}

//  Функция следования Брусенцова
//	--+----------
//	  | -  0  +
//	--+----------
//	- | +  0  -
//	0 | 0  0  0
//	+ | 0  0  +
//	--+----------
func following_brusentsov_t(a trit, b trit) trit {
	if a == false && b == false {
		return true
	} else if a == false && b == nil {
		return nil
	} else if a == false && b == true {
		return false
	} else if a == nil && b == false {
		return nil
	} else if a == nil && b == nil {
		return nil
	} else if a == nil && b == true {
		return nil
	} else if a == true && b == false {
		return nil
	} else if a == true && b == nil {
		return nil
	} else if a == true && b == true {
		return true
	}
	return nil
}

// Троичное сложение двух тритов с переносом
// (Для измерения производительности операций int)
func sum_t(a int8, b int8, p0 int8) (int8, int8) {

	if a > 0 {
		a = 1
	}
	if a < 0 {
		a = -1
	}
	s := a + b + p0
	switch s {
	case -3:
		return 0, 1
		break
	case -2:
		return 1, -1
		break
	case -1:
		return -1, 0
		break
	case 0:
		return 0, 0
		break
	case 1:
		return 1, 0
		break
	case 2:
		return -1, 1
		break
	case 3:
		return 0, 1
		break
	default:
		return 0, 0
		break
	}
	return 0, 0
}

//
// Реализация чтение трита
//
func get_trit(t1 uint32, t0 uint32, p uint8) int {
	var trit int
	if (t0 & (1 << p)) == 0 {
		trit = 0
	} else {
		if (t1 & (1 << p)) == 0 {
			trit = -1
		} else {
			trit = 1
		}
	}
	return trit
}

// ----------------------------------------------------
// TRYTE
// ----------------------------------------------------

// t5t4t3t2t1t0
type tryte [6]interface{}

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

// ***********************************************
// *  Троичная арифметика компьютера "Сетунь-1958"
// *----------------------------------------------

/* Константы троичные */
const (
	// Длина слова
	SIZE_WORD_SHORT = 9
	SIZE_WORD_LONG  = 18

	// Описание ферритовой памяти FRAM
	NUMBER_ZONE_FRAM    = 3   // количество зон ферритовой памяти
	SIZE_ZONE_TRIT_FRAM = 54  // количнество коротких 9-тритных слов в зоне
	SIZE_ALL_TRIT_FRAM  = 162 // всего количество коротких 9-тритных слов

	SIZE_PAGES_FRAM     = 2  // количнество коротких 9-тритных слов в зоне
	SIZE_PAGE_TRIT_FRAM = 81 // количнество коротких 9-тритных слов в зоне

	// Адреса зон ферритовой памяти FRAM
	ZONE_M_FRAM_BEG = -120 /* ----0 */
	ZONE_M_FRAM_END = -41  /* -++++ */
	ZONE_0_FRAM_BEG = -40  /* 0---0 */
	ZONE_0_FRAM_END = 40   /* 0++++ */
	ZONE_P_FRAM_BEG = 42   /* +---0 */
	ZONE_P_FRAM_END = 121  /* +++++ */

	// Описание магнитного барабана DRUM
	SIZE_TRIT_DRUM      = 1944 // количество хранения коротких слов из 9-тритов */
	SIZE_ZONE_TRIT_DRUM = 54   // количество 9-тритных слов в зоне
	NUMBER_ZONE_DRUM    = 36   // количество зон на магнитном барабане
)

// Троичный тип данных
type trs struct {
	l  uint8  // длина троичного числа в тритах
	t1 uint32 // двоичное битовое поле троичного числа
	t0 uint32
	// if t0[i] == 0 then trit[i] = NIL emdif
	// if (t0[i] == 1) && (t1[i] == 0)  then trit[i] = FALSE emdif //
	// if (t0[i] == 1) && (t1[i] == 1)  then trit[i] = TRUE emdif //
}

// Основные регистры в порядке пульта управления
var (
	K trs // K(1:9)  код команды (адрес ячейки оперативной памяти)
	F trs // F(1:5)  индекс регистр
	C trs // C(1:5)  программный счетчик
	W trs // W(1:1)  знак троичного числа
	//
	ph1 trs // ph1(1:1) 1 разряд переполнения
	ph2 trs // ph2(1:1) 2 разряд переполнения
	S   trs // S(1:18) аккумулятор
	R   trs // R(1:18) регистр множителя
	MB  trs // MB(1:4) троичное число зоны магнитного барабана
	// Дополнительный
	MR trs // временный регистр для обмена троичным числом
)

// Троичное умножение двух тритов с переносом
func and_trit(a trit, b trit) int8 {
	if a == false && b == false {
		return -1
	} else if a == false && b == nil {
		return -1
	} else if a == false && b == true {
		return -1
	} else if a == nil && b == false {
		return -1
	} else if a == nil && b == nil {
		return 0
	} else if a == nil && b == true {
		return -1
	} else if a == true && b == false {
		return -1
	} else if a == true && b == nil {
		return 0
	} else if a == true && b == true {
		return 1
	}
	return 0
}

func or_trit(a trit, b trit) int8 {
	if a == false && b == false {
		return -1
	} else if a == false && b == nil {
		return 0
	} else if a == false && b == true {
		return 1
	} else if a == nil && b == false {
		return 0
	} else if a == nil && b == nil {
		return 0
	} else if a == nil && b == true {
		return 1
	} else if a == true && b == false {
		return 1
	} else if a == true && b == nil {
		return 1
	} else if a == true && b == true {
		return 1
	}
	return 0
}

func xor_trit(a trit, b trit) int8 {
	if a == false && b == false {
		return -1
	} else if a == false && b == nil {
		return 0
	} else if a == false && b == true {
		return 1
	} else if a == nil && b == false {
		return 0
	} else if a == nil && b == nil {
		return 0
	} else if a == nil && b == true {
		return 0
	} else if a == true && b == false {
		return 1
	} else if a == true && b == nil {
		return 0
	} else if a == true && b == true {
		return -1
	}
	return 0
}

// Очистить троичное число и длину
func clear_full_trs(tr *trs) {
	tr.l = 0
	tr.t0 ^= tr.t0
}

// Очистить троичное число
func clear_trs(tr *trs) {
	tr.t0 ^= tr.t0
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

// Преобразование трита в целое число
func trs2int(tr trs, p uint8) int8 {
	if p > 31 {
		p = 31
	}
	if (tr.t0 & (1 << p)) == 0 {
		return 0
	} else {
		if (tr.t1 & (1 << p)) == 0 {
			return -1
		} else {
			return 1
		}
	}
}

// Преобразование целое число в трит
func int2trs(tr trs, p uint8, i int8) trs {
	if p > 31 {
		p = 31
	}
	if i > 0 {
		tr.t1 |= (1 << p)
		tr.t0 |= (1 << p)
		return tr
	}
	if i < 0 {
		tr.t1 &^= (1 << p)
		tr.t0 |= (1 << p)
		return tr
	}
	tr.t1 &^= (1 << p)
	tr.t0 &^= (1 << p)
	return tr
}

//  Преобразовать трит в символ '-','0','+'
func trs2symb(tr trs, p uint8, i int8) string {
	if p > 31 {
		p = 31
	}
	if tr.t0 == 0 {
		return "0"
	} else {
		if tr.t1 == 0 {
			return "-"
		} else {
			return "+"
		}
	}
}

// Операция знак SGN троичного числа
func sgn_trs(x trs) int8 {
	var i int8
	if x.l > 31 {
		x.l = 31
	}
	for i = int8(x.l) - 1; i >= 0; i -= 1 {
		if ((x.t0 & (1 << i)) > 0) && ((x.t1 & (1 << i)) == 0) {
			return -1
		} else if ((x.t0 & (1 << i)) > 0) && ((x.t1 & (1 << i)) > 0) {
			return 1
		}
	}
	return 0
}

// Операция AND trs
func and_trs(x trs, y trs) trs {

	var r trs
	var i, j uint8
	var a, b, s int8

	x.l %= 32
	y.l %= 32

	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	for i = 0; i < j; i++ {
		a = trs2int(x, i)
		b = trs2int(y, i)
		s = and_trit(int2trit(a), int2trit(b))
		r = int2trs(r, i, s)
	}

	return r
}

// Операция OR trs
func or_trs(x trs, y trs) trs {

	var r trs
	var i, j uint8
	var a, b, s int8

	x.l %= 32
	y.l %= 32

	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	for i = 0; i < j; i++ {
		a = trs2int(x, i)
		b = trs2int(y, i)
		s = or_trit(int2trit(a), int2trit(b))
		r = int2trs(r, i, s)
	}

	return r
}

// Операция XOR trs
func xor_trs(x trs, y trs) trs {

	var r trs
	var i, j uint8
	var a, b, s int8

	x.l %= 32
	y.l %= 32

	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	for i = 0; i < j; i++ {
		a = trs2int(x, i)
		b = trs2int(y, i)
		s = xor_trit(int2trit(a), int2trit(b))
		r = int2trs(r, i, s)
	}

	return r
}

/**
 * Операция сдвига тритов
 * Параметр:
 * if(d > 0) then "Вправо"
 * if(d == 0) then "Нет сдвига"
 * if(d < 0) then "Влево"
 * Возврат: Троичное число
 */
func shift_trs(tr trs, d int8) trs {
	tr.l %= 32
	if d > 0 {
		tr.t1 >>= d
		tr.t0 >>= d
	} else if d < 0 {
		tr.t1 <<= -d
		tr.t0 <<= -d
	}
	return tr
}

// Троичное сложение троичных чисел
func add_trs(x trs, y trs) trs {
	var i, j uint8
	var a, b, s, p0, p1 int8
	var r trs

	if x.l > 31 {
		x.l = 31
	}
	if y.l > 31 {
		y.l = 31
	}

	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	r.l = j
	r.t0 = 0

	p0 = 0
	p1 = 0

	for i = 0; i < j; i++ {

		a = trs2int(x, i)
		b = trs2int(y, i)
		s, p1 = sum_t(a, b, p0)
		r = int2trs(r, i, s)
		p0 = p1

		x.t1 >>= 1
		x.t0 >>= 1
		y.t1 >>= 1
		y.t0 >>= 1
	}
	return r
}

// Троичное сложение троичных чисел
func sub_trs(x trs, y trs) trs {
	var i, j uint8
	var a, b, s, p0, p1 int8
	var r trs

	if x.l > 31 {
		x.l = 31
	}
	if y.l > 31 {
		y.l = 31
	}

	if x.l >= y.l {
		j = x.l
	} else {
		j = y.l
	}

	r.l = j
	r.t0 = 0

	p0 = 0
	p1 = 0

	for i = 0; i < j; i++ {

		a = trs2int(x, i)
		b = 0 - trs2int(y, i)
		s, p1 = sum_t(a, b, p0)
		r = int2trs(r, i, s)
		p0 = p1

		x.t1 >>= 1
		x.t0 >>= 1
		y.t1 >>= 1
		y.t0 >>= 1
	}
	return r
}

// Сложение  в  S (S)+(A*)=>(S)
//func add_tr(x lts ,y lts ) lts  {
//     return nil
//}

// Вычитание в  S (S)-(A*)=>(S)
//func add_tr(x interface ,y interface ) interface  {
//     return nil
//}

// Умножение 0 (S)=>(R); (A*)(R)=>(S)
//func mulz_tr(x trs,y trs) trs {
//     return nil
//}

// Умножение + (S)+(A*)(R)=>(S)
//func mulp_tr(x trs,y trs) trs {
//     return nil
//}

// Умножение - (A*)+(S)(R)=>(S)
//func muln_tr(x trs,y trs) trs {
//     return nil
//}

// Поразрядное умножение (A*)[x](S)=>(S)
//func and_tr(x trs,y trs) trs {
//     return nil
//}

// Посылка в R (A*)=>(R)
// Нормализация	Норм.(S)=>(A*); (N)=>(S)

// Аппаратный сброс.
// Очистить память и регистры
// виртуальной машины "Сетунь-1958"
func reset_setun_1958() {
	//
	// clean_fram();	/* Очистить  FRAM */
	// clean_drum();	/* Очистить  DRUM */
	//
	clear_full_trs(&K) /* K(1:9) */
	K.l = 9
	clear_full_trs(&F) /* F(1:5) */
	F.l = 5
	clear_full_trs(&C) /* K(1:5) */
	C.l = 5
	clear_full_trs(&W) /* W(1:1) */
	W.l = 1
	//
	clear_full_trs(&ph1) /* ph1(1:1) */
	ph1.l = 1
	clear_full_trs(&ph2) /* ph2(1:1) */
	ph2.l = 1
	clear_full_trs(&S) /* S(1:18) */
	S.l = 18
	clear_full_trs(&R) /* R(1:18) */
	R.l = 18
	clear_full_trs(&MB) /* MB(1:4) */
	MB.l = 4
	//
	clear_full_trs(&MR) /* Временный регистр данных MR(1:9) */
	MR.l = 9
}

// ---------------------------------------------------
// Main
// ---------------------------------------------------
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

	// test get_trit()
	var t1 uint32
	var t0 uint32
	var p uint8

	fmt.Printf("Test print trt:\n")
	// -1
	p = 0
	t1 = t1 &^ (1 << p)
	t0 |= (1 << p)
	trt := get_trit(t1, t0, p)
	fmt.Printf(" trt=% 2d\n", trt)
	// 0
	p = 0
	t0 = t0 &^ 1
	trt = get_trit(t1, t0, p)
	fmt.Printf(" trt=% 2d\n", trt)
	// 1
	p = 0
	t1 |= (1 << p)
	t0 |= (1 << p)
	trt = get_trit(t1, t0, p)
	fmt.Printf(" trt=% 2d\n", trt)

	// Сдвиг
	fmt.Printf("Shift to the left {t1,t0}<<= 2:\n")
	t1 <<= 2
	t0 <<= 2
	trt = get_trit(t1, t0, 2)
	fmt.Printf(" shift trt=% 2d\n", trt)

	fmt.Printf("--- Operation Setun-1958 ---\n")

	reset_setun_1958()
	K = int2trs(K, 0, -1)
	fmt.Println(K)
	fmt.Println(sgn_trs(K))
	K = int2trs(K, 0, 0)
	fmt.Println(K)

	K = int2trs(K, 0, 1)
	fmt.Printf(" K=%v\n", K)
	K = add_trs(K, K)
	K = sub_trs(K, K)
	fmt.Printf(" K=K+K=%v\n", K)

	K = int2trs(K, 0, 1)
	K = shift_trs(K, -1)
	K = shift_trs(K, -3)
	fmt.Println(K)
	K = sub_trs(K, K)
	fmt.Println(K)
	K = shift_trs(K, 2)
	fmt.Println(K)
	K = shift_trs(K, 100)
	fmt.Println(K)

	fmt.Println(S)

	fmt.Printf("--------------------------------\n")
}
