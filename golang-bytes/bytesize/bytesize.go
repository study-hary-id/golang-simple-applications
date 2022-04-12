package bytesize

import "fmt"

// This constants provide default calculation of byte sizes.
const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
)

// Elegant constants with iota type.
//const(
//	KB ByteSize = 1<<( 10 * (iota + 1))
//	MB
//	GB
//	TB
//	PB
//)

// ByteSize is a custom type of float64, used to calculate bytes.
type ByteSize float64

// String returns formatted byte size with the unit prefix.
func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return "Very Big"
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.1fB", b)
}
