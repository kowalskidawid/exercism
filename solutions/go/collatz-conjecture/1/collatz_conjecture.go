package collatzconjecture

import (
    "errors"
)

func CollatzConjecture(n int) (int, error) {
    if n == 0 {
        return 0, errors.New("zero is an error")
    } else if n < 0 {
        return 0, errors.New("negative value is an error")
    }
    steps := 0
    for n > 1 {
        steps++
    	isEven := n % 2 == 0
        if isEven {
            n = n / 2
        } else {
            n = (n * 3) + 1
        }
    }
    return steps, nil
}
