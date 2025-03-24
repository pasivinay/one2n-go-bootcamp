package structs

import(
    "testing"
)
func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{2.4,3.6}
    got := Perimeter(rectangle)
    want := 12.0

    if got != want {
        t.Errorf("%v is wrong perimeter, it should be %v", got, want)
    }
}

func TestArea(t *testing.T) {

    checkArea := func(t testing.TB, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()

        if got != want {
            t.Errorf("%v is wrong area, it should be %v", got, want)
        }
    }

    // t.Run("Get Area of Rectangle", func(t *testing.T) {
    //     rectangle := Rectangle{2.4,3.6}
    //     checkArea(t, rectangle, 8.64)
    // })
    // t.Run("Get Area of a circle", func(t *testing.T) {
    //     circle := Circle{10.0}
    //     checkArea(t, circle, 314.1592653589793)
    // })

    areaTests := []struct {
        shape Shape
        want float64
    }{
        {Rectangle{2.4,3.6},8.64},
        {Circle{10.0},314.1592653589793},
        {Triangle{2.4,3.6},4.32},
    }

    for _,tt := range areaTests {
        checkArea(t, tt.shape, tt.want)
    }
}