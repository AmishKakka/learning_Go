package main
import "fmt"

func arraysDemo() {
	//  ... means length is inferred when compiled
	// nums := [5]int{} array of length 5, all zeros
	var nums1 = [...]int{1, 2, 3}
	nums2 := [4]int{1,2,3,4}
	fmt.Println("nums1: ", nums1, "nums2: ", nums2)
	fmt.Println("nums1 length: ", len(nums1))

	// accessing elements
	nums2[1] = 12
	fmt.Println("nums2: ", nums2)

	// slicing array
	slice := nums2[1:3]
	fmt.Println(slice)

	// appending elements in order
	nums3 := append(slice, 4, 5)
	fmt.Println("nums3: ", nums3)
}