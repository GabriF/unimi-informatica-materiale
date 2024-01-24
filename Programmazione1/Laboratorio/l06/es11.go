package main

import . "fmt"
import "unicode"
import "strconv"
import "math"

func BaseXToBase10(n string, b uint) (uint, bool) {
    if b < 2 || b > 16 {
        return 0, false
    }

    var r uint

    for i, c := range n {
        var valore int

        peso := len(n) - i - 1

        if unicode.IsDigit(c) {
            valore, _ = strconv.Atoi(string(c))
        } else {
            switch c {
                case 'A':
                    valore = 10
                case 'B':
                    valore = 11
                case 'C':
                    valore = 12
                case 'D':
                    valore = 13
                case 'E':
                    valore = 14
                case 'F':
                    valore = 15
                default:
                    return 0, false
            }
        }



        if valore > int(b - 1) {
            return 0, false
        }

        r += uint(int(math.Pow(float64(b), float64(peso))) * valore)
    }

    return r, true
}

func main() {
    Println(BaseXToBase10("2FA", 17))
}
