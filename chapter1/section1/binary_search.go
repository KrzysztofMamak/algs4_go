package main

func BinarySearch(key int, a []int) int {
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func BinarySearchRec(key int, a []int) int {
	return binarySearchRecInner(key, a, 0, len(a))
}

func binarySearchRecInner(key int, a []int, lo int, hi int) int {
	if lo <= hi {
		mid := lo + (hi-lo)/2

		if key < a[mid] {
			return binarySearchRecInner(key, a, lo, mid-1)
		} else if key > a[mid] {
			return binarySearchRecInner(key, a, mid+1, hi)
		} else {
			return mid
		}
	} else {
		return -1
	}
}
