package main

import (
	"bufio"
	"fmt"
	"os"
)

func run() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("引数が足りません。")
	}

	src, dst := os.Args[1], os.Args[2]

	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("ファイルが開けませんでした。%s", src)
	}
	// TODO: 関数終了時にファイルを閉じる
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("ファイルを書き出せませんでした。%s", dst)
	}
	// TODO: 関数終了時にファイルを閉じる
	defer df.Close()

	scanner := bufio.NewScanner(sf)
	// TODO: sfから1行ずつ読み込み、"行数:"を前に付けてdfに書き出す。
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		df.WriteString(fmt.Sprintf("%d:%s\n", lineNumber, line))
	}

	// TODO: scannerから得られたエラーを返す
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ファイルの読み込みに失敗しました。%s", src)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
