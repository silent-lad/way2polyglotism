package utils

import (
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)


func ClearTerminal() {
	machineRuntime := runtime.GOOS

	switch machineRuntime {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") 
        cmd.Stdout = os.Stdout
        cmd.Run()
	case "linux":
		cmd := exec.Command("clear") 
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}

func GetHangmanAscii(level int) string {
	x := level < 7 || level >= 0
	if x {
		return hangmanAscii[level]
	} else {
		return "Wrong level info"
	}
}

func GetWord() string {
	return wordList[rand.Intn(len(wordList))]
}


var wordList = []string{ "Each", "word", "is", "accompanied", "by", "four", "antonyms"}

var hangmanAscii = [...]string{`
+---+
|   |
	|
	|
	|
	|
=========`, `
+---+
|   |
O   |
	|
	|
	|
=========`, `
+---+
|   |
O   |
|   |
	|
	|
=========`, `
+---+
 |   |
 O   |
/|   |
	 |
	 |
=========`, `
+---+
 |   |
 O   |
/|\  |
     |
     |
=========`, `
+---+
 |   |
 O   |
/|\  |
/    |
     |
=========`, `
+---+
 |   |
 O   |
/|\  |
/ \  |
     |
=========`}