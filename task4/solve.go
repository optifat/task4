package main

import "unicode"

func RemoveEven(arr []int) []int{
  result := []int{}
  for _, element := range arr{
    if element%2 == 0{
      result = append(result, element);
    }
  }
  return result;
}

func PowerGenerator(input int) func() uint{
  i := uint(2);
  return func() (ret uint){
    ret = i
    i *= 2
    return
  }
}

func DifferentWordsCount(text string)  int{
  var result string
  var word string
  myMap := map[string]bool{}
  for _, symbol := range text{
    if unicode.IsLetter(symbol){
      result = result + string(unicode.ToLower(symbol))
    } else{
      result = result + string(` `)
    }
  }
  for _, symbol := range result{
    if unicode.IsSpace(symbol){
      myMap[word] = true
      word = string(``)
    } else{
      word = word + string(symbol)
    }
  }
  delete(myMap, string(``))
  fmt.Println(myMap)
  return len(myMap);
}
