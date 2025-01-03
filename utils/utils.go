package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HandleLongInput(text *string) {
	reader := bufio.NewReader(os.Stdin)
	dataInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error brok duar:", err)
		return
	}
	*text = strings.TrimSpace(dataInput)
}
