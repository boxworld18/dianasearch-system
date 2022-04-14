// @title	SetData
// @description	此函数的用途为，根据数据 "类型"，在配置文件中找出对应的 "写入行为"，并反馈相关数据到数据处理函数中。
// @auth	ryl			2022/4/13		16:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	filename	string			文件名
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package io

import (
	"dianasdog/setup"
	"io/ioutil"
)

func SetData(resource string, filename string, content []byte) error {

	// 得到此文件的绝对路径
	abspath, _ := setup.GetAbsPath()

	// 查找对应类型的文件路径
	filepath := abspath + "data/" + resource + "/" + filename

	// 写入配置
	err := ioutil.WriteFile(filepath, content, 0644)

	// 无论正确与否都返回 err 的内容
	return err
}
