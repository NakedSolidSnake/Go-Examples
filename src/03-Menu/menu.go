package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	introduction()
	for {
		showOptions()
		executeOption(getOption())
	}

}

func introduction() {
	name := "Cristiano"
	version := 1.1
	fmt.Println("Hello, Sir.", name)
	fmt.Println("Este programa está na versão", version)
}
func showOptions() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func getOption() int {
	var command int
	fmt.Scan(&command)
	return command
}

func executeOption(option int) {
	switch option {
	case 1:
		monitoring()
	case 2:
		showLogs()
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Option unavailable")
		os.Exit(-1)
	}
}

func monitoring() {
	fmt.Println("Monitoring...")
	sites := readSitesFromFile()

	for i, site := range sites {
		fmt.Println("Position", i, "Site:", site)
		siteTest(site)
	}
}

func siteTest(site string) {
	answer, _ := http.Get(site)
	var state bool

	if answer.StatusCode == 200 {
		state = true
	} else {
		state = false
	}

	logRegister(site, state)
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()

	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(file)

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + "- online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	fmt.Println("Showing Logs...")

	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
