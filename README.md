# notify-web-update

# よく使ううコマンド

```
```text
$ GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
$ zip deployment.zip bootstrap

// 関数作成
$ aws lambda create-function --function-name MyLambdaFunction --zip-file fileb://deployment.zip --handler main --runtime provided --role arn:aws:iam::...

// 関数更新
$ aws lambda update-function-code --function-name MyLambdaFunction --zip-file fileb://deployment.zip
```

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


