// @Title  unpackfile
// @Description  Read XML files in folder and unpack them
// @Author  于沛楠
// @Update  2022/3/16
package unpackfile

import (
	"testing"
)

// test for function: getFiles
func TestGetXMLFiles(t *testing.T) {
	filesPath, err := getXMLFiles("./car")
	if err != nil {
		t.Error(err)
	}
	if len(filesPath) != 1 || filesPath[0] != "./car/car_test.xml" { // wrong return value
		t.Error("getFile函数读取错误")
	}

	_, err = getXMLFiles("./cars")
	if err != nil && err.Error() != "open ./cars: no such file or directory" { //wrong error content
		t.Error(err)
	}
}
