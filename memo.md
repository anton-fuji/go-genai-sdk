# go-gemini-agent

Google Gen AI SDK for Go を使った、コンテキスト理解エージェント

## 概要

- ReAct ループベースのエージェント  
- LangChainGo ツール（電卓・HTTP など）  
- 簡易メモリ機能付き CLI / サーバ

## ディレクトリ構成
(↓予定)
```markdown
go-gemini-agent/
├── cmd/               # 実行可能ファイルごとにサブディレクトリ
│   └── agent/
│       └── main.go    # エントリーポイント（CLI / サーバ）
├── internal/          # アプリ固有ロジック（外部公開しない）
│   ├── agent/         # ReAct ループやメモリ実装
│   ├── tools/         # 電卓・HTTP など LangChainGo ツール
│   ├── llm/           # go-genai ラッパ・モデル選択
│   ├── config/        # viper 等で env / YAML 読み込み
│   └── server/        # gin / echo 等の HTTP ハンドラ
├── scripts/           # ローカル開発用スクリプト (make, bash, mage)
└── go.mod
…


無料枠：Studio キーは月 60 req/分＆40K トークン/日相当のライト枠。超えると HTTP 429。課金したい場合は Vertex AI に移行。

## agent.env or .env　ファイルを作成
```bash
export GOOGLE_GENAI_USE_VERTEXAI=false         # Studio キーなら false の
export GOOGLE_API_KEY="YOUR_ACTUAL_API_KEY" 
```

フラグ入れたかったら↓
```
go run ./cmd/agent -model gemini-2.0-flash -temp 0.7
```