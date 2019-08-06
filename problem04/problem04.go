package problem04

func ReplaceSpace(str []byte, length int) {
	count := 0
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			count++
		}
	}
	newlength := length + count*2

	for l, nl := length-1, newlength-1; l >= 0 && nl >= 0; {
		if str[l] == ' ' {
			str[nl] = '0'
			nl--
			str[nl] = '2'
			nl--
			str[nl] = '%'
			nl--
			l--
		} else {
			str[nl] = str[l]
			l--
			nl--
		}
	}
}
