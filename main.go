package main

import (
	"fmt"
	"github.com/alexpfx/linux_wrappers/wrappers/rofi"
	"github.com/alexpfx/linux_wrappers/wrappers/wtype"
	"log"
	"time"
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

func typeIt(text string) {
	w := wtype.NewWType(wtype.WTypeBuilder{
		DelayBetweenKeyStrokes: "5",
		DelayBeforeKeyStrokes:  "50",
	})
	w.Run(text)
}

func getDate() string {
	mmdd := time.Now().Format("02/01")
	typeIt(mmdd)
	return mmdd
}
