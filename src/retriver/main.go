package main

import (
	"fmt"
	"retriver/mock"
	"retriver/real"
	"time"
)

type Retriver interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,form map[string]string) string
}

func download(r *Retriver,url string) string{
	return (*r).Get(url)
}

func post(poster Poster,url string,args map[string]string){
	poster.Post(url,args)
}

type RetriverPoster interface {
	Retriver
	Poster
	Connect(host string)
}

func session(s RetriverPoster){
}

func main(){
	var r Retriver = &real.Retriver{
		UserAgent:"Mozilla 5.0",
		TimeOut:time.Minute,
	}
	fmt.Printf("%T %v \n",r,r)

	switch v := r.(type) {
	case mock.Retriver:
		fmt.Println("Contents:",v.Contents)
	case *real.Retriver:
		fmt.Println("UserAgent:",v.UserAgent)
	default:
		panic("unsupport type...")
	}



	//fmt.Println(download(r,"https://www.baidu.com"))
	//fmt.Printf("%T %v \n",r,r)
}