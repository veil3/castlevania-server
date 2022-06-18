package logging

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

const tag = "[CASTLEVANIA]"

func Motd() {

	r := color.New(color.FgRed)
	g := color.New(color.FgHiBlack)

	g.Println(
		`
                   .d:....:h.
                .:!!!!!!!!!!!!:.
           .::!!!!!!!!!!!!!!!!!!!!::.
    ..::!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!::..
..'''. . eeee .... ... '~' ... .... eeee . .'''..
 '!h:. $ $$$$ $$$$ $$$$b d$$$$ $$$$ $$$$ $ .:h!'
  '!!!!. '$$$ '$$' '$$$' '$$$' '$$' $$$'.!!!!!'
    '!!!!.'$$ .   .  ......   .   . $$'.!!!!'
     '!!!! $$ !!!!!!!!!!!!!!!!!!!!! $$ !!!!'
       '!!h ' !!!!!!!!!!!!!!!!!!!!! ' d!!'
         '!h !!!!!!!!!!!!!!!!!!!!!!! d!'
		  ''!!!!!!!!!!!!!!!!!!!!!!!!''
             ''!!!!!!!!!!!!!!!!!!''
                 '''!!!!!!!!'''

	 [  C  A  S  T  L  E  V  A  N  I  A ]
`)
	g.Print("    Author: ")
	r.Print("github.com/CycloneOrg\n")
	r.Print(tag)
	fmt.Printf(" Welcome, ")
	r.Print(os.Getenv("USERNAME"))
}

func getCurrentTime() string {
	return time.Now().Format("15:04")
}

func Success(message string) {
	g := color.New(color.FgGreen)
	r := color.New(color.FgRed)
	date := getCurrentTime()

	r.Print(tag)
	fmt.Printf(" [%s] ", date)
	g.Print("[OK]")
	fmt.Printf(" %s\n", message)
}

func Info(message string) {
	c := color.New(color.FgCyan)
	r := color.New(color.FgRed)
	date := getCurrentTime()

	r.Print(tag)
	fmt.Printf(" [%s] ", date)
	c.Print("[INFO]")
	fmt.Printf(" %s\n", message)
}

func Warning(message string) {
	r := color.New(color.FgRed)
	date := getCurrentTime()

	r.Print(tag)
	fmt.Printf(" [%s] ", date)
	r.Print("[WARNING]")
	fmt.Printf(" %s\n", message)
}

func Error(message string) {
	r := color.New(color.FgRed)
	date := getCurrentTime()

	r.Print(tag)
	fmt.Printf(" [%s] ", date)
	r.Print("[ERROR]")
	fmt.Printf(" %s\n", message)

	os.Exit(1)
}

func Response(rw http.ResponseWriter, message string) {
	fmt.Fprintln(rw, message)
}
