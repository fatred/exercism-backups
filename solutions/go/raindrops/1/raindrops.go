package raindrops

import "fmt"

func Convert(number int) string {
    var result string
    // 3 == "Pling"
    if number % 3 == 0 {
        result += "Pling"
    }
    // 5 == "Plang"
	if number % 5 == 0{
        result += "Plang"
    }
    // 7 == "Plong"
    if number % 7 == 0 {
        result += "Plong"
    }
    if result == "" {
        result += fmt.Sprintf("%v", number)
    }
    return result
}
