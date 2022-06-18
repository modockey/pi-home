# home-management

## やりたいこと

- 家に取り付けたセンサーのデータを収集する
- 収集したデータを可視化し、自宅内LAN+外部からアクセスできるようにする

## 管理方針

### TODO管理
やることは基本的にIssueとして管理する。

### branch管理
当面はmainを基本ブランチとし、mainブランチに対してPull Requestを立てる。
開発時のブランチ名はIssue番号を用いて`feature/#{Issue番号}`とする。