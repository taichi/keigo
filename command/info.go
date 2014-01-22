package command

import (
	"github.com/spf13/cobra"
	. "github.com/taichi/keigo/core"
	//"strings"
	"fmt"
)

var info = &cobra.Command{
	Use:   "info",
	Short: "Print the information of Keiko-chan",
	Long: `get information from Keiko-chan that contains 
farmware version, unit id,
and some maintenance informations.`,
	Run: newRunFn(func(cmd *cobra.Command, args []string) {
		results := execute(VERN, UTID, RDCD, RDCN, RDMN, RDPD, RDSN)
		labels := []string{
			// TODO NLS
			"Version",         // "ファームウェアバージョン",
			"UNIT ID",         // "ユニットＩＤ",
			"Effective Date",  // "保守契約期限",
			"Contract Number", // "保守契約番号",
			"Model Name",      // "モデル名",
			"Production",      // "製造年月",
			"Serial Number",   // "シリアル番号",
		}
		for i, v := range results {
			// 以下のフォーマットだと文字数基準でパディングされる為、2byte文字が混ざると綺麗に揃わない
			fmt.Printf("%-20s: %s\n", labels[i], v)
			// 以下のコードは、ラベル部分で1byteと2byteが混ざってると上手く動作しない
			//cmd.Printf("%s%s: %s\n", labels[i], strings.Repeat(" ", (12-len(labels[i])/3)*2), v)
		}
	}),
}
