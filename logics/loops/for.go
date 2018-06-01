package loops

func IterativeFactorial (x int) int{
    var prod, i int = 1, 1;
    for i=1; i<=x; i++ {
        prod *= i;
    }
    return prod;
}

func RecursiveFactorial(x int) int{
    if x <= 1 {
        return 1
    }
    return x * RecursiveFactorial(x-1);
}