package cmd

import (
	"fmt"
	"github.com/leancodebox/goose/fileopt"
	"github.com/leancodebox/goose/jsonopt"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "tool:yzw:json_tool",
		Short: "json处理",
		Run:   runYzwJsonTool,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func runYzwJsonTool(_ *cobra.Command, _ []string) {
	vJson := fileopt.StorageGet("vlist.json")
	vList := jsonopt.Decode[[]schoolInfo](vJson)
	var newVList []schoolInfo

	fileopt.PutContent("t.csv", "城市,博士点,招生单位,院系,专业,方向,学习方式,招生人数,考试方式,政治,英语,专业课1,专业科2,信息源\n")
	fmt.Println(len(vList))

	// Less
	// 排序原则
	// 完成现阶段目标的当置于中间位置
	// 不期望的当滞后
	// 离目标越接近的越靠前
	// 离目标同样接近的期望越高越高前
	sort.SliceStable(vList, func(i, j int) bool {
		return vList[i].City < vList[j].City
	})
	hMap := map[string]schoolInfo{}
	for _, vItem := range vList {
		if _, ok := hMap[vItem.Name]; ok {
			continue
		}
		hMap[vItem.Name] = vItem
		newVList = append(newVList, vItem)
		for _, vtItem := range vItem.SpecialityList {
			for _, sItem := range vtItem.Info {
				t := fmt.Sprintf("%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,%schoolInfo,研招网2022信息\n",
					replaceDot(vItem.City),      // 城市
					replaceDot(vtItem.IsDoctor), // 博士点
					replaceDot(sItem.Zsdw),      // 招生单位
					replaceDot(sItem.Yx),        // 院系
					replaceDot(sItem.Zy),        // 专业
					replaceDot(sItem.Fx),        // 方向
					replaceDot(sItem.StudyFunc), // 学习方式
					replaceDot(sItem.Number),    // 招生人数
					replaceDot(sItem.Kxfs),      // 考试方式
					replaceDot(sItem.Zz),        // 政治
					replaceDot(sItem.Wy),        // 英语
					replaceDot(sItem.B1),        // 专业课1
					replaceDot(sItem.B2),        // 专业科2
				)
				fileopt.AppendPutContent("t.csv", t)
			}
		}
	}
	vListString := jsonopt.Encode(newVList)
	fileopt.PutContent("final2.json", vListString)
}

func replaceDot(s string) string {
	return strings.NewReplacer("\n", "", ",", "_", "，", "_").Replace(s)
	//return strings.ReplaceAll(s, ",", "")
}
