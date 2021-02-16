package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// return ind < 0 if subdir not found
func lastDirSearch(path string) (ind int, err error) {
	subFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return -1, err
	}
	dirInd := -1
	for ind, dir := range subFiles {
		if dir.IsDir() {
			dirInd = ind
		}
	}
	return  dirInd, nil
}

func dirRead(out io.Writer, path string, printFiles bool, depth int, depthSign []bool) (returnSign []bool, err error) {
	subFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	tabSymb := ""
	for i := 0; i < depth; i++ {
		if depthSign[i] == true {
			tabSymb += "│\t"
		} else {
			tabSymb += "\t"
		}

	}
	tabSymb += "├───"
	lastDir := -1
	if printFiles == false {
		lastDir, err = lastDirSearch(path)
		if err != nil {
			return nil, err
		}

	}
	for ind, subFile := range subFiles {
		if ind + 1 == len(subFiles) || (lastDir == ind && printFiles == false){
			tabSymb = strings.ReplaceAll(tabSymb, "├───", "└───")
		}
		if (!subFile.IsDir() && printFiles) || subFile.IsDir() {
			fmt.Fprintf(out, fmt.Sprintf("%s %s\n", tabSymb, subFile.Name()))
		}
		if subFile.IsDir() {
			if ind + 1 == len(subFiles) ||  (lastDir == ind && printFiles == false) {
				if len(depthSign) < depth + 1 {
					depthSign = append(depthSign, false)
				} else {
					depthSign[depth] = false
				}
			} else {
				if len(depthSign) < depth + 1 {
					depthSign = append(depthSign, true)
				} else {
					depthSign[depth] = true
				}
			}
			depthSign, err = dirRead(out, filepath.Join(path, subFile.Name()), printFiles, depth + 1, depthSign)
			if err != nil {
				return nil, err
			}
		}
	}
	//depthSign[depth] = false
	return depthSign, nil
}

func dirTree(out *bytes.Buffer, path string, printFiles bool) error {
	var depthSign []bool
	depthSign = append(depthSign, true)
	_, err := dirRead(out, path, printFiles, 0, depthSign)
	return err
}

func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	//path := "testdata"
	//printFiles := false

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(out)
	return
}
