# react-flow

## 概要

DB に登録されているデータを API で取得し、react-flow で表示します

## 初回起動

### コンテナ起動

`server`ディレクト配下に`.env`ファイルを作成し、`.env.sample`の中身をコピーして貼り付けます。

以下のコマンドを実行し、コンテナを起動します。

```
$ docker-compose up -d
```

### server のコンテナ内での作業

以下のコマンドを実行します。

```
$ go mod tidy
```

サーバーを起動します。

```
$ go run main.go
```

`server/sql`ディレクトリ配下の`sample-data.sql`ファイルの SQL を実行する

### client のコンテナ内での作業

以下のコマンドを実行します。

```
$ npm install
```

サーバーを起動します。

```
$ npm run dev
```

### db のコンテナ内での作業

以下のコマンドを実行し、MySQL に接続します。

```
$ mysql -u root -ppassword
```

`sql/sample-data.sql`ファイルに記載の SQL を以下の通り実行します。

```
mysql> INSERT INTO edges (id, source_id, target_id) VALUES (1, 1, 2), (2, 2, 3), (3, 3, 4);
```

```
mysql> INSERT INTO nodes (id, type, x, y, label) VALUES (1, "input", 0, 0, "start"), (2, null, 0, 100, "middle"), (3, "end", 0, 200, "end");
```

## 2 回目以降のコンテナ起動

`docker-compose.yml`内の`command: go run main.go`および`command: npm run dev`のコメントアウトを解除します。

以下のコマンドを実行し、コンテナを起動します。

```
$ docker-compose up -d
```

## コンテナ停止

以下のコマンドを実行します。

```
$ docker-compose down
```
