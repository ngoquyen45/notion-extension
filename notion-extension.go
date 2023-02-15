package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// replace this string
	oldStr := `const windowCreationArgs = Object.assign(Object.assign({}, desiredWindowBounds), { show: false, backgroundColor: colorHelpers_1.electronColors.notionBackground[state_1.Store.getState().app.theme.mode], titleBarStyle: "hiddenInset", trafficLightPosition: isTabsFeatureEnabled`
	// with this string
	newStr := `const windowCreationArgs = Object.assign(Object.assign({}, desiredWindowBounds), { show: false, frame: false, backgroundColor: colorHelpers_1.electronColors.notionBackground[state_1.Store.getState().app.theme.mode], titleBarStyle: "hiddenInset", trafficLightPosition: isTabsFeatureEnabled`

	// get the file path
	appData, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(appData, "AppData", "Local", "Programs", "Notion", "resources", "app", "main", "WindowController.js")

	// read the file
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	content := string(bytes)

	// replace the string
	if len(os.Args) > 1 && os.Args[1] == "-reset" {
		content = strings.Replace(content, newStr, oldStr, -1)
	} else {
		content = strings.Replace(content, oldStr, newStr, -1)
	}

	// write the updated content back to the file
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}
