package CsDir

import (
	"fmt"
	"testing"
)

func TestCsDir(t *testing.T) {
	SrcDir := "E:/golang/csfioletext/go"
	//ShareDir := "E:/golang/csfioletext/Share"
	ShareDir := "./"
	_, dirs, _ := WalkDirFile(SrcDir, "")
	// fmt.Println(files, err)
	//ListFileFunc(files)
	//ListFileFunc(dirs)
	DirHead := GetDirHead(SrcDir)
	err := Mkdirs(dirs, ShareDir, DirHead)
	if err != nil {
		t.Error("error!\n")
	} else {
		fmt.Println("nice day!")
	}
}
