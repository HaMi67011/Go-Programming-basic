package main
import "fmt"


//declare function by func 
func main() {
    
    //declare variables
    var a int16
    var b uint
    var c float64
    var d string
    
    //assigning values
    a=6
    b=10
    c=10.7
    d="hello"
    
    
    //dynamic 
    
    y := 32
    
    //printing type 
    fmt.Printf("type is %T",y)
    
    //constant 
    const x uint = 10
    
    //printing
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d) 
    
    //if else
    if (y>0){
        fmt.Println("val greater then 0")
    }else if (y==0){
        fmt.Println("val is equal to 0")
    }else{
        fmt.Println("val is less then 0")
    }
    
    //switch
    switch{
        case y>0:
            fmt.Println("val is equal to 0")
        case y>0:
           fmt.Println("val greater then 0")
        default:
            fmt.Println("default")
    }
    
    //for loop
    //just like c++
    for i:=0;i<5;i++{
        //break and continue like other languages
        if i<5{
            continue
        }
        if i>5{
            break
        }
    }
    
    //function
    var res uint = findmax(3,5)
    fmt.Println(res)
    
    //1d arrays
    
    //initialized
    var arrIn = [3] uint{1,2,3}
    
    //uninitialized
    var arrUn = [3] uint{}
    
    //size not declare
    var arrSizeNotDef = [] uint {7,8,9}
    
    var q uint = 4
    for i:=0;i<3;i++{
        arrUn[i] = q
        q++
    }
    
    //display
    for j:=0;j<3;j++{
        fmt.Printf("Elements [%d] = %d\n",j,arrIn[j])
        fmt.Printf("Elements [%d] = %d\n",j,arrUn[j])
        fmt.Printf("Elements [%d] = %d\n",j, arrSizeNotDef[j])
    }
    
    //2d array
    //initialized
    var arr2dIn = [2][2] uint { {0,0} ,{1,1}} 
    
    //uninitialized
    var arr2dUn = [2][2] uint {{}}
    
    //undefined size
    var arr2dUnDef = [][] uint { {4,4} ,{5,5}} 
    
    var t uint = 2
    
    for i:=0;i<2;i++{
        for j:=0;j<2;j++{
            arr2dUn[i][j] = t
        }
        t++
    }
    
    //display
    for i:=0;i<2;i++{
        for j:=0;j<2;j++{
            fmt.Printf("Elements [%d][%d] = %d\n",i,j,arr2dIn[i][j])
            fmt.Printf("Elements [%d][%d] = %d\n",i,j,arr2dUn[i][j])
            fmt.Printf("Elements [%d][%d] = %d\n",i,j,arr2dUnDef[i][j])
        }
    }
    
    //array passing into function
    
    var avg float32 = arrPassWithSize(arrIn)
    fmt.Printf("avg is %f\n",avg)
    
    
    var avgres float32 = arrPassWithOutSize(arrSizeNotDef,3)
    fmt.Printf("avg is %f\n",avgres)
    
    
    //pointers
    //same as in c++
    var intpointer *int
    var vI int=20
    
    var floatpointer *float32
    var vF float32=30
    
    var Strpointer *string
    var vS string="Heloo"
    
    intpointer = &vI
    floatpointer=&vF
    Strpointer=&vS

    //double pointers
    
    var pointer_pointer **int
    
    pointer_pointer=&intpointer

    //displaying the address
    fmt.Printf("%x\n",intpointer)
    fmt.Printf("%x\n",floatpointer)
    fmt.Printf("%x\n",Strpointer)
    fmt.Printf("%x\n",pointer_pointer)
    
    //passing to function just like c++
    inc(&vI)
    fmt.Printf("%d",vI)
    
    //struct like in c++
    var book1 Books
    
    book1.title = "FAST"
    book1.id=1122
    
    fmt.Printf("%s\n",book1.title)
    fmt.Printf("%d\n",book1.id)
    
    printstruct(book1)
    
    //recursion like other languages
    var factres int = fact(5)
    fmt.Printf("%d\n",factres)
    
    //slicing just like we do in python 
    arr := [5] int{1,2,3,4,5}
    slc:=arr[0:3]
    fmt.Println(slc)
    fmt.Println(len(slc))
    //append fucntion
    slc = append(slc,45)
    fmt.Println(slc)
    
    //variadic fucntion
    variadic(1,2,3,4,5,6,7)
    
    //maps just like dict in python
    var mymaps map[string] int
    mymaps = map[string]int {}
    mymaps["a"] = 1
    mymaps["b"] = 2
    mymaps["c"] = 3
    
    //displaying
    fmt.Println(mymaps)
    
    delete(mymaps,"c")
    
    fmt.Println(mymaps)
    
}
//parameters types and return types
func findmax(no1,no2 uint) uint{
    if no1 > no2{
        return no1
    }else if no1 < no2{
        return no2
    }
    return 0
}

func arrPassWithSize(arr[3] uint) float32{
    var sum uint
    var avg float32
    for i:=0;i<3;i++{
        sum+=arr[i]
    }
    avg=float32(sum/3)
    return avg
}


func arrPassWithOutSize(arr[] uint,size uint) float32{
    var sum uint
    var avg float32
    var isize int = int(size)
    for i:=0;i<isize;i++{
        sum+=arr[i]
    }
    avg=float32(sum/size)
    return avg
}

//pointer passing in function
func inc (x *int){
    *x++
}

//struct
type Books struct{
    title string
    id int
}

//struct passing in function
func printstruct (book Books){
    fmt.Printf("%s\n",book.title)
    fmt.Printf("%d\n",book.id)
}

//recursion example
func fact (no int) int {
    if (no <=1){
        return 1
    }
    return no*fact(no-1)
}

//variadic function
//infinte amount of parameters you can pass
func variadic(params ... int){
    fmt.Printf("%T\n %v\n",params,params)
}
