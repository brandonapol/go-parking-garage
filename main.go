package main


import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strcov"
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
        fmt.Println("")
        fmt.Println("How can we help? \n - Get Ticket (\"Get\") \n - Checkout (\"Checkout\") \n - Show Current Ticket(s) (\"Show\") \n - Help (\"Help\")\n - Quit (\"Quit\")")
        fmt.Println("")
        input, _ := reader.ReadString('\n')
        input = strings.ToLower(input)
        // the underscore is for a returned error value from ReadString, but we don't need it so the _ will discard it
        if input == "get\n" {
            takeTicket(g)
        } else if input == "checkout\n" {
            checkout(g)
        } else if input == "show\n" {
            showSpaces(g)
        } else if input == "quit\n" {
            break
        } else if input == "help\n" {
            fmt.Println(`Please type one of the following prompts: 

            For recieving a ticket and gaining access to the garage, 
            please type "Get".
            For checking out and escaping the parking 
            garage, please type "Checkout".
            If you'd like to quit the process and not get a ticket, 
            please type "Checkout". 
            If you'd like to show your current ticket(s) please type "Show".
            If you'd like to quit the process and not get a ticket, 
            please type "Quit".`)
        } else {
            fmt.Println("Please try a valid command.")
        }
    }
    

}

func takeTicket(g *Garage) {
    fmt.Println("taking ticket")
    g.spaces -= 1
    fmt.Println(g.spaces)
    ticketNum := g.tickets[0]
    g.tickets = g.tickets[1:len(g.tickets)]
    g.current_ticket[ticketNum] = false
}

func showSpaces(g *Garage) {
    fmt.Printf("Spaces available: %d\n", g.spaces)
    fmt.Printf("Current tickets: %v\n", g.current_ticket)
}

func checkout(g *Garage) {
    if g.spaces >= 10 {
        fmt.Println("The garage is empty")
        return
    } else {
        fmt.Println("Would you like to leave the garage? y/n")
        reader := bufio.NewReader(os.Stdin)
        leave, _ := reader.ReadString('\n')
        if leave == "y" {
            fmt.Println("Please give your ticket number: ")
            reader := bufio.NewReader(os.Stdin)
            ticketNum, _ := reader.ReadString('\n')
            ticketNum = strings.TrimSuffix(ticketNum, "\n")
            ticketNum, err := strcov.Atoi(ticketNum)
            if err != nil {
                fmt.Println("Invalid ticket number:", err)
            } 
            // else {
        
            //     for _, num := range g.tickets {
            //         if num == ticketNum {
        
            //             fmt.Println("Ticket is not available")
            //             return
            //         }
            //     fmt.Println("Please submit payment.")
            //     fmt.Println("Thank you for parking with us! \nHave a great day!")
            //     }
            // }
        }
    }
}

func main() {
	fmt.Println("garage")

    g := makeGarage()
    runGarage(g)
}