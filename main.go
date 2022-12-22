package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main(){
	fmt.Print("Please enter number(separate with space) : ")
	br := bufio.NewReader(os.Stdin)
	a, _, _ := br.ReadLine()
	ns := strings.Split(string(a), " ")

	var values []int
	for _, s := range(ns){
		n, _ := strconv.Atoi(s)
		values = append(values, n)
	}
	bubblesort(values)
	fmt.Print(values,"\n")
}

func bubblesort(a []int){
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a)-1-i;j++{
			if a[j+1] < a[j]{
				Swap(a,j)
			}
		}
	}
}

func Swap(a []int, j int){
	temp := 0
	temp = a[j]
	a[j] = a[j+1]
	a[j+1] = temp
}