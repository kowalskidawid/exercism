package robotname

import "math/rand"
import "fmt"

func random(maxRange int) int {
    return rand.Intn(maxRange - 1) + 1
}

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
    if r.name != "" {
        return r.name, nil
    }
    letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    r.name = fmt.Sprintf("%c%c%d%d%d", letters[random(26)], letters[random(26)], random(10), random(10), random(10))
	return r.name, nil
}

func (r *Robot) Reset() {
    r.name = ""
}
