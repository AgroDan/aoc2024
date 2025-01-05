package utils

import (
	"bufio"
	"fmt"
	"os"
)

func FlushScreen(s string) {
	// This will take whatever you type here and write it to the screen immediately.
	// It will also append \r so that it can be written to again so that you can write
	// updates to the screen as quickly as possible without scrolling through repeated
	// text
	writer := bufio.NewWriter(os.Stdout)
	// spaceBuf := int(math.Abs(float64(300 - len(s))))
	// expandedStr := fmt.Sprintf("%-*s", spaceBuf, s)
	fmt.Fprintf(writer, "%s\r", s)
	writer.Flush()
}
