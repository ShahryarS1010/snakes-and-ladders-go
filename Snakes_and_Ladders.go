package main
import "fmt"
// Import to roll the dice
import "math/rand"
import "time"

func main() {
  
	
	fmt.Printf("Welcome to Snakes and Ladders\n" + "\n")
	
	// Start the Game
	Game()
	
	fmt.Printf("Thanks for Playing")
	
}

func Game() {
  
    // Get the two players names
    var players [2]string
    
    fmt.Printf("Enter Player 1's Name: ")
    fmt.Scanf("%s", &players[0])
    
    fmt.Printf("Enter Player 2's Name: ")
    fmt.Scanf("%s", &players[1])
    
    // Each player starts on Square 0
    var player1 int = 0
    var player2 int = 0
    
    // Player 1 starts first
    turn := 0
    
    for { 
        
        for {
            // Take pauses before each round
            fmt.Print("\n" + "Press enter to continue" + "\n")
            fmt.Scanf("%s")
            break
        }
        
        fmt.Print("\n")
        
        // Roll the dice
        randomdice := rand.NewSource(time.Now().UnixNano()) 
        diceroll := rand.New(randomdice) 
        dice := diceroll.Intn(6)
       
        // Dice rolls are from 1 to 6
        dice += 1
        
        // Display the current squares of each player and whose turn it is
        fmt.Print(players[0] + "'s current square: ")
        fmt.Print(player1)
        fmt.Print("\n")
        fmt.Print(players[1] + "'s current square: ")
        fmt.Print(player2)
        fmt.Print("\n")
        
        fmt.Print(players[turn] + "'s turn " + "\n")
        
        // Player 1's Turn
        if turn == 0 {
            
            // Player moves to square based on dice roll
            player1 += dice
            
            // 100 is the maximum square
            if player1 > 100 {
                player1 = 100
            }
            
            // Display which square the player went to
            fmt.Print(players[0] + " rolls a ")
            fmt.Print(dice)
            fmt.Print(" to square ")
            fmt.Print(player1)
            fmt.Print("\n")
        
            // Player wins if Square 100 is reached
            if player1 == 100 {
            
                fmt.Print(players[turn] + " reached Square 100 and wins" + "\n")
                
                // End the game
                break
            
        }   
        
            // Numbers divisible by 4 are ladders
            if player1 % 4 == 0 {
            
                player1 += 6
                fmt.Print("Ladder, go up 6 squares to square ")
                fmt.Print(player1)
                fmt.Print("\n")
                
                // Switch turns 
                turn = 1
                continue
        }
        
            // Numbers divisible by 5 are snakes
            if player1 % 5 == 0 {
            
                player1 -= 6
                
                // 0 is the minimum square
                if player1 < 0 {
                player1 = 0
                }
                
                fmt.Print("Snake, go down 6 squares to square ")
                fmt.Print(player1)
                turn = 1
                continue
        }
          
            turn = 1
            continue
        }
        
        // Same as Player 1 but for Player 2
        if turn == 1 {
            
            player2 += dice
            
            if player2 > 100 {
                player2 = 100
            }
            
            fmt.Print(players[1] + " rolls a ")
            fmt.Print(dice)
            fmt.Print(" to square ")
            fmt.Print(player2)
            fmt.Print("\n")
        
            if player2 == 100 {
            
                fmt.Print(players[turn] + " reached Square 100 and wins" + "\n")
                break
            
        }   
        
            if player2 % 4 == 0 {
            
                player2 += 6
                fmt.Print("Ladder, go up 6 squares to square ")
                fmt.Print(player2)
                fmt.Print("\n")
                turn = 0
                continue
        }
        
            if player2 % 5 == 0 {
            
                player2 -= 6
                
                if player2 < 0 {
                player2 = 0
                }
                
                fmt.Print("Snake, go down 6 squares to square ")
                fmt.Print(player2)
                turn = 0
                continue
        }
          
            turn = 0
            continue
        }
        
        
        
    } 
    
 
    // Game is completed
    return
}


