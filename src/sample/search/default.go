package search

type defaultMathcer struct {

}

func init() {
	var matcher defaultMathcer
	Register("default",matcher)

}



func (m defaultMathcer) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil,nil
}
