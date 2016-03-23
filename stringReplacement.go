package main
// Program is an excercise to find and replace words in a string.
// Author: TS Burns

import "fmt"
import "strings"


func main() {
    target := "apes"
    replace := "monkeys"
    sentence := "The apes like the bananas."

    if strings.Contains(sentence, target){
// split sentence into an array
      sentenceArray := strings.Split(sentence, " ")
// for loop for length of array.
      for i := 0; i < len(sentenceArray); i++{
// comparsion of array with target word
        if sentenceArray[i] == target {
// assignemnt of array index with replacement word
          sentenceArray[i] = replace
          fmt.Println(sentenceArray)
          if sentenceArray[i] == replace {
// breaking of loopl
            break
          }
        } else {
          if sentenceArray[i] == replace {
            break
          }
        }
      }
    }

}
