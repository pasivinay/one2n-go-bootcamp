package structs

type Rectangle struct {
    Width, Height float64
}

type Circle struct {
    Radius float64
}

type Triangle struct {
    Base, Height float64
}

type Shape interface {
	Area() float64
}