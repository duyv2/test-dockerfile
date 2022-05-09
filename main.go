package main

import (
  "fmt"
  "strings"
)
func main(){
  var test string="abc def"

  var result string= reverseSentence(test)

  fmt.Println(result)
}
func reverseWord (word string) string{
  var result string;
  for i:=len(word)-1;i>=0;i--{
    result+=string(word[i])
  }
  return result
}

func reverseSentence(chuoi string) string{
  var result string;
  arrWord := strings.Fields(chuoi)
  for i:=0;i<len(arrWord);i++{
    if i==(len(arrWord)-1){
      result+=reverseWord(arrWord[i])
    }else{
      result+=reverseWord(arrWord[i])+" "
    }
  }
  return result
}
