package recu

func Factorial(x int) int{
    if x<=1{
        return 1
    }
    return x * Factorial(x-1)
}