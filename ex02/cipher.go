/*package main
import "fmt"
*/
package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

/////////////Caesar////////////////
type Caesar struct {
}

func (point Caesar) Encode(str1 string) string {
	var str2 [100]byte
	l := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] >= 'A' && str1[i] <= 'Z' {
			str2[l] = str1[i] + 32
			l++
		} else if str1[i] >= 'a' && str1[i] <= 'z' {
			str2[l] = str1[i]
			l++
		}
	}
	str1 = string(str2[:l])

	for j := 0; j < len(str1); j++ {
		if str1[j] == 'x' || str1[j] == 'y' || str1[j] == 'z' {
			str2[j] = str1[j] - 23
		} else {
			str2[j] = str1[j] + 3
		}
	}
	return string(str2[:len(str1)])
}

func (point Caesar) Decode(str1 string) string {
	var str2 [100]byte
	for j := 0; j < len(str1); j++ {
		if str1[j] == 'a' || str1[j] == 'b' || str1[j] == 'c' {
			str2[j] = str1[j] + 23
		} else {
			str2[j] = str1[j] - 3
		}
	}
	return string(str2[:len(str1)])
}

func NewCaesar() Cipher {
	return Caesar{}
}

//////////////////////Shift/////////////////////////////
type Shift struct {
	n byte
}

func (point Shift) Encode(str1 string) string {
	var str2 [100]byte
	l := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] >= 'A' && str1[i] <= 'Z' { //переделываем большие буквы в маленькие
			str2[l] = str1[i] + 32
			l++
		} else if str1[i] >= 'a' && str1[i] <= 'z' {
			str2[l] = str1[i]
			l++
		}
	}
	str1 = string(str2[:l])

	for j := 0; j < len(str1); j++ {
		if str1[j]+point.n < 'a' {
			str2[j] = str1[j] + point.n + 26
		} else if str1[j]+point.n > 'z' {
			str2[j] = str1[j] + point.n - 26
		} else {
			str2[j] = str1[j] + point.n
		}
	}
	return string(str2[:len(str1)])
}

func (point Shift) Decode(str1 string) string {
	var str2 [100]byte
	for j := 0; j < len(str1); j++ {
		if str1[j]-point.n > 'z' {
			str2[j] = str1[j] - point.n - 26
		} else if str1[j]-point.n < 'a' {
			str2[j] = str1[j] - point.n + 26
		} else {
			str2[j] = str1[j] - point.n
		}
	}
	return string(str2[:len(str1)])
}

func NewShift(n int) Cipher {
	if n >= -25 && n <= 25 && n != 0 {
		res := byte(n)
		return Shift{res}
	} else {
		return nil
	}
}

//////////////////////ViginereShift////////////////////////////
type Vigenere struct {
	n string
}

func Len_key(str1 string, n string) string {
	var str2 [100]byte
	j := 0
	for i := 0; i < len(str1); i++ {

		if j > len(n)-1 {
			j = 0 //counter по строке ключа
		}
		str2[i] = n[j]
		j++
	}
	ret := string(str2[:len(str1)])
	return ret
}

func (point Vigenere) Encode(str1 string) string {
	var str2 [100]byte
	if len(point.n) < len(str1) {
		point.n = Len_key(str1, point.n)
	}
	l := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] >= 'A' && str1[i] <= 'Z' {
			str2[l] = str1[i] + 32
			l++
		} else if str1[i] >= 'a' && str1[i] <= 'z' {
			str2[l] = str1[i]
			l++
		}
	}
	str1 = string(str2[:l])

	for j := 0; j < len(str1); j++ {
		if str1[j]+(point.n[j]-'a') > 'z' {
			str2[j] = str1[j] + (point.n[j] - 'a') - 26
		} else {
			str2[j] = str1[j] + (point.n[j] - 'a')
		}
	}
	return string(str2[:len(str1)])
}

func (point Vigenere) Decode(str1 string) string {
	var str2 [100]byte
	if len(point.n) < len(str1) {
		point.n = Len_key(str1, point.n)
	}
	for j := 0; j < len(str1); j++ {
		if str1[j]-(point.n[j]-'a') < 'a' {
			str2[j] = str1[j] - (point.n[j] - 'a') + 26
		} else {
			str2[j] = str1[j] - (point.n[j] - 'a')
		}
	}
	return string(str2[:len(str1)])
}

func right_symbols(str1 string) bool {
	if str1 == "" {
		return false
	}
	count := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] == 'a' {
			count++
		}
		if !(str1[i] >= 'a' && str1[i] <= 'z') {
			return false
		}
	}
	if count == len(str1) {
		return false
	}
	return true
}

func NewVigenere(n string) Cipher {
	if right_symbols(n) == true {
		return Vigenere{n}
	} else {
		return nil
	}
}

///////////////////////////////////////////////////
/*func main () {

	var c =NewCaesar()
	encode_caesar := c.Encode("go go go phers")
	fmt.Println(encode_caesar)
	fmt.Println(c.Decode(encode_caesar))
	var v =NewShift(3)
	encode_shift := v.Encode("go go gophers")
	fmt.Println(encode_shift)
	fmt.Println(v.Decode(encode_shift))
	var d =NewVigenere("abc")
	encode_vigenere := d.Encode("Vla  d")
	fmt.Println(encode_vigenere)
	fmt.Println(d.Decode(encode_vigenere))

}*/
