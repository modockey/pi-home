# PHOTO-AWARDとは

家族内で「月間写真アワード」を開催することになった。これをシステムで実現するためこのプロジェクトを立ち上げた。

##　月間写真アワードの概要

月に一度、ユーザーは写真を登録する。これをノミネートという。

### ノミネートのルール

ノミネートする写真は以下のいずれかを満たす。
- 自分で撮影した写真
- 自分が写っている写真

なお、ノミネートする写真はユーザーあたり3枚までとする。

## 受賞

ノミネートされた写真のうち一枚を「月間写真アワード」とする。決定方法は未決。
「特別賞」などを自由に付与できる。

# 実現方法

## アルバム

アルバムを作成できるものとする。
「月間写真アワード」はこれを毎月作成することで実現する。

### ノミネート

アルバムに写真を登録することを「ノミネート」とする。

### タグ機能

写真にはタグをつけることができ、タグの内容はアルバムを横断して編集できる。

現在想定しているタグは以下

- 最優秀賞
- 特別賞

# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type aware lint rules:

- Configure the top-level `parserOptions` property like this:

```js
export default {
  // other rules...
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
    project: ['./tsconfig.json', './tsconfig.node.json'],
    tsconfigRootDir: __dirname,
  },
}
```

- Replace `plugin:@typescript-eslint/recommended` to `plugin:@typescript-eslint/recommended-type-checked` or `plugin:@typescript-eslint/strict-type-checked`
- Optionally add `plugin:@typescript-eslint/stylistic-type-checked`
- Install [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) and add `plugin:react/recommended` & `plugin:react/jsx-runtime` to the `extends` list
