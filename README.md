# 環境構築

ソースコードをダウンロード。

```bash
git clone https://github.com/junichi4250/monitoring-system.git
```

docker 立ち上げ

```bash
docker-compose up -d --build
```

# 実行方法

コンテナの中に入る

```bash
docker-compose exec golang sh
```

実行

```bash
go run *.go
```

# 実行結果

## 課題 1

```bash
故障サーバーip: 10.20.30.1/16
故障期間:
20201019133324 ~ 20201019133326
20201019133327 ~ 20201019133331
20201019133345 ~ 20201019133347
故障サーバーip: 10.20.30.2/16
故障期間:
20201019133325 ~ 20201019133328
```

# 内容補足

故障しているサーバーのみ実行結果に出力されるようにしています。

# ディレクトリ構成
```bash
└── monitoring-system
    ├── main.go.     main
    ├── model.go     serverのstructを管理
    ├── breakDown.go 故障サーバーと故障期間を出力している
    └── access.log   アクセスログ
```
