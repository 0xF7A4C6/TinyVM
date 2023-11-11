package main

func add(a int, b int) int {
    return a + b 
}

func main() {
    for {
        x := "tiny vm on top"

        a := 3
        b := 5

        if add(a, b) == add(b, b) {
            continue
        }

        break
    }
}