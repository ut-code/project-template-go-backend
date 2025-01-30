# すごいアプリ

これは、すごいアプリのテンプレートです。

## 使い方

このアプリを使うには、 <https://utcode.net> にアクセスします。

## デプロイ

1. 任意のホスティングサービス (例: Render, Heroku など) と CDN (例: Netlify,
   Cloudflare Pages など) を用意します。
2. ホスティングサービスの Go ランタイムを選択し、バックエンドをデプロイします。
   (サポートされてない場合や VPS を使う場合は Docker なり CD
   でビルドするなりで対応してください)
3. CDN にフロントエンドをデプロイします。
4. フロントエンド (CDN 側で与えられるURL) にアクセスします。

各必要な設定は以下のとおりです。

```yml
バックエンド:
  環境変数:
    CORS_ALLOW_ORIGINS: (フロントエンドの URL)
    DSN: (データベースの DSN)

フロントエンド:
  ルートディレクトリ: frontend
  ビルドコマンド: npm run build
  配信ディレクトリ: dist
  環境変数:
    VITE_API_ENDPOINT: (バックエンドの URL)
```

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

`npm run frontend:dev`
を実行してください。そうすると、フロントエンドの開発環境が起動します。

`npm run backend:dev`
を実行してください。そうすると、バックエンドの開発環境が起動します。

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
