package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// コマンドライン引数をチェック
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [path to folder]")
		return
	}

	// 引数からフォルダパスを取得
	folderPath := os.Args[1]

	// 指定されたフォルダ内のファイルを取得
	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// フォルダ内の全てのファイルに対してループ
	for _, entry := range files {
		if !entry.IsDir() {
			filename := entry.Name()

			// ファイルがJPGであるかチェック
			if filepath.Ext(filename) == ".jpg" {
				// ファイル名から拡張子を除外
				nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))

				// Tesseractコマンドを構築
				outFilename := fmt.Sprintf("%s/%s_out.txt", folderPath, nameWithoutExt)
				cmd := exec.Command("tesseract", fmt.Sprintf("%s/%s", folderPath, filename), outFilename)

				// コマンドを実行
				if err := cmd.Run(); err != nil {
					fmt.Println("Error running tesseract:", err)
					continue
				}
				fmt.Println("Processed:", filename)
			}
		}
	}
}
