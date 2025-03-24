package arraysandslices

func SumAll(numberlist ...[]int) []int {
    var sumofall []int
    for _,numbers := range numberlist{
        sumofall = append(sumofall,Sum(numbers))
    }
    return sumofall
}