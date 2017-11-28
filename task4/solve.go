package main

//import "fmt"
import "unicode"

func RemoveEven(arr []int) []int{
  result := []int{}
  for _, element := range arr{
    if element%2 == 1{
      result = append(result, element);
    }
  }
  return result;
}

func PowerGenerator(b int) (func() int){
    res := 1
    return func() (ret int){
        ret = res*b
        res = ret
        return
    }

}

func DifferentWordsCount(text string) int {
    word := ``
    myMap := make(map[string]bool)
    for i := 0; i < len(text); i++ {
        for ;(i < len(text) && unicode.IsLetter(rune(text[i]))) ; i++{
            str += string(unicode.ToLower(rune(text[i])))
        }
        if(len(word) > 0){
            myMap[word] = true;
        }
        str = ``
    }
    return len(myMap)
}
/*
func main(){
  array := []int{ 98, 93, 77, 82, 83 }
  fmt.Println(RemoveEven(array));
  gen := PowerGenerator(2)
  fmt.Println(gen()) // Должно напечатать 2
  fmt.Println(gen()) // Должно напечатать 4
  fmt.Println(gen()) // Должно напечатать 8
  fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
}
*/
