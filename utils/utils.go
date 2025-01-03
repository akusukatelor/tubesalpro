package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HandleLongInput(text *string) {
	fmt.Scanln()
	reader := bufio.NewReader(os.Stdin)
	dataInput, _ := reader.ReadString('\n')
	*text = strings.TrimSpace(dataInput)
}
