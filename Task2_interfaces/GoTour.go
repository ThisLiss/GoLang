package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
)

type IPAddr [4]byte
type ErrNegativeSqrt float64
type MyReader struct{}

type rot13Reader struct {
	r io.Reader
}

type Image struct{
	w, h int
}

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0],i[1],i[2],i[3])
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0,ErrNegativeSqrt(x)
	}
	z := 1.0
	old := 0.0
	for i:=0; i <= 100; i++	{
		old = z
		z = z - (z*z - x)/(2*z)
		if Equals(old,z) {
			break
		}
	}
	return z, nil
}

func Equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	if delta < 1e-10 {
		return true
	}
	return false
}

func (m MyReader) Read(r []byte) (int, error) {

	for i := range r {
		r[i] = 'A'
	}

	return len(r), nil
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	val, err := rot13.r.Read(b)
	for i := range b {
		if b[i] >= 97 && b[i] <= 109 ||  b[i] >= 65 && b[i] <= 77{
			b[i]+=13
		} else if b[i] >= 110 && b[i] <= 122 || b[i] >= 78 && b[i] <= 90{
			b[i]-=13
		}
	}
	return val, err
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.RGBA {
	return color.RGBA{R: uint8(x*y), G: uint8(x*y), B: 255, A: 255}
}


func main() {

	//Упражнение: Stringers
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}


	//Упражнение: ошибки
	sqrt, err := Sqrt(-2)
	//sqrt, err := Sqrt(4)
	if err == nil {
		fmt.Printf("Sqrt = %v\n", sqrt)
	} else {
		fmt.Println(err)
	}


	//Упражнение: Reader
	reader.Validate(MyReader{})

	//Упражнение: rot13Reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	//Упражнение: изображения
	myImage := Image{100,100}
	fmt.Println(myImage.ColorModel())
	fmt.Println(myImage.Bounds())
	fmt.Println(myImage.At(30, 30))
}
