# Slack Bot作成手順書

この手順書では、Slack Botの作成、Tokenの取得、およびEvent SubscriptionsのURL設定を説明します。

## 1. Slack Botの作成

1. [Slack APIページ](https://api.slack.com/apps)にアクセスし、右上の**Create New App**ボタンをクリックします。

2. **App Name**に、Botの名前を入力し、**Development Slack Workspace**でBotを開発するためのワークスペースを選択します。

3. **Create App**ボタンをクリックして、アプリを作成します。

## 2. Tokenの取得

1. 作成したアプリの**Basic Information**ページに移動します。

2. 左側のメニューから**OAuth & Permissions**を選択します。

3. **Bot Token Scopes**の下の**Add an OAuth Scope**をクリックし、必要なスコープを追加します。必要なスコープは以下になります。
* `app_mentions:read`
* `channels:history`
* `chat:write`
* `groups:history`
* `im:history`
* `mpim:history`

4. **Install App**の下にある**Install to Workspace**ボタンをクリックして、Botをワークスペースにインストールします。

5. インストールが完了すると、**Bot User OAuth Token**が表示されます。このTokenをコピーし、安全な場所に保管してください。
