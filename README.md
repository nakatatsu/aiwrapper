# aiwrapper

ChatGPT API を呼び出し、リクエストとレスポンスをログとして保存するCLIツール。

## 特長

* 標準入力に渡されたプロンプトをそのまま API に送信します。
* `-log` フラグで指定したディレクトリ（既定値 `./logs`）に、タイムスタンプ付きでリクエスト・レスポンスを保存します。
* `-model` フラグで使用するモデルを選択できます（既定値 `gpt-4o`）。
* レスポンス本文は標準出力へそのまま出力します。


## 使い方

```bash
echo "HTTP/3 をわかりやすく説明してください" | aiwrapper -model gpt-4o -log mylogs
```

### フラグ

| フラグ      | 既定値      | 説明            |
| -------- | -------- | ------------- |
| `-log`   | `logs`   | ログを保存するディレクトリ |
| `-model` | `gpt-4o` | 使用するモデル ID    |

### 必要な環境変数

| 変数               | 概要              |
| ---------------- | --------------- |
| `OPENAI_API_KEY` | OpenAI の API キー |

## ログファイル

実行ごとに次の 2 ファイルがログとして生成されます。

| ファイル例                         | 内容                                   |
| ----------------------------- | ------------------------------------ | 
| `20250514094530_request.log`  | リクエストを保存 |                   |
| `20250514094530_response.log` | API から返された JSON レスポンスを保存             | 

