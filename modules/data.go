package modules

import "strings"

// データを均等に指定したサイズに割り当てる
func DataGroupe(slice []string, dataSize int) []string {
	// データを初期化
	var datas []string

	// i を引数sliceの長さまでデータサイズを増加
	for i := 0; i < len(slice); i += dataSize {
		
		// データサイズをendに保持
		end := i + dataSize

		// endがスライスの長さを超える場合
		if end > len(slice) {

			// emdをスライス長さに設定
			end = len(slice)
		}
		// i番目の要素からend-1番目の要素までを区切り文字", "で結合したスライスを作成しデータに追加
		datas = append(datas, strings.Join(slice[i:end], ", "))
	}
	// 結合し追加したデータ群を返す
	return datas
}