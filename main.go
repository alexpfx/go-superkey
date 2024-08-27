package main

import (
	_ "embed"
	"github.com/alexpfx/go-superkey/action"
	"github.com/alexpfx/go-superkey/util"
	"github.com/alexpfx/linux_wrappers/wrappers/rofi"

	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

//go:embed actions/actions.yaml
var defaultActionFile []byte
var userConfigDir, _ = os.UserConfigDir()
var appActionDir = filepath.Join(userConfigDir, "go-superkey", "actions")

func main() {
	err := util.Init(appActionDir, "actions.yaml", defaultActionFile)
	if err != nil {
		log.Fatal(err)
	}
	actionFiles, _ := loadActionFiles()
	actionMap := make(map[rune]rofi.KeyAction)
	for _, filename := range actionFiles {
		var actF action.ActionsFile
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(data, &actF)
		if err != nil {
			continue
		}

		for _, act := range actF.Actions {
			if act.Key == "" {
				continue
			}
			actionMap[rune(act.Key[0])] = rofi.KeyAction{
				Label: act.Label,
				Action: func() string {
					sessionType := getSessionType()
					if cmd, ok := act.Scripts[sessionType]; ok {
						out := util.BashExec(cmd)
						return out
					}
					cmd := act.Scripts["default"]
					out := util.BashExec(cmd)
					return out
				},
			}
		}

	}

	kbm := rofi.NewKeyboardMenu(actionMap)
	out, err := kbm.Show()
	if err != nil {
		log.Fatal(err)
	}
	util.Typeit(out)
}

func getSessionType() string {
	return os.Getenv("XDG_SESSION_TYPE")
}

func loadActionFiles() ([]string, error) {
	acts := make([]string, 0)

	filepath.WalkDir(appActionDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".yaml" {
			acts = append(acts, path)
		}
		return err
	})
	return acts, nil

}
