# GO-ITDDD-05-REPOSITORY

zenn の記事「[Go でファクトリを実装（「入門ドメイン駆動設計」Chapter9）](https://zenn.dev/msksgm/articles/20220415-go-itddd-09-factory)」のサンプルコードです。

# 実行環境

- Go
  - 1.18
- docker compose

# 実行方法

## コンテナを起動・マイグレーション

コンテナの起動

```bash:コンテナの起動
> make up
docker compose up -d
# 完了までまつ
```

```bash:マイグレーション
> make run-migration
docker compose exec app bash db-migration.sh
1/u user (13.817ms)
```

## 実行

test-user 登録 1 回目

```bash:test-user 登録 1回目
> make run
docker compose exec app go run main.go
2022/04/14 21:36:19 successfully connected to database
2022/04/14 21:36:19 user name of test-user is successfully saved
```

test-user 登録 2 回目

```bash:test-user 登録 2回目
> make run
docker compose exec app go run main.go
2022/04/14 21:36:30 successfully connected to database
2022/04/14 21:36:30 userapplicationservice.Register err: user name of test-user is already exists.
exit status 1
make: *** [run] Error 1
```

# テスト

```bash
> make test
docker compose exec app go test ./...
?       github.com/msksgm/go-itddd-09-factory   [no test files]
ok      github.com/msksgm/go-itddd-09-factory/domain/model/user 0.003s
```
