package differenceofsquares

func SquareOfSum(n int) int {
    var sum int
    for i := 0; i <= n; i++ {
        sum += i
    }
    return sum * sum
}

func SumOfSquares(n int) int {
    var sum int
    for i := 0; i <= n; i++ {
        sum += i * i
    }
    return sum
}

func Difference(n int) int {
    squareOfSum := SquareOfSum(n)
    sumOfSquares := SumOfSquares(n)
    if squareOfSum > sumOfSquares {
        return squareOfSum - sumOfSquares
    }
    return sumOfSquares - squareOfSum
}
