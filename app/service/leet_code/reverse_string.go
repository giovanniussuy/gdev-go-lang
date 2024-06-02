package leetcode

func reverseString(s []byte) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func reverseStringBestMs(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseStringBestMemory(s []byte) {
	wordLength := len(s)
	for leftPointer := 0; leftPointer < wordLength/2; leftPointer++ {
		rightPointer := wordLength - leftPointer - 1
		s[leftPointer], s[rightPointer] = s[rightPointer], s[leftPointer]
	}
}
