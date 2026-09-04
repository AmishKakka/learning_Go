package main
import "fmt"

func make2DArray(dx int, dy int) [][]uint8 {
	// creating 2D arrays using the 'make' keyword
	// this creates 'dy' columns
	ss := make([][]uint8, dy)
	for i:=0; i < dy; i++ {
		// create an array of length 'dx'
		row := make([]uint8, dx)
		for j := range dx {
			row[j] = uint8((i + j) / 2)
		}
		ss[i] = row
	}
	return ss
}

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

	// slicing array; slice is a dynamically-sized, flexible view into the elements of an array
	slice := nums2[1:3]
	fmt.Println(slice)

	// appending elements in order
	nums3 := append(slice, 4, 5)
	fmt.Println("nums3: ", nums3)

	// making a 2D array
	fmt.Println(make2DArray(3, 4))
}