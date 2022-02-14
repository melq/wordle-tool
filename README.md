# Wordle-tool

このプログラムは、Wordleを解いている途中の回答情報をもとに答えとしてあり得る単語の一覧を出力する。

## 使用方法

回答情報はコマンドライン引数として入力する。

例として、

![img.png](img.png)

のような回答の状況の場合は、
`go run . ra?ise !cou?nt`
のように、色付きの文字の前に記号を挿入し引数として入力する。すると結果として以下のような出力が得られる。

```text
> go run . ra?ise !cou?nt
suggest:
cynic cinch 
```