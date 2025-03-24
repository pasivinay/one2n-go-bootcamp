package arraysandslices

func Sum(List []int) int {
    sum := 0
    for _, number := range List {
        sum += number
    }
    return sum
}