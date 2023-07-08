# notify-web-update

# プロジェクト構成

```text
    /modules
        /line
            notifier.go // LINEに通知を送るロジック
        /web
            scraper.go // ウェブスクレイピングのロジック
    /tests
        // ユニットテストや統合テストのコード
    main.go // AWS Lambdaから呼び出されるメインのハンドラー関数
    Dockerfile
    .gitignore
    README.md
    go.mod
    go.sum
```


