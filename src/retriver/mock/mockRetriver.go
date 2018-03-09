package mock

import "fmt"

type Retriver struct {
	Contents string
}

func (r Retriver) String() string {
	fmt.Println("Retriver: {Contents = %s};",r.Contents)
	return r.Contents
}



func (r Retriver) Post(url string, form map[string]string) string {
	panic("implement me")
}



func (r Retriver) Get(url string) string{
	return r.Contents
}

