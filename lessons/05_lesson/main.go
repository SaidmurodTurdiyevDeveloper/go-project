package main

import (
	"fmt"
	"sort"
)

// this is Slices
func main() {
	fmt.Println("\nSlices\n")

	// var fruitList = []string{}
	// //add new item
	// fruitList = append(fruitList, "Apple")
	// fruitList = append(fruitList, "Banana", "Orange")
	// fruitList = append(fruitList, "Cherry", "Peach")
	// fruitList = append(fruitList, "Strawberry", "Pineapple", "Ananas", "Apricot", "Pear")
	// fruitList = append(fruitList)

	// fmt.Println("Our Fruits ", fruitList)
	// fmt.Println("fruit list size ", len(fruitList))

	// // add with input
	// var input = bufio.NewReader(os.Stdin)

	// fmt.Print("Add new fruit ")
	// newFruit, error := input.ReadString('\n')

	// if error == nil {
	// 	//clear "\n" because added every time
	// 	newFruit = strings.TrimSpace(newFruit)
	// 	fruitList = append(fruitList, newFruit)
	// } else {
	// 	fmt.Println("new frouit did not add cause ", error)
	// }
	// fmt.Println("After added input data ", fruitList)
	// fmt.Println("fruit list size ", len(fruitList))
	// fmt.Println("\n=================\n")

	// //slice cut any start index example startIndex=2
	// fruitList = append(fruitList[2:])
	// fmt.Println("After cutted start index=2 list ", fruitList)
	// fmt.Println("fruit list size ", len(fruitList))
	// fmt.Println("\n=================\n")

	// //slice cut any start  and end indexes example startIndex=1 endIndex=5
	// fruitList = append(fruitList[1:5])
	// fmt.Println("After cutted start index=1 and end index=5 list ", fruitList)
	// fmt.Println("fruit list size ", len(fruitList))
	// fmt.Println("\n=================\n")

	// //slice cut any end index example endIndex=3
	// fruitList = append(fruitList[:3])
	// fmt.Println("After cutted end index=3 list ", fruitList)
	// fmt.Println("fruit list size ", len(fruitList))
	// fmt.Println("\n=================\n")

	var vegitables = make([]string, 6)
	//added with
	vegitables[0] = "Poteto"
	vegitables[1] = "Garlic"
	vegitables[2] = "Carrot"

	fmt.Println("Before append list size ", len(vegitables))
	fmt.Println("\n=================\n")

	//if with add with append list size is changed
	vegitables = append(vegitables, "Tometo")
	fmt.Println("After append list size ", len(vegitables))

	fmt.Println("\n=================\n")

	// fmt.Println("Our Vegitables\n")
	// for i := 0; i < len(vegitables); i++ {
	// 	fmt.Println(i, vegitables[i])
	// }
	// fmt.Println("\n=================\n")

	//sort vegitables
	vegitables[3] = "Onion"
	vegitables[4] = "Cucumber"
	vegitables[5] = "Mushroom"
	fmt.Println(vegitables)

	fmt.Println("Slice is sorted ", sort.StringsAreSorted(vegitables))
	sort.Strings(vegitables)
	fmt.Println(vegitables)
	fmt.Println("Slice is sorted ", sort.StringsAreSorted(vegitables))
	fmt.Println("\n=================\n")

	//remove value of slices based index
	// var index int = 2
	// vegitables =  append(vegitables[:index], vegitables[index+1:]...)

	//this is with pointer
	remove(1, &vegitables)
	fmt.Println("after remove")
	fmt.Println(vegitables)
	add(&vegitables, "apricot")
	addIndex(0, &vegitables, "mushrum")
	addIndex(3, &vegitables, "murshmellow")
	addIndex(7, &vegitables, "last")
	fmt.Println("after add")
	fmt.Println(vegitables)

}

func remove(index int, list *[]string) {
	var ls = *list
	if index < len(*list) {
		*list = append(ls[:index], ls[index+1:]...)
	}
}
func addIndex(index int, list *[]string, args ...string) {
	var ls = *list
	if index == 0 {
		*list = append(args, ls...)
	} else if index > 0 && index < len(*list) {
		*list = append(ls[:index], args...)
		*list = append(*list, ls[index+1:]...)
	}
}
func add(list *[]string, args ...string) {
	var ls = *list
	*list = append(ls, args...)
}
