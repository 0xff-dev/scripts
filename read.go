package main


import (
    "fmt"
     "bufio"
)


func main() {
    var (
        firstName, lastName string
        age int
        format = "%s %s %d"
    )

    fmt.Println("Enter your full name: ")
    fmt.Scanln(&firstName, &lastName)
    fmt.Println(firstName, lastName)

    fmt.Scanf(format, &firstName, &lastName, &age)
    fmt.Printf("Hello %s %s you are %d\n", firstName, lastName, age)

    reader := bufio.NewReader()
    fmt.Println("Input something\n")
    input, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    fmt.Println("input is: ", input)

    /*
        sanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            data := strings.ToUpper(scanner.Text())
            fmt.Println(data)
        }
        if err := scanner.Err(); err != nil {
            os.Exit(1)
        }
    */
}
