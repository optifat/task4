package main

import "fmt"

func RemoveEven(arr []int) []int{
  result := []int{}
  for _, element := range arr{
    if (element%2 == 0){
      result = append(result, element);
    }
  }
  return result;
}

func PowerGenerator (input int){
  number *= 2;
  fmt.Println(number)
}

//func main(){
//  array := []int{ 98, 93, 77, 82, 83 }
//  fmt.Println(RemoveEven(array));
//  PowerGenerator(2);
//  PowerGenerator(2);
//  PowerGenerator(2);
//}
