        
package main // src/hello/hello.go 
 
import . "fmt"
import . "hellosub"
 
type C1 class {
    C1( s1 string )
    Huh() string
    S1 string
}

func ( c C1 ) C1( s1 string ) {
    c.S1 = s1
}

func ( c C1 ) Huh() string {
    return c.S1
}


type C2 class extends C1 { 
    C2( s1 string, hi string  ) 
    Wow() string 
    s2 string  
} 

func ( cc C2 ) C2( s1 string, hi string ) { 
    cc.C1( s1 ) 
    cc.s2 = hi 
} 

func ( cc C2 ) Wow() string { 
    return cc.s2 + " " + cc.Huh() + "!"  
} 


func main() {
    var hi Hi = new(Hello)   
    hi.Init() 

    var c C1 = new( C1 )  
    c.C1( "World" ) 

    var cc C2 = new( C2 )  
    cc.C2( "World", hi.Huh() ) 

    Println( cc.Wow() )
    Printf( "\nWelcome to Go++ - the friendlier golang...\n\n" )  
}
