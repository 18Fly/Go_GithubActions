package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"time"

	"golang.org/x/tour/pic"
)

type Vertes struct {
	X, Y int
}

func (v Vertes) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v *Vertes) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)

	} else {
		return float64(f)

	}
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("二倍的 %v 是 %v\n", v, v*2)
	case string:
		fmt.Printf("%s 的长度是 %v\n", v, len(v))
	default:
		fmt.Printf("我不知道 %v 类型 %T!\n", v, v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		return x, nil
	}
}

type Image struct {
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	// v := Vertes{3, 4}
	// fmt.Println(v.Abs()) // 方法即函数
	// f := MyFloat(-math.Sqrt2)
	// fmt.Println(f.Abs())
	// v.Scale(2)
	// fmt.Println(v)

	// var i interface{} = "hello"
	// v, ok := i.(int) // 如果断言触发，会讲v置位断言要求类型的零值
	// fmt.Printf("%T\t%v", v, ok)

	// do(21)
	// do("hello")
	// do(true)

	// a := Person{"Alice", 18}
	// b := Person{"Bob", 20}
	// fmt.Println(a, b)

	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	// for name, ip := range hosts {
	// 	fmt.Printf("%v: %v\n", name, ip)
	// }

	// if err := run(); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(Sqrt(2))
	// fmt.Println(Sqrt(-2))
	// fmt.Println('A')

	pic.ShowImage(Image{})
}
