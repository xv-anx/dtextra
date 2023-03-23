# Data Extraction

![imege](https://user-images.githubusercontent.com/128390652/227150689-481f225c-9644-4b26-863c-9afaccb26d46.png)

## 使い方

1. cheatsフォルダを作成し、cheatsフォルダにcode.txtファイルを作成します。
    またはbuild.batを実行すると作成されます。  
```
$ go run main.go
```
2. code.txtの例は以下の通り
```
[Unlimited kinsect extract time v1.4]
00BF2FC0 E3510000
00BF2FC4 E3A01445
00BF2FC8 1580131C
00BF2FCC ED900AC7
00BF2FD0 E12FFF1E
!%
00A88D10 E590131C
00A88D14 EA05A8A9
{操虫棍エキス時間無限}
%!
```
  - [] は必須、[]内のチート名を関数名に変換します。  
  - サブルーチン以外のコードは !% ~ %! で囲んでください。  
  - %! の前に {} で囲むとNoteを出力します、ここはコメントアウトとして出力します。
  - 複数のコードでも問題なくビルドできますが必ず [] は必要です。   
  - 現在Gateway条件(D3, D2など)はサポートしておりません。  
3. 最後にbuild.batを起動しビルドします。    

## ビルド方法
### Windows
```
$ go build -o [ファイル名].exe
```
### Linux
```
$ go build -o [ファイル名]
```
