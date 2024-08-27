package util

import (
	"github.com/alexpfx/linux_wrappers/wrappers/wtype"
	"github.com/alexpfx/linux_wrappers/wrappers/xdtype"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func BashExec(script string) string {
	cmd := exec.Command("bash", "-c", script)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar script: %v", err)
	}
	return string(output)
}

func Init(actionDir string, filename string, data []byte) error {
	file := filepath.Join(actionDir, filename)

	err := os.MkdirAll(actionDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Erro ao criar diretório de ação: %v", err)
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		err = os.WriteFile(file, data, 0644)
		if err != nil {
			log.Fatalf("Erro ao criar arquivo de ação: %v", err)
			return err
		}
	}
	return err

}

func Typeit(text string) {
	stype := os.Getenv("XDG_SESSION_TYPE")
	if stype == "wayland" {
		w := wtype.New(wtype.Builder{
			DelayBetweenKeyStrokes: "5",
			DelayBeforeKeyStrokes:  "501",
		})

		w.Type(strings.TrimSpace(text))
		return
	}

	x := xdtype.New(xdtype.Builder{
		Delay: "50",
	})

	x.Type(strings.TrimSpace(text))

}
