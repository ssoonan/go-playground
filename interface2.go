package main

import "fmt"

type AppsV1Interface interface {
	Deploy()
}

type AppsV1beta2Interface interface {
	Deploy()
}

type Client struct{}

func (c Client) Deploy() {
	fmt.Println("Deploying...")
}

func GetClientV1() AppsV1Interface {
	return Client{} // 정상
}

func GetClientV1beta2() AppsV1beta2Interface {
	return Client{} // 정상
}

// 하지만 다른 인터페이스 타입에 할당할 수 없음
func InvalidAssignment() {
	var v1 AppsV1Interface = GetClientV1beta2() // 컴파일 오류
	v1.Deploy()
}

func main() {
	client := GetClientV1()
	client.Deploy()

	InvalidAssignment()
}
