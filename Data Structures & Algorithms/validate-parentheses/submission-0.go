type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) push (v T){
    s.items = append(s.items,v)
}

func (s*Stack[T]) pop () (T, bool) {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    n := len(s.items)-1
    v := s.items[n]
    s.items = s.items[:n]
    return v,true
}

func (s*Stack[T]) peek() (T,bool) {
    var zero T
    if len(s.items) == 0{
        return zero, false
    }
    return s.items[len(s.items)-1],true
}

func (s*Stack[T]) len() int{
 return len(s.items)
}

func isValid(s string) bool {
    var st Stack[rune]
    pairs := map[rune]rune{')' : '(', '}' :  '{', ']' : '['}
    for _,ch := range s {
        switch ch {
        case '(','[','{' : 
            st.push(ch)
        case ']','}',')' :
            top,ok:= st.peek()
            if !ok || top!=pairs[ch]{
                return false
            }
            st.pop()
        }
    }
   return st.len() == 0
}
