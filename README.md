# pi-home

自宅のRaspberry Piの管理用リポジトリ。

## 機能

- リバースプロキシ
- MYDNSへの定期リクエスト更新
- 月間写真アワード用機能の提供

## 管理方針

### TODO管理
やることは基本的にIssueとして管理する。

### branch管理
当面はmainを基本ブランチとし、mainブランチに対してPull Requestを立てる。
開発時のブランチ名はIssue番号を用いて`feature/#{Issue番号}`とする。