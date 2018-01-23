package easy

type parenStack []byte

func (ps *parenStack) Peek() byte {
	if len(*ps) == 0 {
		panic("can't peek at an empty stack")
	}
	return (*ps)[len(*ps)-1]
}

func (ps *parenStack) Pop() byte {
	if len(*ps) == 0 {
		panic("can't pop an empty stack")
	}
	res := (*ps)[len(*ps)-1]
	*ps = (*ps)[0 : len(*ps)-1]
	return res
}

func (ps *parenStack) Push(b byte) {
	*ps = append(*ps, b)
}

func (ps parenStack) Len() int {
	return len(ps)
}

var parenMap = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}

func isValid(s string) bool {
	var stack parenStack
	for _, b := range []byte(s) {
		var (
			closingFor byte
			isClosing  bool
		)
		switch b {
		case '{', '(', '[':
			stack.Push(b)
		case '}', ']', ')':
			closingFor = parenMap[b]
			isClosing = true
		default:
			panic("isValid does not accept non-parenthesis inputs")
		}
		if isClosing {
			if stack.Len() == 0 {
				return false // no opening parenthesis to check
			}
			if stack.Peek() == closingFor {
				stack.Pop()
			} else {
				return false // expecting a corresponding closing parenthesis
			}
		}
	}
	return stack.Len() == 0
}
