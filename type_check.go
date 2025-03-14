package main

// type A struct{}
// type B struct{}

// func main() {
// 	var a A
// 	var b B = a // 에러: cannot use a (type A) as type B
// }
ㅋ
type A interface { Foo() }
type B interface { Foo() }

func GetB() B { return nil }
func main() {
	var a A = GetB() // 정상: 메서드 집합이 동일
}