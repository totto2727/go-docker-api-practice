# 注意
プログラムでDocker Socketを直接操作します。イメージやコンテナなどの作成・削除も行うため、必ず本番環境以外で実験してください。
## 実行環境
- 新規の仮想環境
- `docker:dind`イメージのようなDinD環境にGoのRuntimeを追加する
- VSCodeの[Dev container features](https://code.visualstudio.com/docs/remote/containers#_dev-container-features-preview)によるDinD環境（ベースはGoのテンプレート）

# go-docker-api-practice
`Go`の`Docker SDK`の練習用リポジトリです。以下の操作は実演しています。
- image 
    - pull
- container
    - ps
    - create
    - start
    - stop
    - remove