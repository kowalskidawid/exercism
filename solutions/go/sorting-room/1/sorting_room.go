package sorting

import (
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
    return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
    return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
    fn, ok := fnb.(FancyNumber)
    if !ok {
        return 0
    }

    n, err := strconv.Atoi(fn.Value())
    if err != nil {
        return 0
    }

    return n
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    return fmt.Sprintf(
        "This is a fancy box containing the number %.1f",
        float64(ExtractFancyNumber(fnb)),
    )
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
    switch t := i.(type) {
        case float64:
            return DescribeNumber(t)
        case int:
            return DescribeNumber(float64(t))
        case NumberBox:
            return DescribeNumberBox(t)
        case FancyNumber:
        	return DescribeFancyNumberBox(t)  
        case FancyNumberBox:
        	return DescribeFancyNumberBox(t)
    }

    return "Return to sender"
}
