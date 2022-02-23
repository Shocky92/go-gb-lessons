package main

import "unsafe"

// Вопрос 1:
// a: В каких случаях необходима была явная передача указателя в качестве входных параметров и возвращаемых результатов или в качестве приёмника в методах?

// Ответ:  при работе со структурами
type student struct {
	name string
	age  int
}

func birthday(s *student) {
	s.age++ // в этом случае возраст изменится и запишется в структуру
}

func notBirthday(s student) {
	s.age++ // в этом случае возраст не изменится
}

// b: В каких случаях мы фактически имеем дело с указателями при передаче параметров, хотя явно их не указываем?
// Ответ: При работе с мапами, она представлена в виде хэш-таблицы, где представлен массив указателей на бакеты

type hmap struct {
	count int   // количество занятых ячеек
	B     uint8 // 2^B - всего ячеек

	buckets unsafe.Pointer // массив бакетов
}

// Вопрос 2:
// Для арифметического умножения и разыменования указателей в Go используется один и тот же символ — оператор (*).
// Как вы думаете, как компилятор Go понимает, в каких случаях в выражении имеется в виду умножение, а в каких — разыменование указателя?

// Ответ: при арифметическом умножении используется два операнда, при разыменовании указателя знак * помещается перед ним
