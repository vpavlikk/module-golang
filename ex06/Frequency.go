package letter

/*package main*/
import (
	//"fmt"
	"time"
)

var n = make(map[byte]int)

func Frequency(str string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}
	return m
}
func Frequency_chan(str string) {

	for i := 0; i < len(str); i++ {
		n[str[i]]++
	}

}
func ConcurrentFrequency(array []string) map[byte]int {

	for i := 0; i < len(array); i++ {
		go Frequency_chan(array[i])
		time.Sleep(1 * time.Millisecond)
	}
	//	time.Sleep(500 * time.Millisecond)
	return n
}

/*func main() {
	fmt.Println(Frequency("aswefawf"+"gtrgege"+"bbbswkefnk"))
	name:=[]string{"aswefawf","gtrgege","bbbswkefnk"}
	fmt.Println(ConcurrentFrequency(name))

}*/
