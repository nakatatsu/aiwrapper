// cmd/proxy/main.go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/openai/openai-go"
)

func main() {
	// ──────────────────────────────
	// 0. フラグ解析
	// ──────────────────────────────
	logDir := flag.String("log", "logs", "ログ保存ディレクトリ（省略時は ./logs）")
	model := flag.String("model", "gpt-4o", "使用するモデル ID を指定")
	flag.Parse()

	// ──────────────────────────────
	// 1. 標準入力の読み取り
	// ──────────────────────────────
	promptBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "stdin 読み込み失敗:", err)
		os.Exit(1)
	}
	prompt := string(promptBytes)

	// ──────────────────────────────
	// 2. ログディレクトリの準備
	// ──────────────────────────────
	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS
	if err := os.MkdirAll(*logDir, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "ログディレクトリ作成失敗:", err)
		os.Exit(1)
	}

	// ──────────────────────────────
	// 3. パラメータ構築 & リクエスト JSON 保存
	// ──────────────────────────────
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: *model,
	}
	reqJSON, _ := json.MarshalIndent(params, "", "  ")
	if err := os.WriteFile(filepath.Join(*logDir, timestamp+"_request.log"), reqJSON, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "リクエストログ書込失敗:", err)
		os.Exit(1)
	}

	// ──────────────────────────────
	// 4. OpenAI API 呼び出し
	// ──────────────────────────────
	client := openai.NewClient()
	resp, err := client.Chat.Completions.New(context.TODO(), params)
	if err != nil {
		fmt.Fprintln(os.Stderr, "OpenAI API エラー:", err)
		os.Exit(1)
	}

	// ──────────────────────────────
	// 5. レスポンス JSON 保存
	// ──────────────────────────────
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	if err := os.WriteFile(filepath.Join(*logDir, timestamp+"_response.log"), respJSON, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "レスポンスログ書込失敗:", err)
		os.Exit(1)
	}

	// ──────────────────────────────
	// 6. 本文を標準出力へ
	// ──────────────────────────────
	fmt.Print(resp.Choices[0].Message.Content)
}
