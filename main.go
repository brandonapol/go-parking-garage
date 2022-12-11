package main

import (
    "fmt"
    "bufio"
    "os"
    // "time"
)

type Garage struct {
    tickets []int
    spaces int 
    current_ticket map[int]bool
    flag bool
}

func makeGarage() *Garage{
    g := Garage {
        tickets: []int{1,2,3,4,5,6,7,8,9,10},
        spaces: 10,
        current_ticket: map[int]bool{},
        flag: true,
    }
    return &g
}

func runGarage(g *Garage) {

    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Println("How can we help? \n - Get Ticket (\"Get\") \n - Checkout (\"Checkout\") \n - Show Current Ticket(s) (\"Show\") \n - Help (\"Help\")\n - Quit (\"Quit\")")
        input, _ := reader.ReadString('\n')
        fmt.Printf("You have chosen %s", input)
    }
    

}

func takeTicket() {

}

func showSpaces(g *Garage) {
    fmt.Printf("Spaces available: %d", g.spaces)
}

func main() {
	fmt.Println("garage")

    g := makeGarage()
    runGarage(g)
}