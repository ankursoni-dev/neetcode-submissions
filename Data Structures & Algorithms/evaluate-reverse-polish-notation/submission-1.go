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

func evalRPN(tokens []string) int {
    var st Stack[string]
    ops_map:= map[string]bool{"+":true,"-":true,"/":true,"*":true}

    for _,token := range tokens {
        if ops_map[token] {
            str2,_:= st.pop() 
            str1,_:= st.pop()
            num2,_:= strconv.Atoi(str2)
            num1,_:= strconv.Atoi(str1)
            res:= strconv.Itoa(compute(num1,num2,token))
            st.push(res)
        }else {
            st.push(token)
        }
    }
    ansStr,_ := st.peek()
    ansNum,_ := strconv.Atoi(ansStr)
    return ansNum
}

func compute(num1 int, num2 int, op string)int{
    switch op {
        case "+" :
            return num1+num2
        case "-" :
            return num1-num2
        case "*" :
            return num1*num2
        case "/" :
            return num1/num2
    }
    return 0
}
