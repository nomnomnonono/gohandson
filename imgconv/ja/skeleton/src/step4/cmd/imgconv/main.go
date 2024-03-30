package main

import (
	"fmt"
	"image"

	// TODO: pngとjpegをデコードできるようにimportする。
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func convert(dst, src string) error {

	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("画像ファイルが開けませんでした。%s", src)
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした。%s", dst)
	}
	defer df.Close()

	// TODO: 入力ファイルから画像をメモリ上にデコードする。
	srcImage, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	// TODO: 拡張子によって保存する形式を変える。
	// ".png"の場合は、png形式で、".jpeg"と".jpg"の場合はjpeg形式で保存する。
	// 拡張子は大文字でも小文字でも動作するようにする。
	// なお、jpegは`jpeg.DefaultQuality`で保存する。
	// エラー処理も忘れないようにする。
	switch strings.ToLower(filepath.Ext(dst)) {
	case ".png":
		png.Encode(df, srcImage)
	case ".jpeg", ".jpg":
		jpeg.Encode(df, srcImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	default:
		return fmt.Errorf("未対応の拡張子です。%s", dst)
	}

	return nil
}

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("画像ファイルを指定してください。")
	}

	return convert(os.Args[2], os.Args[1])
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
