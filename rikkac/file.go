package main

import (
	"errors"
	"flag"
	pathUtil "path/filepath"
	"strings"

	"github.com/Raincal/rikka/client"
	"github.com/Raincal/rikka/common/util"
)

func readFile(filePath string) (string, []byte, error) {
	absFilePath, err := pathUtil.Abs(filePath)
	if err != nil {
		return "", nil, err
	}
	l.Debug("Change to absolute path:", absFilePath)

	if !util.IsFile(absFilePath) {
		return "", nil, errors.New("Path " + absFilePath + " not exists or not a file")
	}
	l.Debug("File", absFilePath, "exists and is a file")

	fileContent, err := client.CheckFile(absFilePath)
	if err != nil {
		return "", nil, err
	}
	return absFilePath, fileContent, nil
}

func getFile(index int) (string, bool) {
	filepath := ""
	if len(flag.Args()) > index {
		filepath = flag.Args()[index]
		if strings.HasPrefix(filepath, "-") {
			return filepath, false
		}
	} else {
		return "", false
	}
	l.Debug("Get path of file want be uploaded:", filepath)
	return filepath, true
}
