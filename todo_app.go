package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var name string
	//var tasks []string
	reader := bufio.NewReader(os.Stdin)

	name := getName(reader)
	fmt.Printf("hello,%s\n", name)
	/*fmt.Println("Your Today's Task:")
	for {
		//	var task string
		fmt.Print("Task:")
		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)
		if strings.ToLower(task) == "done" {
			break
		}
		tasks = append(tasks, task)

	}*/
	// Collect tasks
	tasks := addTask(reader)
	printTask(tasks)

}

// *bufio.Reader upper main function is liye use kiya hai
func getName(reader *bufio.Reader) string {
	var name string
	fmt.Println("Enter the name:")
	fmt.Scanf("%s", &name)
	//fmt.Println("hello,", name)
	reader.ReadString('\n')

	return name

}

func printTask(tasks []string) {

	fmt.Println("Your task for today are:")
	if tasks == nil {
		fmt.Println("Today is free day")

	} else {
		for i, t := range tasks {
			//so here i+1 is used to write the number like  1,2,3,4,5.hello etc , indexing 0 se hoti hai ok
			fmt.Printf("%d.%s\n", i+1, t)
		}
	}
}
func addTask(reader *bufio.Reader) []string {
	var tasks []string
	fmt.Println("Your Today's Task:")
	for {
		//	var task string
		fmt.Print("Task:")
		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)
		if strings.ToLower(task) == "done" {
			break
		}
		tasks = append(tasks, task)
	}
	return tasks
}
func deleteTask(reader *bufio.Reader, tasks []string) []string {
	//TODO:complete the delete functionality

	fmt.Println("Enter the user task you want to delete (id number):")
	fmt.Scan("%d")
	return tasks

}
