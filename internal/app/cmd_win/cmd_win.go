// package cmd_win
// https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
// https://go.dev/doc/tutorial/handle-errors
// https://stackoverflow.com/questions/7151261/append-to-a-file-in-go
// https://gobyexample.com/writing-files
// https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console
// Перейдя к папке с файлом, нажмите клавишу с логотипом Windows  + R,
//  напечатайте shell:startup, затем нажмите ОК. Откроется папка Автозагрузка.
// sc.exe create "wuauserv_stopper" binPath= "C:\Users\Alex1\Documents\cmd_win.exe"
// https://ru.stackoverflow.com/questions/526702/%D0%9A%D0%B0%D0%BA-%D1%81%D0%BE%D0%B7%D0%B4%D0%B0%D1%82%D1%8C-%D1%81%D0%BB%D1%83%D0%B6%D0%B1%D1%83-windows-%D0%BD%D0%B0-golang
// https://stackoverflow.com/questions/3582108/create-windows-service-from-executable
// http://nssm.cc/builds
// C:\Users\Alex1\Documents\nssm-2.24-103-gdee49fc\win64>nssm install "wuauserv_stopper" "C:\Users\Alex1\Documents\cmd_win.exe"
// Service "wuauserv_stopper" installed successfully!
//
// C:\Users\Alex1\Documents\nssm-2.24-103-gdee49fc\win64>nssm start "wuauserv_stopper"
// wuauserv_stopper: START: Операция успешно завершена.
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

const INTERVAL_SEC = 3
const LOG_FILE = "log.txt"

func cmd_log(text string) {
	// d1 := []byte("hello\ngo\n")
	// now := time.Now()
	// now_string := now.Format("YYYY-MM-DD")
	t := time.Now()
	now_string := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d\r\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(now_string)
	// _, err := os.Create(LOG_FILE)
	// check(err)
	f, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()
	// d1 := []byte(now_string + ": " +text + "\r\n")
	// err := os.WriteFile("log.txt", d1, 0644)
	result_code, err := f.WriteString(now_string)
	fmt.Println(fmt.Sprintf("cmd_log %s, %d", LOG_FILE, result_code))
	check(err)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loop_cs() (string, error) {
	// c := exec.Command("cmd", "/C", "del", "D:\\a.txt")
	// cmd, err := exec.Command("cmd", "/C", "dir").Output()
	// cmd, err := exec.Command("sc", "query", "wuauserv").Output()

	cmd, err := exec.Command("sc", "stop", "wuauserv").Output()

	if err != nil {
		// fmt.Print("-> ")
		// fmt.Println("Error: ", err.Error())
		return "", errors.New(err.Error())
	} else {
		// cmd, _ := exec.Command("cmd", "/C", "dir").Output()
		reader := transform.NewReader(bytes.NewReader(cmd), charmap.CodePage866.NewDecoder())
		d, _ := ioutil.ReadAll(reader)
		s := string(d)
		// fmt.Println("Done: ", s)
		return s, nil
	}
}
func main() {

	for {
		result, err := loop_cs()
		if err == nil {
			// fmt.Println(result)
			color.Red(result)
			cmd_log("Done")
		}
		time.Sleep(INTERVAL_SEC * time.Second)
	}
	// cmd_reader := bufio.NewReader(os.Stdin)
	// s,_ := cmd_reader.ReadString('\n')
	// fmt.Println("Done: ", s)

	// bufio.NewScanner(os.Stdin)

}
