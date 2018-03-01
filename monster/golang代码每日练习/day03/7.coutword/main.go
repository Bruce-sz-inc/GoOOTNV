package main

import (
	"io/ioutil"
	"os"
	"log"
	"strings"
	"fmt"
)
// 统计词频
func main(){
	content ,err := ioutil.ReadFile(os.Args[1])
	if err != nil{
		log.Fatal(err)
	}
	count := make(map[string]int)

	words :=strings.Fields(string(content))

	for _,word := range words{
		_,ok := count[word]
		if ok {
			count[word]+=1
		}else{
			count[word] = 1
		}
	}
	for word ,cnt := range count {
		fmt.Println(word,cnt)
	}


}


