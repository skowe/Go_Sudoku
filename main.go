package main

import (
    "fmt"
    "github.com/skowe/Go_Sudoku/game"
)

const Version = 0.0.1

func main(){
    tab := game.NewSolvedTable()
    play := true
    
    for play{
        fmt.Println(tab)
        row, col, val, err := game.ParseIn()
        if err != nil{
            fmt.Println(err)
            continue
        }
        err = tab.Insert(row, col, val)
        
        if err != nil{
            fmt.Println(err)
            continue
        }
        
        if tab.Solved(){
            play = false
        }
    }
}
