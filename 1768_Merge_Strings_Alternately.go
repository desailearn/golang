package main

import "fmt"

func main_() {

	fmt.Println(mergeAlternately("Hello", "Word"))
}

func mergeAlternately(str1 string, str2 string) string {

	m, n := len(str1), len(str2)
	ans := make([]byte, 0, m+n)
	for i := 0; i < m || i < n; i++ {
		if i < m {
			ans = append(ans, str1[i])
		}
		if i < n {
			ans = append(ans, str2[i])
		}
	}
	return string(ans)

}
