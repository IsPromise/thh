package arms

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLine(filePath string, action func(item string)) {

	f, errF := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if errF != nil {
		fmt.Print(errF)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		action(line)
	}
}
