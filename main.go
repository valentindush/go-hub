package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("New bill " + "\"" +b.name + "\"" + " created")

	return b
}

func main()  {

	mybill := createBill()
	mybill.updateTip(10)
	mybill.addItem("wine", 20.5)

	fmt.Print(mybill.format())

}