package downcase

//import "fmt"
func Downcase(str1 string) (string, error) {
	var str2 [100]byte
	for i := 0; i < len(str1); i++ {
		if str1[i] >= 'A' && str1[i] <= 'Z' {
			str2[i] = str1[i] + 32
		} else {
			str2[i] = str1[i]
		}
	}
	red := string(str2[:len(str1)])
	return red, nil
}

/*	func main () {
	str1 := "123 adfaswe"
	fmt.Println(downcase (str1))
}*/
