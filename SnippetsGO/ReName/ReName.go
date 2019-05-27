package ReName

//把一个目录下面，所以文件包含的A名字改成B名字
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	strPath = `E:\game3_girl\trunk\client\GirlProject\Assets\GameResources\model\mob\m104_str` //位置
	srcName = "stoneman"                                                                       //要替换的名字
	newName = "m104_str"                                                                       //新的名字
)

//传入一个路径，原来的名字，你想要变的名字
func Start(path string, oldname string, tarname string) {
	strPath = path
	srcName = oldname
	newName = tarname
	filepath.Walk(strPath, fnwalk)
}

//对每个文件的操作
func fnwalk(path string, info os.FileInfo, err error) error {
	if strings.Contains(info.Name(), srcName) {
		dirPath, fileName := filepath.Split(path)
		newfileName := strings.Replace(fileName, srcName, newName, -1)
		newPath := dirPath + newfileName
		fmt.Println(path)
		fmt.Println(newPath)
		os.Rename(path, newPath)
	}
	return err
}
