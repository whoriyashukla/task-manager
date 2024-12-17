package main

// import (
//     "fmt"

// )
// const s string = "constant"

// VALUES
// fmt.Println("Hello, World!")
// fmt.Println("go"+"Lang")
// fmt.Println("1+1=", 1+1)
// fmt.Println("7.0/3.0 =", 7.0/3.0)
// fmt.Println(true && false)
// fmt.Println(true || false)
// fmt.Println(!true)
// VARIABLES
// var a = "initial"
// fmt.Println(a)

// var b, c int = 1, 2
// fmt.Println(b+c)

// var d = true
// fmt.Println(d)
// var e int
// fmt.Println(e)

// f := "apple"
// fmt.Println(f)
// CONSTANTS
// fmt.Println(s)
// const n = 500000000

// const d = 3e20 / n
// fmt.Println(d)

// fmt.Println(int64(d))
// fmt.Println(math.Sin(n)) //imported math for this

// FOR LOOP //remove imported math if not in use

// i := 1
// for i<=3 {
//     fmt.Println(i)
//     i = i+1;
// }
// for j := 0; j < 3; j++ {
//     fmt.Println(j)
// }
// for i:= range 5{
//     fmt.Println(i)
// }
// for {
//     fmt.Println("loop")
//     break
// }
// for n := range 6 {
//     if n%2 == 0 {
//         continue
//     }
//     fmt.Println(n)
// }

// IF/ELSE STATEMENTS
// if 7%2 == 0 {
//     fmt.Println("7 is even")
// } else {
//     fmt.Println("7 is odd")
// }

//  if num := 9; num < 0 {
//     fmt.Println(num, "is negative")
// } else if num < 10 {
//     fmt.Println(num, "has 1 digit")
// } else {
//     fmt.Println(num, "has multiple digits")
// }

// SWITCH STATEMENT(IMPORT TIME)
// i := 3
// fmt.Println("Write", i, "as:")
// switch i {
// case 1:
//     fmt.Println("one")
// case 2:
//     fmt.Println("two")
// case 3:
//     fmt.Println("three")
// }

// switch time.Now().Weekday(){
// case time.Saturday, time.Sunday:
//     fmt.Println("It's a weekend!!!")
// default:
//     fmt.Println("It's a weekday!")
// }

// t := time.Now()
// switch {
// case t.Hour() < 12:
//     fmt.Println("It's before noon")
// default:
//     fmt.Println("It's after noon")
// }

// whatAmI := func(i interface{}) {
//     switch t := i.(type) {
//     case bool:
//         fmt.Println("I'm a bool")
//     case int:
//         fmt.Println("I'm an int")
//     default:
//         fmt.Printf("Don't know type %T\n", t)
//     }
// }
// whatAmI(true)
// whatAmI(1)
// whatAmI("hey")

// messages := make(chan string)
// go func() { messages <- "Hi!" }()
// fmt.Println(<-messages)


