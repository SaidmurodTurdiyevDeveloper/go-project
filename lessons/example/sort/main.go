package main

import "fmt"

func main() {
	nums := []int{-7087, 12694, -19352, -7660, 12052, -11316, -352, 18321, 15,
		19967, 6331, -1289, 6540, -10454, -19309, -10193, 15074, 8926, 510, -11044,
		5156, 10397, 18477, 7011, -16822, -8281, -13675, -14616, 11103, -3338, 1597,
		-10313, 15808, 16242, 11438, 19029, -3969, -5137, 6955, -11426, 11283, -9429,
		15512, 5109, -16951, 9741, 7024, 7374, -1429, -3822, 16494, 11183, 13581,
		-624, -8692, 9033, 13253, -6781, -8030, -12564, -6495, -8075, 6375, -7638,
		-8894, -15126, 6196, 7565, -8335, 17298, 15965, -10424, -11440, 115, -15303,
		-8185, -14136, -4704, 5078, 19498, -8636, 19068, 1950, 1980, -19147, 1634,
		-634, 5704, -16161, -19512, 18818, 1598, -13237, -4924, 8819, -17508, -6143,
		-922, -1262, 5676, -7286, 19153, 17772, 12016, 7017, -3883, 10757, -5870,
		-4128, 17396, 17387, 14927, 11939, 19301, 13760, 16166, 15332, 10975, -5162,
		-19353, 9945, -295, -13588, -1908, -13244, -348, -12064, 10565, 913, -16781,
		16214, -1846, 7727, 5283, 10492, -2839, 15925, -7995, 12903, -9995, -9204,
		-15275, -9845, -3452, 19892, -1347, -3371, 11224, -14823, -12338, -19412,
		-516, -14527, 6209, 11369, 12343, 18372, 11136, -8844, 15963, 12194, 12863,
		352, -3911, 932, 19855, -1760, 1582, 3047, 19}

	nums2 := []int{-7087, 12694, -19352, -7660, 12052, -11316, -352, 18321, 15,
		19967, 6331, -1289, 6540, -10454, -19309, -10193, 15074, 8926, 510, -11044,
		5156, 10397, 18477, 7011, -16822, -8281, -13675, -14616, 11103, -3338, 1597,
		-10313, 15808, 16242, 11438, 19029, -3969, -5137, 6955, -11426, 11283, -9429,
		15512, 5109, -16951, 9741, 7024, 7374, -1429, -3822, 16494, 11183, 13581,
		-624, -8692, 9033, 13253, -6781, -8030, -12564, -6495, -8075, 6375, -7638,
		-8894, -15126, 6196, 7565, -8335, 17298, 15965, -10424, -11440, 115, -15303,
		-8185, -14136, -4704, 5078, 19498, -8636, 19068, 1950, 1980, -19147, 1634,
		-634, 5704, -16161, -19512, 18818, 1598, -13237, -4924, 8819, -17508, -6143,
		-922, -1262, 5676, -7286, 19153, 17772, 12016, 7017, -3883, 10757, -5870,
		-4128, 17396, 17387, 14927, 11939, 19301, 13760, 16166, 15332, 10975, -5162,
		-19353, 9945, -295, -13588, -1908, -13244, -348, -12064, 10565, 913, -16781,
		16214, -1846, 7727, 5283, 10492, -2839, 15925, -7995, 12903, -9995, -9204,
		-15275, -9845, -3452, 19892, -1347, -3371, 11224, -14823, -12338, -19412,
		-516, -14527, 6209, 11369, 12343, 18372, 11136, -8844, 15963, 12194, 12863,
		352, -3911, 932, 19855, -1760, 1582, 3047, 19}

	sortArrayFirstAnswer(nums)
	sortArraySecondAnswer(nums2)

	var wrongCount = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			wrongCount++
		}
	}
	fmt.Println("Wrong count", wrongCount)
}

func maxMin(firstIndex int, lastIndex int, nums []int) (int, int) {
	maxIndex := lastIndex
	minIndex := firstIndex
	for i := firstIndex; i <= lastIndex; i++ {
		if nums[i] <= nums[minIndex] {
			minIndex = i
		}
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	return minIndex, maxIndex
}

func min(nums []int, endindex int) int {
	minIndex := 0
	for i := 0; i < endindex; i++ {
		if nums[i] < nums[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func sortArrayFirstAnswer(nums []int) []int {
	var startIndex = 0
	var endIndex = len(nums)
	if endIndex < 2 {
		return nums
	}
	if endIndex%2 == 1 {
		repleseData(&nums[0], &nums[min(nums, endIndex)])
		startIndex++
	}
	for startIndex < endIndex {
		min, max := maxMin(startIndex, endIndex-1, nums)
		repleseData(&nums[startIndex], &nums[min])
		if !(min == endIndex-1 && max == startIndex) {
			repleseData(&nums[endIndex-1], &nums[max])
		}
		startIndex++
		endIndex--
	}
	return nums
}

func sortArraySecondAnswer(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				repleseData(&nums[i], &nums[j])
			}
		}
	}
	return nums
}
func repleseData(firstItem *int, secondItem *int) {
	if firstItem == secondItem {
		return
	}
	*firstItem += *secondItem
	*secondItem = *firstItem - *secondItem
	*firstItem = *firstItem - *secondItem
}
