package main

import (
	"fmt"
)

func main() {
	Float32()
	Float64()
}

func Float32() {
	fmt.Println("Float32")
	var y,yp1 float32
	y = 1
	y *= 2
	yp1 = y + 1
	for yp1 - y == 1 {
		y *= 2
		yp1 = y + 1
	}

	var x float32
	x = 1
	for y + x == y {
		x += 1
	}
	var base int
	base = int((y+x)-y)
	
	fmt.Printf("Base: %d\n",base)

	y = 1
	var p int
	p = 0
	var ep1 float32

	y *= float32(base)
	p++
	ep1 = y + 1

	for ep1 - y == 1 {
		y *= float32(base)
		p++
		ep1 = y + 1
	}

	fmt.Printf("P: %d\n",p)

	var eps float32
	eps = float32(base) / y;

	fmt.Printf("Eps by formula: %e\n",eps)

	var epsp1 float32
	x = 1
	eps = x
	x = x / 2
	epsp1 = x + 1
	for epsp1 > 1 {
		eps = x
		x = x / 2
		epsp1 = x + 1
	}

	fmt.Printf("Brute force eps: %e\n",eps)

}

func Float64() {
	fmt.Println("Float64")
	var y,yp1 float64
	y = 1
	y *= 2
	yp1 = y + 1
	for yp1 - y == 1 {
		y *= 2
		yp1 = y + 1
	}

	var x float64
	x = 1
	for y + x == y {
		x += 1
	}
	var base int
	base = int((y+x)-y)
	
	fmt.Printf("Base: %d\n",base)

	y = 1
	var p int
	p = 0
	var ep1 float64

	y *= float64(base)
	p++
	ep1 = y + 1

	for ep1 - y == 1 {
		y *= float64(base)
		p++
		ep1 = y + 1
	}

	fmt.Printf("P: %d\n",p)

	var eps float64
	eps = float64(base) / y;

	fmt.Printf("Eps by formula: %e\n",eps)

	var epsp1 float64
	x = 1
	eps = x
	x = x / 2
	epsp1 = x + 1
	for epsp1 > 1 {
		eps = x
		x = x / 2
		epsp1 = x + 1
	}

	fmt.Printf("Brute force eps: %e\n",eps)

}
