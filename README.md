##  .env　ファイルを作成
```bash
GOOGLE_GENAI_USE_VERTEXAI=false         # Studio キーなら false の
GOOGLE_API_KEY="YOUR_API_KEY" 
```
実行方法
```sh
go run cmd/main.go
```

フラグ入れたかったら↓
```
go run ./cmd/agent -model gemini-2.0-flash -temp 0.7
```