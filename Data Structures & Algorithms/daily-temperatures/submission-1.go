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

func dailyTemperatures(temperatures []int) []int {
    var st Stack[int] // stores indexes
    re := make([]int,len(temperatures))

    for idx,temp := range temperatures {
        for {
            top_idx,ok:= st.peek()
            if !ok || temperatures[top_idx] >= temp {
                break
            }
            st.pop()
            re[top_idx] = idx - top_idx
        }
        st.push(idx)
    }

    return re

}
