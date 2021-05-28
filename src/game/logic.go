package game

func NewSolvedTable() Table{
    // Solve a newely created table and mask a random num of values
    // Return a Table object
    tab := newTable()
    
    solve(&tab, 1, 1, true)
    tab.Mask()
    return tab
}

func newTable() Table{
    // create a Table full of 0 values
    var tab Table
    for i:=0; i < 9 ; i++{
        for j:=0; j < 9; j++{
            tab.Values[i][j].Mask()
        }
    }
    return tab
}

func solve(t *Table, row, col int, rnd bool) error{
    // First row and column are solved by randomizing them
    // Rest of the table solved by using a backtracking approach
    // Values are initialized using Table.Init(row, col, val) method
    // Invalid values are overwriteen with Field.Mask(val)
    // Reinitialize the field with a new value that meets the requirements
    // Return nil when row value reaches 9 otherwise return an error
    
    if rnd == true{
        t.Randomize()
    }
    if col == 9{
        col = 1
        row = row+1
    }
    if row == 9{
        return nil
    }
    var err error
    t.Values[row][col].Mask()
    for i:=1; i < 10; i++{
        err = t.Init(row, col, i)
        if err == nil{
            err = solve(t, row, col+1, false)
        }
        if err!=nil{
            t.Values[row][col].Mask()
        } else {
            // Table solved break the loop
            break
        }
    }
    
    return err
}
