package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// 옵션 구조체 정의
type Options struct {
	// ConfigFlags는 kubeconfig, context, namespace 등 쿠버네티스 관련 플래그들을 포함합니다
	configFlags *genericclioptions.ConfigFlags

	// IOStreams는 입출력 스트림을 관리합니다
	genericclioptions.IOStreams
}

// 새 Options 인스턴스를 생성하는 함수
func NewOptions(streams genericclioptions.IOStreams) *Options {
	// 입력으로 만들어준 iostreams를 사용하여 Options 인스턴스를 생성
	return &Options{
		configFlags: genericclioptions.NewConfigFlags(true), // true는 기본 플래그들을 모두 활성화한다는 의미
		IOStreams:   streams,
	}
}

func main() {
	// 입출력 스트림 설정
	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	// 옵션 인스턴스 생성
	options := NewOptions(streams)

	// 루트 명령어 생성
	// 이 옵션을 cobra에 넘기는 구조
	cmd := &cobra.Command{
		Use:   "k8s-config-info",
		Short: "쿠버네티스 설정 정보를 출력합니다",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 쿠버네티스 설정 로더 생성
			configLoader := options.configFlags.ToRawKubeConfigLoader()

			// 현재 네임스페이스 정보 가져오기
			namespace, overridden, err := configLoader.Namespace()
			if err != nil {
				return err
			}

			// 결과 출력
			fmt.Fprintf(options.Out, "현재 네임스페이스: %s\n", namespace)
			if overridden {
				fmt.Fprintf(options.Out, "네임스페이스가 플래그로 재정의되었습니다.\n")
			} else {
				fmt.Fprintf(options.Out, "네임스페이스는 kubeconfig에서 가져왔습니다.\n")
			}

			// kubeconfig 파일 경로 출력
			kubeconfigPath := ""
			if options.configFlags.KubeConfig != nil {
				kubeconfigPath = *options.configFlags.KubeConfig
			}
			if kubeconfigPath == "" {
				fmt.Fprintf(options.Out, "기본 kubeconfig 경로를 사용합니다.\n")
			} else {
				fmt.Fprintf(options.Out, "사용 중인 kubeconfig 경로: %s\n", kubeconfigPath)
			}

			// 현재 컨텍스트 정보 출력
			currentContext := ""
			if options.configFlags.Context != nil {
				currentContext = *options.configFlags.Context
			}
			if currentContext == "" {
				fmt.Fprintf(options.Out, "현재 컨텍스트는 kubeconfig에서 가져옵니다.\n")
			} else {
				fmt.Fprintf(options.Out, "현재 컨텍스트: %s\n", currentContext)
			}

			return nil
		},
	}

	// 쿠버네티스 관련 플래그들을 명령어에 추가
	options.configFlags.AddFlags(cmd.Flags())

	// 명령어 실행
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(options.ErrOut, "오류: %v\n", err)
		os.Exit(1)
	}
}
