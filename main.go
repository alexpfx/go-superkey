package main

import (
	"fmt"
	"github.com/alexpfx/linux_wrappers/wrappers/pm"
	"github.com/alexpfx/linux_wrappers/wrappers/rofi"
	"github.com/alexpfx/linux_wrappers/wrappers/wtype"
	"log"
	"strings"
	"time"
	"os"
)

func main() {
	actionMap := make(map[rune]rofi.KeyAction)
	actionMap['a'] = rofi.KeyAction{
		Label:  "Time",
		Action: getTime,
	}

	actionMap['d'] = rofi.KeyAction{
		Label:  "Date",
		Action: getDate,
	}
	actionMap['p'] = rofi.KeyAction{
		Label:  "New Pass",
		Action: genNewPass,
	}

	actionMap['c'] = rofi.KeyAction{
		Label:  "CPF",
		Action: getCpf,
	}


	kbm := rofi.NewKeyboardMenu(actionMap)
	out, err := kbm.Show()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}

func getTime() string {
	hhmm := time.Now().Format("15:04")
	typeIt(hhmm)

	return hhmm
}

func getCpf() string {
	cpf := os.Getenv("CPF")
	typeIt(cpf)

	return cpf
}

func typeIt(text string) {
	w := wtype.New(wtype.Builder{
		DelayBetweenKeyStrokes: "5",
		DelayBeforeKeyStrokes:  "50",
	})
	w.Type(strings.TrimSpace(text))
}

func getDate() string {
	mmdd := time.Now().Format("02/01")
	typeIt(mmdd)
	return mmdd
}


func genNewPass() string {
	pmg := pm.NewDefaultMin12()
	pass, err := pmg.Gen()
	if err != nil {
		log.Fatal(err)
	}
	typeIt(pass)
	return pass
}
