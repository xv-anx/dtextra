# Data Extraction

![imege](https://user-images.githubusercontent.com/128390652/227150689-481f225c-9644-4b26-863c-9afaccb26d46.png)

## 使い方

1. cheatsフォルダを作成し、cheatsフォルダにcode.txtファイルを作成します。  
    またはbuilf.batを実行すると作成されます。  

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
  - [] は絶対必須で、[]内のチート名を出力します。  
  - サブルーチン以外のコードは !% ~ %! で囲んでください。  
  - また、%! の前に !{} で囲むとNoteを出力します。  
  - コードは複数あっても問題ありません。  
  
3. 最後にbuild.batを起動しビルドします。  
