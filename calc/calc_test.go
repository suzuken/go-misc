package calc

import "testing"

// TestSumは加算のテストをします。
// 引数には *testing.T を渡します。
// 必ず Test から始まる名前にしましょう。
// すると、go testでの実行対象になります。
func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		// t.Fatalはテストが失敗したことを返すAPIです。
		// 多くのGoのテストコードでは条件分岐とt.Fatalを組み合わせて書くことになります。
		// t.Fatal以外にも、t.Fatalfもあります。
		// これらはテスト失敗時のエラーメッセージを加工するものです。
		// 別の例で詳しく見ていきます。
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}
