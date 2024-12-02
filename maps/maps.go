package maps

import "fmt"

func main() {
	var websites = map[string]string{"Google": "www.google.com", "Facebook": "www.fb.com"}
	websites["aws"] = "www.aws.com"

	delete(websites, "Google")

	fmt.Println(websites)
}