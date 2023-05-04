package main

import (
    "os"
    "bufio"  
    "fmt"
    "strings"
    "runtime"
)

const(
    DOT byte = 46
)

type InputBuffer struct {
    buffer string
}

type MetaCommandResult string
const (
    META_COMMAND_UNRECOGNIZED_COMMAND MetaCommandResult = "Unrecognized_Meta"
    META_COMMAND_SUCCESS MetaCommandResult = "META_COMMAND_SUCCESS"
)

type PrepareResult string
const (
    PREPARE_UNRECOGNIZED_STATEMENT PrepareResult = "Unrecognized_Prepare"
    PREPARE_SUCCESS PrepareResult = "PREPARE_SUCCESS"
)

type StatementType string
const (
    STATEMENT_UNDEFINED StatementType ="UNDEFINED"
    STATEMENT_INSERT StatementType ="INSERT"
    STATEMENT_SELECT StatementType ="SELECT"
)

type Statement struct {
    statementType StatementType
}

func prepare_statement(input *InputBuffer, statement *Statement) PrepareResult {

    if strings.Compare(input.buffer,"select")==0 {
        statement.statementType = STATEMENT_SELECT
        return PREPARE_SUCCESS

    } else if strings.Compare(input.buffer,"insert")==0 {
        statement.statementType = STATEMENT_INSERT
        return PREPARE_SUCCESS
    }
    return PREPARE_UNRECOGNIZED_STATEMENT

}

func execute_statement(statement *Statement){
    switch(statement.statementType){
    case STATEMENT_INSERT:
        print("INSERTING\n")
    case STATEMENT_SELECT:
        print("SELECTING\n")
    }
}

func printPrompt(){
    fmt.Print(" db > ")
}

func print(val string){
    fmt.Print(val)
}

func do_meta_command(input *InputBuffer) MetaCommandResult {
    if strings.Compare(input.buffer,".exit")==0 {
        exit(0)
     }
     return META_COMMAND_UNRECOGNIZED_COMMAND
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
    for true {        
        printPrompt()
        readInput(inputBuffer)

         if(inputBuffer.buffer[0]==DOT){ 
            
            switch do_meta_command(inputBuffer) {
            case META_COMMAND_SUCCESS:
                continue
            case META_COMMAND_UNRECOGNIZED_COMMAND:   
                print("Unrecognized command "+inputBuffer.buffer+"\n")             
            default:
                print("Unrecognized command "+inputBuffer.buffer+"\n")                
            }

            continue

         }

         var statement *Statement
         statement = &Statement{statementType:STATEMENT_UNDEFINED,}
         switch(prepare_statement(inputBuffer,statement)){
         case PREPARE_SUCCESS:
            break
         case PREPARE_UNRECOGNIZED_STATEMENT:
            fmt.Println(fmt.Sprintf("Unrecognized keyword at start of %s",inputBuffer.buffer))
        default:
            fmt.Println(fmt.Sprintf("Unrecognized keyword at start of %s",inputBuffer.buffer))
         }


         execute_statement(statement)
        
    }
}