               
package hellosub // src/hellosub/hellosub.go 
import . "fmt"
 
type Hi interface { 
    Init()
    Huh() string
}

type Hello struct implements Hi { 
    str string 
}


func ( h *Hello ) Copy( fromh *Hello ) { 
    h.str = fromh.str
}  


func (h *Hello) Init() { 
    var slice sliceof byte  
    var chk bool 
    hello := "olleH" 
    while NOT chk AND len(slice) < 5 { 
        slice = append( slice, hello[ len(hello) - 1 ] ) 
        hello = hello[0:len(hello)-1]  
        if len(slice) < 6 { 
	    Println( string(slice) ) 
	} 
        else { 
            chk = true 
        } 
    } 
    h.str = string( slice )    
} 


func (h *Hello) Huh() string { 
    return h.str 
} 


var Test Hi 

func init() { 
    var hlo Hello = Hello{ "What?" }
    var h2 Hello 

    h2.Copy( hlo ) // This is a golang error - Go++ let's it slide 
    Test = h2  // Ditto here since Hello methods using pointer receivers 
}  
