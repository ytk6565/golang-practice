package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var np = flag.Bool("n", false, "行番号を表示するか")

// ファイルを開く
func openFile(filepath string) *os.File {
	sf, err := os.Open(filepath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ファイルが見つかりませんでした。", err)
	}

	return sf
}

// ファイルを1行ずつ読み込む
func scanFile(sf *os.File) []string {
	var output []string
	scanner := bufio.NewScanner(sf)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました。", err)
	}

	return output
}

// cat コマンド
func mycat(filepaths []string, enableLineNumber bool) {
	if len(filepaths) == 0 {
		return
	}

	// 出力用の文字列の配列を生成
	var output []string

	for _, filepath := range filepaths {
		sf := openFile(filepath)
		if sf == nil {
			break
		}
		result := scanFile(sf)
		output = append(output, result...)
	}

	if len(output) == 0 {
		return
	}

	for index, text := range output {
		if enableLineNumber {
			fmt.Printf("%d: %s\n", index+1, text)
		} else {
			fmt.Printf("%s\n", text)
		}
	}
}

func main() {
	flag.Parse()
	mycat(flag.Args(), *np)
}
