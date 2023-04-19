# chatgpt-slack-bot-on-cloudrun

## 概要
ChatGPTと会話できるSlack BotをGCP上に構築する  
Botに対するMentionイベントをトリガーに起動、返信を返す。
Slackの [3秒ルール](https://api.slack.com/apis/connections/events-api#responding) に対応するために、2つのCloud Runコンテナ、Pub/Subを利用する

## 前準備
* [Slack API Tokenを作成](./docs/create-slack-bot.md)
* [OpenAI APIキーを作成](./docs/create-openAI-APIKey.md)

## デプロイ手順

infra配下の[README](./infra/README.md)を参照

## ローカル開発

application配下の[README](./application/README.md)を参照