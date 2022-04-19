// @title	TestSetData
// @description	此函数的用途为检查 SetData 函数的正确性
// @auth	ryl		2022/4/14	10:30
// @param	t		*testing.T	testing 用参数

package io

import (
	"dianasdog/setup"
	"io/ioutil"
	"testing"
)

func TestSetData(t *testing.T) {

	// 得到此文件的绝对路径
	abspath, _ := setup.GetAbsPath()

	// 取得测试文件
	filepath := abspath + "data/testcase/testcase_car.xml"

	// 读入文件
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Error(err)
	}

	err = SetData("testcase", "testcase_car.xml", file)
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

}