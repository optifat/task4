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

func DifferentWordsCount(c string) int {
    str := ``
    x := make(map[string]bool)
    for i := 0; i < len(c); i++ {
        for ;(i < len(c) && unicode.IsLetter(rune(c[i]))) ; i++{
            str = str + string(unicode.ToLower(rune(c[i])))
        }
        if(len(str) > 0){
            x[str] = true;
        }
        str = ``
    }
    return len(x)
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
