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
func carFleet(target int, position []int, speed []int) int {
	pos_time_map := make(map[int]float64)
	for idx := range position {
		time :=( float64(target) - float64(position[idx]) ) / float64(speed[idx])
		pos_time_map[position[idx]] = time
	}
	fmt.Println(pos_time_map)
	var st Stack[int]
	sort.Ints(position)
	for i:= len(position) - 1; i >=0 ; i-- {
		top_el,ok := st.peek()
		if !ok || pos_time_map[top_el] < pos_time_map[position[i]] {
			st.push(position[i])
		}
	}
	return st.len()
}
