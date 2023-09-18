package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"slices"
	"strings"
)

func init() {
	cmd := &cobra.Command{
		Use:   "json_2_md",
		Short: "",
		Run:   runJson2Md,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runJson2Md(cmd *cobra.Command, args []string) {
	jsonStr := `{
  "string_value": "Hello World",
  "number_value": 123,
  "boolean_value": true,
  "array_value": [
    "apple",
    2,
    false,
    {
      "nested_string": "This is a nested string",
      "nested_number": 456,
      "nested_boolean": false
    }
  ],
  "object_value": {
    "nested_string": "This is a nested string",
    "nested_number": 456,
    "nested_array": [
      "banana",
      null,
      true,
      {
        "nested_property": "Nested object in array"
      }
    ],
    "nested_object": {
      "nested_boolean": false,
      "nested_null": null
    }
  }
}`
	// param, _ := cmd.Flags().GetString("param")
	fmt.Println("success")

	stats := JSONStats{
		Keys:  make([]string, 0),
		Types: make(map[string][]string),
	}
	jsonStr = d

	err := parseJSON("", []byte(jsonStr), &stats)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("|key \t| value \t|")
	fmt.Println("|---| ---|")
	for _, valueType := range stats.Keys {
		fmt.Printf("|%s|%s|\n", valueType, strings.Join(stats.Types[valueType], ","))

	}
}

type JSONStats struct {
	Keys  []string
	Types map[string][]string
}

func parseJSON(prefix string, jsonData []byte, stats *JSONStats) error {
	var data interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}
	parseJSONObject(prefix, data, stats)
	return nil
}

func parseJSONObject(prefix string, jsonObj interface{}, stats *JSONStats) {
	switch obj := jsonObj.(type) {
	case map[string]interface{}:
		for key, value := range obj {
			if !slices.Contains(stats.Keys, prefix+key) {
				stats.Keys = append(stats.Keys, prefix+key)
			}
			switch value.(type) {
			case []interface{}:
				if !slices.Contains(stats.Types[prefix+key], "array") {
					stats.Types[prefix+key] = append(stats.Types[prefix+key], "array")
				}
				parseJSONArray(prefix+key+".[]", value.([]interface{}), stats)
			case map[string]interface{}:
				if !slices.Contains(stats.Types[prefix+key], "object") {
					stats.Types[prefix+key] = append(stats.Types[prefix+key], "object")
				}
				parseJSONObject(prefix+key+".", value.(map[string]interface{}), stats)
			default:
				nodeType := fmt.Sprintf("%T", value)
				if !slices.Contains(stats.Types[prefix+key], nodeType) {
					stats.Types[prefix+key] = append(stats.Types[prefix+key], fmt.Sprintf("%T", value))
				}
			}
		}
	}
}

func parseJSONArray(prefix string, jsonArray []interface{}, stats *JSONStats) {
	for _, item := range jsonArray {
		switch item.(type) {
		case []interface{}:
			if _, ok := stats.Types[prefix]; !ok {
				stats.Keys = append(stats.Keys, prefix)
				stats.Types[prefix] = append(stats.Types[prefix], "array")
			}
			parseJSONArray(prefix+".[]", item.([]interface{}), stats)
		case map[string]interface{}:
			parseJSONObject(prefix, item.(map[string]interface{}), stats)
		default:
			if _, ok := stats.Types[prefix]; !ok {
				stats.Keys = append(stats.Keys, prefix)
			}
			nodeType := fmt.Sprintf("%T", item)
			if !slices.Contains(stats.Types[prefix], nodeType) {
				stats.Types[prefix] = append(stats.Types[prefix], nodeType)
			}
		}
	}
}

// 测试-提前终止-tmp

var d = `{
    "applicantName": "霍思盾（huosidun）",
    "applicantDept": "商业智能中心-交易研发部-企业系统部",
    "cityContractNum": "12",
    "bpmSn": "21",
    "constractType": {
        "id": "1",
        "value": "流量对赌"
    },
    "contractCity": {
        "id": "2048",
        "value": "三明"
    },
    "isDirect": {
        "id": "1",
        "value": "是"
    },
    "agentEntityName": "xx",
    "ourEntityName": "xxx",
    "terminatedDate": "2023-09-13",
    "refundContractAmount": "0",
    "refundSecurityAmount": "0",
    "invoiceRefund": {
        "id": "1",
        "value": "是"
    },
    "reason": "12",
    "attachment": [
        {
            "name": "out2023_02_27_15_38_1991.png",
            "url": "http://testv3.58v5.cn/FHgHBazYxWFki/attachment/5f633b22171a4732bd2d9e57c09f21cc.png"
        }
    ],
    "processNameKey": "jacBgCGbvZ",
    "processNameValue": "代理终止协议",
    "p_number": "20230913114251695869",
    "processCurrentState": "1",
    "submitUserName": "huosidun",
    "submitUserDept": "201307121704152ae4f131",
    "submitUserPost": "201101101655507f728cae",
    "submitUserLocation": "LOCATION-3-129",
    "submitUserCity": "1",
    "submitUserCostCenterTopTwo": "50",
    "submitUserBudgetOrg": "913663770788651008",
    "submitUserEntryDate": "20220830",
    "businessGroupId": "3",
    "token": "xxx"
}`
