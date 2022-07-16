package logs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func RegisterLogs(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func PrintLogs() {
	fmt.Println("\n Showing Logs:")
	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))
}
