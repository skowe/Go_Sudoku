package game

import ( 
    "errors"
    "fmt"
    "math/rand"
    "time"
)

type Field struct{
    value int
    mut bool
}
func (f *Field) Init(val int){
    f.value = val
    f.mut = false
}

func (f *Field) Mask(){
    f.value = 0
    f.mut = true
}

func (f *Field) FillIn(val int) error{
    if !f.mut{
        return errors.New("This field is imutable")
    }
    f.value = val
    return nil
}

func (f Field) IsZero() bool{
    return f.value == 0
}
//========================================================================================

type Table struct{
    Values [9][9]Field
}

func (t *Table) Solved() bool{
    //Checks if the table is solved
    //Table is solved once it contains no 0s
    //Return true if it is otherwise false
    for _,row := range t.Values{
        for _, fld := range row{
            if fld.IsZero(){
                return false
            }
        }
    }
    
    return true
}

func (t *Table) ValInRow(row, val int) error{
    //Checks if row contains a value
    //Return error if it does nil otherwise
    
    for i := 0; i < 9; i++ {
        if t.Values[row][i].value == val{
            return errors.New("Value in row")
        }
    }
    return nil
}

func (t *Table) ValInCol(col, val int) error{
    //Checks if column  contains a value
    //Return error if it does nil otherwise
    
    for i := 0; i < 9; i++ {
        if t.Values[i][col].value == val{
            return errors.New("Value in column")
        }
    }
    return nil
}

func (t *Table) ValInCube(row, col, val int) error{
    // Check if the value provided is within a cube
    // Return an error if value is found nil othervise
    cube := FindCube(row, col)
    for r:=3*(cube/3); r< 3*(cube/3)+3; r++{
        for c:=3*(cube%3); c<3*(cube%3)+3; c++{
            if val == t.Values[r][c].value{
                return errors.New("Value in cube")
            }
        }
    }
    
    return nil
}

func (t *Table) Insert(row, col, val int) error{
    // Like Table.Init(row, col, val) but checks if the field is mutable
    // Return an error for invalid value or if a field is not mutable
    var err error
    err=t.ValInRow(row, val)
    if err == nil{
        err=t.ValInCol(col, val)
    } 
    if err == nil{
    err=t.ValInCube(row, col, val)
    } 
    if err == nil{
        err= t.Values[row][col].FillIn(val)
    }
    
    return err
}

func (t *Table) Init(row, col, val int) error{
    // Used to initialize all of the table with immutable values 
    // immutable fields can be mad mutable by using the Table.Mask() to hide them
    // Return an error type if value is invalid otherwise 
    
    var err error
    err=t.ValInRow(row, val)
    if err == nil{
        err=t.ValInCol(col, val)
    } 
    if err == nil{
    err=t.ValInCube(row, col, val)
    } 
    if err == nil{
        t.Values[row][col].Init(val)
    }
    
    return err
}

func (t Table) String() string{
    // Satisfy the Stringer interface
    res := ""
    
    for i:=0; i < 9; i++{
        res+="["
        for j:=0; j < 9; j++{
            res+= fmt.Sprintf(" %v", t.Values[i][j].value)
            switch j{
                case 2,5:
                    res+=" ] ["
            }
        }
        
        res += " ]\n"
        switch i{
            case 2,5:
                res += "-----------------------------\n"
        }
    }
    return res
}
func (t *Table) Mask(){
    // Masks random fields within the table
    // Masks anywhere between 5 and 7 fields per row
    // Masked fields are made mutable
    for i:=0; i<9; i++{
        mask := rand.Intn(3) + 5
        for c:=0; c< mask; c++{
            x:= make(map[int]bool)
            temp := rand.Intn(9)
            
            if x[temp] == false{
                t.Values[i][temp].Mask()
                x[temp] = true
            }
        }
    }
}
func (t *Table) Randomize(){
    //Randomizes the first row and column of the table
    
    rand.Seed(int64(time.Now().Nanosecond()))
    for i:=0 ; i<9; i++{
        val:= rand.Intn(9)+1
        err := t.Init(0,i, val)
        if err != nil{
            i--
        }
    }
    
    for i:=1 ; i<9; i++{
        val:= rand.Intn(9)+1
        err := t.Init(i,0, val)
        if err != nil{
            i--
        }
    }
}


