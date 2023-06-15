package containsDuplicate

func containsDuplicate(nums []int) bool {
  counter := make(map[int]int)
  for _, n := range nums {
    counter[n] += 1
    if counter[n] >= 2 {
      return true
    }
  }

  return false
}