package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// etcd 클라이언트 생성
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd 서버 주소
		DialTimeout: 5 * time.Second,            // 연결 타임아웃 설정
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer client.Close()

	fmt.Println("Watching for changes on /my-key...")

	// Watch 채널 생성
	watchChan := client.Watch(context.Background(), "/my-key")

	// Watch 이벤트 감지
	for watchResp := range watchChan {
		for _, ev := range watchResp.Events {
			fmt.Printf("Type: %s | Key: %s | Value: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
