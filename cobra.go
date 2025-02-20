package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// 루트 커맨드 생성
	rootCmd := &cobra.Command{
		Use:   "mycli",
		Short: "A simple CLI using Cobra",
	}

	// `get` 명령 추가
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve resources",
	}

	// `get pods` 명령 추가
	getPodsCmd := &cobra.Command{
		Use:   "pods",
		Short: "Get list of pods",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing all pods")
		},
	}

	// 명령어 계층 구조 설정
	getCmd.AddCommand(getPodsCmd) // `get` 아래 `pods` 추가
	rootCmd.AddCommand(getCmd)    // 루트에 `get` 추가

	// 실행
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
