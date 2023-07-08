# notify-web-update

# プロジェクト構成

```text
    /src
        /lambda
            handler.go // AWS Lambdaから呼び出されるメインのハンドラー関数
            scraper.go // ウェブスクレイピングのロジック
            notifier.go // LINEに通知を送るロジック
        main.go // ローカル開発環境で実行するためのエントリーポイント
    /tests
        // ユニットテストや統合テストのコード
    Dockerfile
    .gitignore
    README.md
    go.mod
    go.sum
```


