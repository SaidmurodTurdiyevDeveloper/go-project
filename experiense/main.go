package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Birinchi sonni kiriting ")
	birinchi_son, _ := reader.ReadString('\n')
	fmt.Print("Birinchi sonni kiriting ")
	ikkinchi_son, _ := reader.ReadString('\n')
	first_number, _ := strconv.ParseInt(strings.TrimSpace(birinchi_son), 0, 8)
	second_number, _ := strconv.ParseInt(strings.TrimSpace(ikkinchi_son), 0, 8)
	natija := first_number + second_number
	fmt.Println(natija)

}
