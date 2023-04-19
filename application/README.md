# Slack Botのローカル開発環境の構築
※コマンドの実行は `chatgpt-slackbot-on-cloudrun/application` ディレクトリ内で行う
## 1. 必要なツールのインストール

DockerとDocker Composeがインストールされていることを確認してください。インストールされていない場合は、公式ドキュメントに従ってインストールを行ってください。

## 2. サンプルファイルのリネーム

`docker-compose.yml.sample`ファイルを`docker-compose.yml`にリネームします。

## 3. APIキーの設定

`docker-compose.yml` 内の `SLACK_BOT_TOKEN` と `OPENAI_API_KEY` を、それぞれ取得済みのものに置き換えてください。

## 4. プロジェクトディレクトリの構成

applicationディレクトリには以下の3つのディレクトリが格納されています。
- gpt-app: Pub/Sub Pushを受け取り、ChatGPTの返答作成、Slackへメッセージを送信する
- pubsub-app: SlackEventを受け取り、Pub/SubへPush
- pubsub-emulator: Pub/Subエミュレータ

## 5. Docker Composeでコンテナを起動

プロジェクトディレクトリで以下のコマンドを実行して、Dockerコンテナを起動します。

```
docker-compose up --build
```

## 6. ngrokのインストールと起動

ngrokがインストールされていない場合は、[公式サイト](https://ngrok.com/download)からダウンロードし、インストールしてください。

インストールが完了したら、以下のコマンドを実行してngrokを起動し、pubsub-appのポート8081を公開します。
```
ngrok http 8081
```

表示されるForwardingのURL（例: https://xxxxxx.ngrok.io）をコピーしてください。

## 7. Slack Event Subscriptionsの設定

Slack APIの[Event Subscriptions](https://api.slack.com/apps)ページで、以下の設定を行います。

- Enable Eventsをオンにする
- Request URLに、手順6で取得したngrokのURLを入力する

設定が完了したら、Save Changesボタンをクリックして保存します。

## 8. Slack Botの動作確認

ChatGPT Botにメンションを飛ばし、正常に動作していることを確認してください。  


## 9. 開発が完了したら、Docker Composeでコンテナを停止

開発が完了したら、以下のコマンドを実行してDockerコンテナを停止します。
```
docker-compose down
```

これでSlack Botのローカル開発環境の構築が完了です。開発が進んで新しい機能を追加する際や、コードを変更した際には、手順5と8を繰り返して動作確認を行ってください。

ngrokを終了する際は、ngrokの起動ターミナルでCtrl+Cを押してプロセスを停止してください。