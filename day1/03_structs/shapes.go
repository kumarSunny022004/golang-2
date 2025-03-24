package structs


type Rectangle struct{
	width float64
	height float64
}

type Circle struct{
	radius float64
}

type Shape interface {
	Area() float64
}

 func Perimeter(r Rectangle) float64{
	 return 2 * (r.width + r.height)
 }

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func (c Circle) Area() float64{
	return 3.141592653589793 * c.radius * c.radius
}