package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// Embedded
type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// Goのインスタンス生成の書き方
// package.Newが基本
// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

// Alias
type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func do(i interface{}) {
	// switch type文
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

type PersonMore struct {
	Name string
	Age  int
}

func (p PersonMore) String() string {
	return fmt.Sprintf("My name is %v.", p.Name)
}

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func (v Vertex) Plus() int {
	return v.x + v.y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!\n", v.x, v.y)
}

func main() {
	v39 := Vertex{3, 4}
	v39.Scale(10)
	fmt.Println(v39.Area())

	// Embedded
	v40 := New(3, 4, 5)
	v40.Scale3D(10)
	fmt.Println(v40.Area3D())

	// non-struct
	myInt := MyInt(10)
	fmt.Println(myInt.Double())

	// interface & duck typing
	// Humanでinterfaceを指定する
	var mike43 Human = &Person{"Mike"}
	var x43 Human = &Person{"X"}
	DriveCar(mike43)
	DriveCar(x43)

	// type assertion
	// interface{}って書くとどの値も利用できる
	// i.(int) これがtype assertion
	do(10)
	do("Mike")
	do(true)

	// これはinterfaceからではなくintからfloatに変換するもの
	// → type conversion。type assertionとは違うもの
	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)

	// Stringer
	// toStringの内容を変更したい場合に使う
	mike45 := PersonMore{"Mike", 22}
	fmt.Println(mike45)

	// Errorの判別をしたいときに単純に値渡しで比較すると
	// 全部同じとみなされる
	// UserNotFound{"mike1"} = UserNotFound{"mike2"}
	// なので&つけて、ポインタとして利用する
	if err46 := myFunc(); err46 != nil {
		fmt.Println(err46)
	}

	v47 := Vertex{3, 4}
	fmt.Println(v47.Plus())
	v48 := Vertex{3, 4}
	fmt.Println(v48)
}
