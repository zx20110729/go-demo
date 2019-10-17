package official

import (
	"math"
	"fmt"
	"time"
	"strings"
	"io"
	"image"
)

/**
方法和接口
 */
type Vertex3 struct {
	X, Y float64
}

// 1、方法，Go 没有类。不过你可以为结构体类型定义方法。
//给Vertex3添加方法
func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Func3() {
	v := Vertex3{3, 4}
	fmt.Println(v.Abs())
}

// 2、方法即函数
func abs(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Func4() {
	v := Vertex3{3, 4}
	fmt.Println(abs(v))
}

// 3、方法，为非结构体类型声明方法

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Func5() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// 4、指针接受者，此处若是用值传递的话，只会对原始 Vertex3 值的副本进行操作
func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Func6() {
	v := Vertex3{4, 5}
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}

// 5、指针与函数

func scale(v Vertex3, f float64) Vertex3 {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

func Func7() {
	v := Vertex3{3, 4}
	v = scale(v, 10)
	fmt.Println(v)
	fmt.Println(abs(v))
}

// 6、方法与指针重定向
func scale2(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Func8() {
	v := Vertex3{3, 4}
	v.Scale(2)
	fmt.Println(v)

	scale2(&v, 10)
	fmt.Println(v)

	p := &Vertex3{3, 4}
	p.Scale(3)
	fmt.Println(p)

	scale2(p, 8)
	fmt.Println(p)
}

// 7、方法与指针重定向

func Func9() {
	v := Vertex3{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(abs(v))

	p := &Vertex3{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(abs(*p))
}

// 8、选择值或指针作为接收者,首先，方法能够修改其接收者指向的值；其次，这样可以避免在每次调用方法时复制该值。
func Func10() {
	v := &Vertex3{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

// 9、接口
type Tester interface {
	test() float64
}

func Interface() {
	var a Tester
	f := MyFloat(-math.Sqrt2)
	v := Vertex3{3, 4}
	//a = f // MyFloat实现了Tester接口
	//a = &v // *Vertex3实现了Tester接口
	//a = v
	fmt.Println(f, v)
	fmt.Println(a.test())
}

// 10、接口与隐式实现
type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(t.S)
	}
}
func Interface2() {
	var i I = &T{"Hello"}
	i.M()
}

// 11、接口值
type F float64

func (f F) M() {
	fmt.Println(f)
}
func Interface3() {
	var i I = &T{"Hello"}
	fmt.Printf("(%v, %T)\n", i, i)

	i = F(math.Pi)
	fmt.Printf("(%v, %T)\n", i, i)
}

// 12、底层值为nil的接口值

func Interface4() {
	var i I
	var t *T
	i = t
	fmt.Printf("(%v, %T)\n", i, i)
	i.M()

	i = &T{"HaHa"}
	fmt.Printf("(%v, %T)\n", i, i)
	i.M()
}

// 13、nil接口值,nil 接口值既不保存值也不保存具体类型。

func Interface5() {
	var i I
	fmt.Printf("(%v, %T)\n", i, i)
	i.M()
}

// 14、空接口

func Interface6() {
	var i interface{}
	fmt.Printf("(%v, %T)\n", i, i)

	i = 42
	fmt.Printf("(%v, %T)\n", i, i)
	i = "hello"
	fmt.Printf("(%v, %T)\n", i, i)
}

// 15、类型断言
func Interface7() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}

// 16、类型选择
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case Vertex3:
		fmt.Printf("Vertex3 is %v\n", v)
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}
func Interface8() {
	do(21)
	do("Hello")
	do(Vertex3{3, 4})
	do(true)
}

// 17、Stringer
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func PrintString() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 42}
	fmt.Println(a, z)
}

// 18、错误
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
func run() error {
	return &MyError{time.Now(), "It didn't work."}
}
func Error() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// 19、Reader
func Reader() {
	r := strings.NewReader("Hello Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// 20、图像

func Image() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
