package card

import "fmt"

// FormatDuration converts milliseconds to "min:sec" format (e.g., 225000 -> "03:45")
func FormatDuration(ms int) string {
	totalSeconds := ms / 1000
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
