package main

func isPalindrome(x int) bool {
	original := x
	reverse := 0
	
	for x > 0 {
		d := x % 10	
		reverse = reverse * 10 + d
		x = x / 10
	}
	return original == reverse
}

func main() {
	num := 10

	println(isPalindrome(num))
}
