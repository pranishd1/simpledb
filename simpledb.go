package main

import (
    "os"
    "bufio"  
    "fmt"
    "strings"
    "runtime"
)

type InputBuffer struct {
    buffer string
}

func printPrompt(){
    fmt.Print(" db > ")
}

func print(val string){
    fmt.Print(val)
}

func printI(val int){
    fmt.Print(val)
}

func readInput(input *InputBuffer) {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    if runtime.GOOS == "windows" {
        text = strings.TrimRight(text, "\r\n")
      } else {
        text = strings.TrimRight(text, "\n")
      }
    input.buffer = text
}

func exit(val int){
    os.Exit(val)
}

func main() {
    var inputBuffer *InputBuffer
    inputBuffer = &InputBuffer{buffer:"",}
    c := 0
    for true {
        
        printPrompt()
        readInput(inputBuffer)
        if strings.Compare(inputBuffer.buffer,".exit")==0 {
           exit(0)
        }else{
            print("Unrecognized command "+inputBuffer.buffer+"\n")
        }
        c = c + 1
    }
}