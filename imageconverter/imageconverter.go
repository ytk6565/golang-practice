package imageconverter

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

/*
	- ディレクトリを指定する
	- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
	- ディレクトリ以下は再帰的に処理する
	- 変換前と変換後の画像形式を指定できる（オプション）
*/

/*
	TODO: 型の絞り込みをしたい
*/

const _JPG = "jpg"
const _JPEG = "jpeg"
const _PNG = "png"
const _GIF = "gif"

var before = flag.String("before", "jpg", "変換前の拡張子")
var after = flag.String("after", "png", "変換後の拡張子")

// デコードする関数
type tdecode = func(r io.Reader) (image.Image, error)

// エンコードする関数
type tencode = func(w io.Writer, img image.Image) error

/*
	TODO: switch で記述したい
*/

// デコードする関数を取得する
func getDecode(ext *string) (tdecode, error) {
	if *ext == _JPG || *ext == _JPEG {
		return jpeg.Decode, nil
	}
	if *ext == _PNG {
		return png.Decode, nil
	}
	if *ext == _GIF {
		return gif.Decode, nil
	}
	return nil, fmt.Errorf("拡張子は「%+v」「%+v」「%+v」「%+v」のいずれかを利用してください。", _JPG, _JPEG, _PNG, _GIF)
}

/*
	TODO: switch で記述したい
*/

// エンコードする関数を取得する
func getEncode(ext *string) (tencode, error) {
	if *ext == _JPG || *ext == _JPEG {
		return func(w io.Writer, img image.Image) error {
			return jpeg.Encode(w, img, nil)
		}, nil
	}
	if *ext == _PNG {
		return png.Encode, nil
	}
	if *ext == _GIF {
		return func(w io.Writer, img image.Image) error {
			return gif.Encode(w, img, nil)
		}, nil
	}
	return nil, fmt.Errorf("拡張子は「%+v」「%+v」「%+v」「%+v」のいずれかを利用してください。", _JPG, _JPEG, _PNG, _GIF)
}

/*
	TODO: switch で記述したい
*/

// 対応している拡張子であるか検証する
func makeExt(ext *string) (string, error) {
	if *ext == _JPG || *ext == _JPEG || *ext == _PNG || *ext == _GIF {
		return "." + *ext, nil
	} else {
		return "", fmt.Errorf("拡張子は「%+v」「%+v」「%+v」「%+v」のいずれかを利用してください。", _JPG, _JPEG, _PNG, _GIF)
	}
}

// 対象のファイルパスの拡張子を変換する関数のファクトリー
func convertImageFactory(decode tdecode, encode tencode) func(w io.Writer, r io.Reader) error {
	return func(w io.Writer, r io.Reader) error {
		img, err := decode(r)
		if err != nil {
			return err
		}
		return encode(w, img)
	}
}

// 画像を指定の拡張子に変換する
func Convert() {
	var err error

	flag.Parse()
	dir := flag.Arg(0)

	beforeExt, err := makeExt(before)
	if err != nil {
		fmt.Printf("変換前の拡張子に問題があります。%v\n", err)
		return
	}

	afterExt, err := makeExt(after)
	if err != nil {
		fmt.Printf("変換後の拡張子に問題があります。%v\n", err)
		return
	}

	decode, err := getDecode(before)
	if err != nil {
		fmt.Printf("変換前の拡張子に問題があります。%v\n", err)
		return
	}

	encode, err := getEncode(after)
	if err != nil {
		fmt.Printf("変換後の拡張子に問題があります。%v\n", err)
		return
	}

	convertImage := convertImageFactory(decode, encode)

	err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// 対象の情報を出力する
		if info.IsDir() {
			fmt.Printf("ディレクトリ「%+v」の検証をしています。\n", path)
		} else {
			fmt.Printf("ファイル「%+v」の検証をしています。\n", path)
		}

		// 指定された拡張子のファイル以外は何もしない
		if info.IsDir() || filepath.Ext(path) != beforeExt {
			fmt.Println("何もせずに終了します。")
			return nil
		}

		// 変換前のファイルを開く
		reader, err := os.Open(path)
		if err != nil {
			fmt.Printf("ファイル「%q」を開けませんでした。%v\n", path, err)
			return err
		}
		defer reader.Close()

		// 変換後のファイルを作成する
		writer, err := os.Create(strings.Replace(path, beforeExt, afterExt, 1))
		if err != nil {
			fmt.Printf("ファイル「%q」の作成に失敗しました。%v\n", path, err)
			return err
		}
		defer writer.Close()

		// 変換を行う
		err = convertImage(writer, reader)
		if err != nil {
			fmt.Println("変換に失敗しました。")
		} else {
			fmt.Println("変換が完了しました。")
		}

		return nil
	})

	if err != nil {
		fmt.Printf("指定されたディレクトリの検証中にエラーが発生しました。%v\n", err)
		return
	}
}
