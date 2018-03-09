package quene


type Quene []interface{}

func (q *Quene) Push(v interface{}){
	*q = append(*q,v)
}

func (q *Quene) pop() interface{}{
	head:= (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Quene) isEmpty() bool {
	return len(*q) == 0
}

func main() {
	
}
