/*
 * @Description: 处理测试文件分割
 * @Version: 1.0
 * @Author: zoetu 1753095374@qq.com
 * @Date: 2022-05-25 09:05:00
 * @LastEditTime: 2022-05-27 14:43:43
 */

package FileSplit

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

/**
 * @Description:按照size切割文件
 * @param {string} fileToBeChunked
 * @param {float64} size
 * @return {*}
 * @author: tanyan <1753095374@qq.com>
 */
func Split_xMB(fileToBeChunked string, size float64) {

	// 输入分割文件名
	file, err := os.Open(fileToBeChunked)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	//文件信息包括文件大小
	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()
	// fmt.Println("File Size = ", (fileSize / (1 << 20)), "MB")

	const fileChunk = 50 * (1 << 20) // 1 MB
	// TODO: 1<<20也就是1*2^20=1MB

	//计算需要切分的总块数
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / (float64(fileChunk))))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(1); i <= totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64((i-1)*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		// write to disk
		fileName := "./data/test/url_data_" + strconv.FormatUint(i, 10) + ".txt"
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		fmt.Println("Piece No", i, "Split to : ", fileName)

		// break //去掉则生成多个XMB文件
	}
}

/**
 * @Description: 按照行数分割文件
 * @param {string} fileToBeChunked
 * @param {float64} size
 * @return {*}
 * @author: tanyan <1753095374@qq.com>
 */
func Split_xLINE(fileToBeChunked string, lineChoose int) {

	// 输入分割文件名
	file, err := os.Open(fileToBeChunked)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	// 文件总行数
	lines := GetFileLines(fileToBeChunked)
	if lines > lineChoose {
		ReadFilexLine(fileToBeChunked, lineChoose)
	}
}

/**
 * @Description: 获取文件总行数
 * @param {string} fileName
 * @return {*}
 * @author: tanyan <1753095374@qq.com>
 */
func GetFileLines(fileName string) (lines int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fd := bufio.NewReader(file)
	lines = 0
	for {
		_, err := fd.ReadString('\n')
		if err != nil {
			break
		}
		lines++

	}
	fmt.Println("lines = ", lines)
	return lines
}

/**
 * @Description: 读取文件前n行
 * @param {string} fileName
 * @param {int} lineChoose
 * @return {*}
 * @author: tanyan <1753095374@qq.com>
 */
func ReadFilexLine(fileName string, lineChoose int) {
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// 读文件
	fd := bufio.NewReader(file)
	var dic []string
	dic = make([]string, lineChoose)
	line := 0

	for line < lineChoose { //超过选定行数，直接结束

		str, err := fd.ReadString('\n')
		dic[line] = str
		if err != nil {
			break
		}
		line++
	}
	// for i, _ := range dic {
	// 	fmt.Println("line", i, " = ", dic[i])
	// }

	//将[]string写入文件

	// 创建新文件
	NfileName := "./data/test/url_data_" + strconv.FormatInt(int64(lineChoose), 10) + "line.txt"
	f, error := os.Create(NfileName)

	if error != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 写入新文件

	// ioutil.WriteFile(NfileName, dic, os.ModeAppend)
	for i, _ := range dic {
		// fmt.Println("line", i, " = ", dic[i])
		_, err := f.WriteString(dic[i])
		if err != nil {
			fmt.Println("Write file err =", err)
			return
		}
	}

	fmt.Println(lineChoose, "lines Split to : ", NfileName)
}
