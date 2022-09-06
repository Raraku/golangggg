package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "time"
    "sync"
)

/*
This function is meant to solve a question to determine all possible combinations; It works but currently runs into timelimit errors
 */
 
 func addOne(num int32) int32{
     return num + 1
 }
 func addTwo(num int32) int32{
     return num + 2
 }
 func multiplyOne(num int32) int32{
     return num * 1
 }
 func multiplyTwo(num int32) int32 {
     return num * 2
 }
 
 func performOperation(num int32, ones int32, twos int32, hasBeenAddedTo bool, out chan <- int32){
     if(ones > 0){
         go performOperation(addOne(num), ones - 1 , twos, true, out)
        if(!hasBeenAddedTo){      
         go performOperation(multiplyOne(num), ones - 1 , twos, false, out)
     }
     }
     if(twos > 0 ){
         go performOperation(addTwo(num), ones, twos -1 , true, out)
        if(!hasBeenAddedTo){      
         go performOperation(multiplyTwo(num), ones , twos - 1, false, out)
     }
     }
     out <- num
 }
 func waiter(ch chan int32, wg *sync.WaitGroup, chan2 chan map[int32]int32)  {
     defer wg.Done()
    uniqueAnswers := make(map[int32]int32)   
    life := 2
      
         for {
        select {
        case s:=<-ch:
            uniqueAnswers[s] = 1
        default:
            time.Sleep(50 * time.Millisecond)
            
            if life--;life==0 {
                // fmt.Println("e don finish", uniqueAnswers)
                chan2 <- uniqueAnswers
                // counter := int32(0);
                // for range uniqueAnswers {
                //     counter +=1
                // }  
                // return counter 
                return
            }
        }
         }
 }

func onesAndTwos(a int32, b int32) int32 {
    // Write your code here
    var wg sync.WaitGroup
    ch := make(chan int32,200)
    maxNum := a + b
    if maxNum == 0 {
        return 0
    }
    uniqueAnswers := make(map[int32]int32)

    if(a > 0 ){
        performOperation(1, a-1, b, false, ch)

    }
    if (b > 0){
            go performOperation(2, a, b-1, false, ch)
    }   
    chan2 := make(chan map[int32]int32,5)
    for i:=0; i< 5; i++ {
        wg.Add(1)
        go waiter(ch, &wg, chan2)       
    }
    wg.Wait()
    for i:=0; i< 5; i++ {
        t := <- chan2
            for k := range t {
                uniqueAnswers[k] = 1
            }
        }
    counter := int32(0);
    for range uniqueAnswers {
        counter +=1
    }  
    return counter 
    //     maps.Copy(uniqueAnswers, <-chan2)
    // }
    //  maps.Copy(dst, src)
    
    // go func(){
        
    //     wg.Wait()
    //     close(ch)
    // }()

    // for i:= range ch {
    //     uniqueAnswers[i] = 1
    // }
 

    // wg.Wait()
    // close(ch)


}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    var wg sync.WaitGroup

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int(tTemp)
    answerss := make([]int32, t, t)

    for tItr := 0; tItr < int(t); tItr++ {
        wg.Add(1)
        ta := tItr
        fmt.Println(ta)

        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")
        aTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        a := int32(aTemp)
        

        bTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        b := int32(bTemp)
         go func(){
        result := onesAndTwos(a, b)
        // fmt.Println(ta)
        answerss[ta] = result
        
        wg.Done()
        }()
    }

    wg.Wait()
    for _,v := range answerss {
            fmt.Fprintf(writer, "%d\n", v)
    }


    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
