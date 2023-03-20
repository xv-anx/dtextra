package main

import (
	"dtextra/modules"
	"strings"

	"github.com/fatih/color"
)

func main() {
	code := `
00DCDCD4 E92D4006
00DCDCD8 E2800FA4
00DCDCDC E1A02000
00DCDCE0 E3A01006
00DCDCE4 E5D00003
00DCDCE8 E1A00001
00DCDCEC E1A06000
00DCDCF0 E5C26003
00DCDCF4 E8BD8006
!%
00B54048 E1C109B6
00B540EC EB09E6F8
00B79EAC EB094F88
%!`

	// 文字列のコードを改行して分割した結果をスライス型で返す
	codeArrays := strings.Split(code, "\n") /*
		[ 00DCDCD4 E92D4006 00DCDCD8 E2800FA4 00DCDCDC E1A02000 00DCDCE0 E3A01006 00DCDCE4 E5D00003 00DCDCE8 E1A00001 00DCDCEC E1A06000 00DCDCF0 E5C26003 00DCDCF4 E8BD8006 !%
		00B54048 E1C109B6 00B540EC EB09E6F8 00B79EAC EB094F88 %!] */

	var romAddress []string
	var romData []string
	var judge bool

	var startAddress string
	var values []string

	// スライス型のコードをrangeで値を順番に取り出す、ループ回数は持たない
	for _, codeArray := range codeArrays {
		/*  00DCDCD4 E92D4006
		00DCDCD8 E2800FA4
		00DCDCDC E1A02000
		00DCDCE0 E3A01006
		00DCDCE4 E5D00003
		00DCDCE8 E1A00001
		00DCDCEC E1A06000
		00DCDCF0 E5C26003
		00DCDCF4 E8BD8006
		!%
		00B54048 E1C109B6
		00B540EC EB09E6F8
		00B79EAC EB094F88
		%! */

		// 取り出したスライスの値に "!%" がある場合は、%!まで処理をする
		if codeArray == "!%" {
			judge = true
			continue
		}
		//　%! までの処理
		if judge && codeArray != "%!" {
			// スペースで分割された各部分を抽出し、アドレスとデータを取得
			code := strings.Fields(codeArray)
			/* [00B54048 E1C109B6][00B540EC EB09E6F8][00B79EAC EB094F88]

			アドレス		データ
			romDatas[0]  romDatas[1] */

			// 各行が2つの部分に分割される場合のみ、アドレスとデータを抽出する
			if len(code) == 2 {
				// [00B54048 E1C109B6] アドレスとデータ

				// アドレスの接頭辞を0xに置換
				address := "0x" + code[0]
				// 0x00B540480x00B540EC0x00B79EAC

				// データの接頭辞を0xに置換
				data := "0x" + code[1]
				// 0xE1C109B60xEB09E6F80xEB094F88

				// それぞれアドレスとデータに置換した要素を追加、このとき末尾に追加される
				romAddress = append(romAddress, address)
				// 0x00B540480x00B540EC0x00B79EAC

				romData = append(romData, data)
				// E1C109B60xEB09E6F80xEB094F88
			}
			// 判定なしの場合、!% が見つからない場合
		} else if !judge {
			code := strings.Fields(codeArray)
			if len(code) == 2 {

				// 最初の要素から0xを削除
				address := strings.TrimPrefix(code[0], "0x")

				// アドレス桁数が8未満の場合
				if len(address) < 8 {

					// アドレス先頭に0を追加
					address = "0" + address
					// 00DCDCD4
				}

				// 開始アドレスに何もない場合
				if startAddress == "" {

					// 開始アドレス先頭に0xを追加
					startAddress = "0x" + address
					// 0x00DCDCD4
				}
				// 値の接頭辞を0xに置換
				value := "0x" + code[1]
				// 値に置換した要素を追加
				values = append(values, "0x"+value)
			}

		}
		// %!が見つかった場合は
		if codeArray == "%!" {
			// 判定なし
			judge = false
		}
	}
	c1 := color.New(color.FgRed)
	c1.Add(color.Bold)

	c2 := color.New(color.FgGreen)
	c2.Add(color.Bold)

	c3 := color.New(color.FgYellow)
	c3.Add(color.Bold)

	c1.Printf("Offset: %s\n\n", startAddress)
	c2.Printf("Values: \n%s\n\n", strings.Join(modules.DataGroupe(values, 4), ", \n"))
	for i := range romAddress {
		c3.Printf("Process::Write32(%s, %s);\n", romAddress[i], romData[i])
	}
}
