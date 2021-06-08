package game

import (
    "errors"
    "fmt"
)


func ParseIn() (int, int, int, error){
    // Parses the user input and checks it against constraints 
    // return row index, column index, the value to be written, and nil for error
    // Return -1 for all 3 values if invalid input + a non nil error value
    fmt.Print("Enter 3 values from 1 to 9 for row, column and value in that order: ")
    var row, col, val int
    
    fmt.Scanf("%d%d%d", &row, &col, &val)
    
    if row > 9 || col > 9 || val > 9{
        return -1,-1,-1, errors.New("A value is out of acceptable range")
    }
    
    if row < 0 || col < 0 || val < 0{
        return -1,-1,-1, errors.New("A value is out of acceptable range")
    }
    
    return row-1, col-1, val, nil
}

func InSlice(val int, slc []int) error{
    // Helper function to check if a value is in slice
    // Also helps to keep track of exhausted values
    // Return an error slice contains the value othervise append to slice and return nil.
    for _,x := range slc{
        if x == val {
            return errors.New("Value tried.")
        }
    }
    slc = append(slc, val)
    return nil
} 

func FindCube(row, col int) int{
    // Helper function to find the number of a cube the value is located in
    return 3*(row/3) + col/3
}
