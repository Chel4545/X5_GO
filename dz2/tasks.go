package homework

// PairInput соответствует объекту nums/target из условия контеста.
type PairInput struct {
	Nums   []int
	Target int
}

// longestPrefix возвращает самый длинный общий префикс строк.
func longestPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}

	prefix := []rune(words[0])

	for _, word := range words[1:] {
		letters := []rune(word)

		i := 0
		for i < len(prefix) && i < len(letters) && prefix[i] == letters[i] {
			i++
		}

		prefix = prefix[:i]

		if len(prefix) == 0 {
			return ""
		}
	}

	return string(prefix)
}

// pairIndices возвращает индексы двух разных элементов по возрастанию.
func pairIndices(input PairInput) [2]int {
	// TODO: найти единственную пару с суммой input.Target.
	var m = make(map[int]int)

	for i, num := range input.Nums {
		if _, ok := m[input.Target-num]; ok {
			return [2]int{m[input.Target-num], i}
		}
		m[num] = i
	}

	return [2]int{0, 0}
}

// hasDuplicate сообщает, встречается ли какое-либо значение повторно.
func hasDuplicate(nums []int) bool {
	// TODO: проверить наличие повторов.
	var m = make(map[int]struct{})

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}

	return false
}

// compactSorted создаёт независимый слайс уникальных значений.
func compactSorted(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	result := make([]int, 0, len(nums))
	result = append(result, nums[0])
	for _, num := range nums[1:] {
		if num != result[len(result)-1] {
			result = append(result, num)
		}
	}
	return result
}

// balancedBrackets проверяет пары и порядок закрытия скобок.
func balancedBrackets(text string) bool {
	if len(text)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(text)/2)
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != text[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}
