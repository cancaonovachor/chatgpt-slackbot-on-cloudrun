# GCP環境へのTerraform実行手順
※コマンドの実行は `chatgpt-slackbot-on-cloudrun/infra` ディレクトリ内で行う

## 1. 必要なツールのインストール

- [Terraform](https://www.terraform.io/downloads.html)をインストールしてください。
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)をインストールしてください。

## 2. 認証情報の設定

以下のコマンドを実行して、GCPへのアクセスに必要な認証情報を設定します。

```
gcloud auth login
gcloud auth application-default login
```
※ 必ず両方のコマンドを実行してください


## 3. サンプルファイルのリネーム

`vars.tfvars.sample`ファイルを`vars.tfvars`にリネームします。

## 4. 設定値の入力

`vars.tfvars`ファイルを開き、以下の設定値を自前のものに置き換えます。

- project_id
- openai_api_key
- slack_token

## 5. Terraformの初期化

プロジェクトディレクトリで以下のコマンドを実行し、Terraformを初期化します。

```
terraform init
```


## 6. Terraformの実行

以下のコマンドを実行して、TerraformでGCPリソースを作成します。

```
terraform apply -var-file="vars.tfvars"
```


## 7. Cloud RunのURLの取得

Terraform実行後、ターミナルに表示されるCloud RunのURLをコピーしてください。
```
Outputs:

pubsubapp_cloudrun_url = "https://xxxxxxxxxx.a.run.app"
```

## 8. Slack Event Subscriptionsの設定

Slack APIの[Event Subscriptions](https://api.slack.com/apps)ページで、以下の設定を行います。

- Enable Eventsをオンにする
- Request URLに、手順7で取得したCloud RunのURLを入力する

設定が完了したら、Save Changesボタンをクリックして保存します。


## 9. デプロイ済みのSlack Botの動作確認

ChatGPT Botにメンションを飛ばし、正常に動作していることを確認してください。  



## APPENDIX. リソースの削除

リソースを削除する必要がある場合は、以下のコマンドを実行してTerraformで作成したリソースを削除します。
```
terraform destroy -var-file="vars.tfvars"
```
