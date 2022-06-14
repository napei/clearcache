package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

var check = green("✓")
var cross = red("✗")

func CheckExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}

func removeDir(path string) error {
	if CheckExists(path) {
		err := os.RemoveAll(path)
		if err != nil {
			return fmt.Errorf("remove dir: %v", err)
		}
	} else {
		return fmt.Errorf("path not found: %s", path)
	}
	return nil
}

func main() {

	appDataRoot, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	basePath := appDataRoot + "\\FiveM\\FiveM.app\\data\\"

	if CheckExists(basePath) {
		fmt.Println(check + " Found FiveM AppData folder")
	} else {
		fmt.Println(cross + "FiveM AppData not found...Exiting")
		return
	}

	proceed := true
	prompt := &survey.Confirm{
		Message: "Are you sure you want to clear FiveM cache?",
		Default: true,
	}
	survey.AskOne(prompt, &proceed)

	if !proceed {
		return
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	cacheDir := basePath + "cache\\"

	s.Suffix = " Cleaning cache directory..."
	err = removeDir(cacheDir)
	if err != nil {
		fmt.Println(cross + " " + err.Error())
	} else {
		fmt.Println(check + " " + cacheDir + " cleaned")
	}

	serverCacheDir := basePath + "server-cache\\"

	s.Suffix = " Cleaning server cache directory..."
	err = removeDir(serverCacheDir)
	if err != nil {
		fmt.Println(cross + " " + err.Error())
	} else {
		fmt.Println(check + " " + serverCacheDir + " cleaned")
	}

	serverCachePrivDir := basePath + "server-cache-priv\\"

	s.Suffix = " Cleaning server cache private directory...\n"
	err = removeDir(serverCachePrivDir)
	if err != nil {
		fmt.Println(cross + " " + err.Error())
	} else {
		fmt.Println(check + " " + serverCachePrivDir + " cleaned")
	}

	s.FinalMSG = check + " Cache cleaned!"
	s.Stop()
}
