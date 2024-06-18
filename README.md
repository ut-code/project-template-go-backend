# すごいアプリ

これは、すごいアプリのテンプレートです。

## 使い方

このテンプレートをつかうには、
1. git clone します。
2. .git を消します。
3. git init します。
4. GitHub でリポジトリを作り、 git remote add origin [URL] とします。
6. README などを開発するアプリのものに書き換えます。
7. 開発します。
8. リリースします。
9. goto 7;

## 技術スタック

Frontend:
- React
- TypeScript
- Vite

Backend:
- Go
- Echo
- GORM

## 開発

### 要件

- Node.js >= v20
- Go >= v1.22

### 環境構築

このリポジトリをクローンしてから、プロジェクトのルートディレクトリに移動してください。

`npm run setup` を実行してください。初期設定が行われます。

`npm run frontend:dev` を実行してください。そうすると、フロントエンドの開発環境が起動します。

`npm run backend:dev` を実行してください。そうすると、バックエンドの開発環境が起動します。

http://localhost:5173/ でアプリケーションにアクセスできます。

### コミット前

コミット前には、以下のコマンドを実行して、コードスタイルと型のチェックを行ってください。

```sh
npm run lint && npm run type-check
```

## LICENSE

このテンプレートは MIT ライセンスのもとで公開されています。

## コントリビューション

Issue や PR はいつでもお待ちしております。
