package stringGo

import (
	"fmt"
	"strings"
)

func CustomerStrings(val, sqval int) {
	var string strings.Builder

	str := fmt.Sprintf("The original val is %d and Squre is %d", val, sqval)

	string.WriteString(str)

	fmt.Println(string.String())
}
