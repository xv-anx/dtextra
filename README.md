# Data Extraction

<img src"https://user-images.githubusercontent.com/128390652/227145272-f13ab318-eb12-44b1-95f5-404fbe921e1f.png" width="100%"

## 使い方

1. cheatsフォルダを作成し、cheatsフォルダにcode.txtファイルを作成します。  
    または以下のコマンドで実行し作成できます  
```
$ go run main.go
```
2. code.txtの例は以下の通り
```
[Monster HP 0]
00DCDD08 E59213E8
00DCDD0C E1A09001
00DCDD10 E2691000
00DCDD14 E12FFF1E
!%
00904598 E1A00000
009045A0 E1A00000
0090A260 EB130EA8
!{This is note.}
%!
```
  [] は絶対必須で、[]内のチート名を出力します。  
  サブルーチン以外のコードは !% ~ %! で囲んでください。  
  また、%! の前に !{} で囲むとNoteを出力します。  
3. 最後にビルドをします。
### Windows
```
$ go build -o [ファイル名].exe
```
### Linux
```
$ go build -o [ファイル名]
```
