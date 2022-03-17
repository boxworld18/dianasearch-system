// @title		StoreItem
// @description	此函数会拆解上层接口提供的 item，并根据 get_config 的配置，将 item 中的数据截取并分类放入不同的数据库中。
// @auth		ryl				2022/3/17		12:00
// @param		data			*etree.Element	上层数据提供的 item 树
// @param		resourceName	string			特型卡类型
// @param		operation		string			对数据库的操作
// @param		docid			int				item 编号
// @return		err				error			错误值

package write_setting

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
)

func StoreItem(data *etree.Element, resourceName string, operation string, docid int) error {
	var itemSettings []ItemSetting
	var err error

	// 查找对应特型卡的配置
	itemSettings, err = GetConfig(resourceName)
	if err != nil {
		return err
	}

	// 根据配置信息写入数据库
	for _, itemSetting := range itemSettings {

		// 根据路径选取对应数据
		path := strings.Replace(itemSetting.itemPath, ".", "/", -1)

		for _, value := range data.FindElements(path) {
			// 数据写入摘要
			if itemSetting.dumpDigest {
				fmt.Println(value.Text())
				// Push(resourceName, docid)
			}

			// 数据写入倒排引擎
			if itemSetting.dumpInvertIdx {
				fmt.Println(value.Text())
				// Push(docid, value.Text())
			}

			// 数据写入词典
			if itemSetting.dumpDict {
				fmt.Println(value.Text())
				// Push(value.Text())
			}
		}

	}
	return nil
}
