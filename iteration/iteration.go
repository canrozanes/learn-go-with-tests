package iteration

//Repeat repeats a string
func Repeat(s string, count int) string {
	repeat := ""
	for i := 0; i < count; i++ {
		repeat = repeat + s
	}
	return repeat
}

// func main() {
// 	Repeat("a", 5)
// }
