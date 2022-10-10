package emojinumbers

import (
	"fmt"
	"strings"
)

func ToEmoji(number int) string {
	if number == 10 {
		return "🔟"
	}
	if number == 100 {
		return "💯"
	}
	str := fmt.Sprintf("%d", number)
	str = strings.ReplaceAll(str, "0", "0️⃣")
	str = strings.ReplaceAll(str, "1", "1️⃣")
	str = strings.ReplaceAll(str, "2", "2️⃣")
	str = strings.ReplaceAll(str, "3", "3️⃣")
	str = strings.ReplaceAll(str, "4", "4️⃣")
	str = strings.ReplaceAll(str, "5", "5️⃣")
	str = strings.ReplaceAll(str, "6", "6️⃣")
	str = strings.ReplaceAll(str, "7", "7️⃣")
	str = strings.ReplaceAll(str, "8", "8️⃣")
	str = strings.ReplaceAll(str, "9", "9️⃣")
	return str
}
