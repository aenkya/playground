//go:build !codeanalysis
// +build !codeanalysis

package practice

import (
	"enkya.org/playground/practice/algo"
)

func Practice() {
	// golang.GoPractice()
	// designpatterns.DesignPatterns()
	ap := algo.Practice()
	ap.SetChallengeName("ReverseArray")
	ap.RunAlgo()

	// if _, err := io.Copy(w, os.Stdout); err != nil {
	// 	log.Fatal(err)
	// }

	// go func() {
	// 	buf := make([]byte, 1024*1024)
	// 	for {
	// 		n, err := r.Read(buf)
	// 		if err != nil {
	// 			if err == io.EOF {
	// 				break
	// 			}

	// 			fmt.Println("Error reading from pipe: ", err)
	// 			break
	// 		}
	// 		fmt.Printf("%s", buf[:n])

	// 	}
	// }()
}
