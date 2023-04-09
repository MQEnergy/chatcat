package chelp

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hashicorp/go-uuid"
	"github.com/shirou/gopsutil/host"
	"gorm.io/gorm/schema"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// GenerateBaseSnowId 生成雪花算法ID
func GenerateBaseSnowId(num int, n *snowflake.Node) string {
	if n == nil {
		localIp, err := GetLocalIpToInt()
		if err != nil {
			return ""
		}
		node, err := snowflake.NewNode(int64(localIp) % 1023)
		n = node
	}
	id := n.Generate()
	switch num {
	case 2:
		return id.Base2()
	case 32:
		return id.Base32()
	case 36:
		return id.Base36()
	case 58:
		return id.Base58()
	case 64:
		return id.Base64()
	default:
		return gconv.String(id.Int64())
	}
}

// GenerateUuid 生成随机字符串
func GenerateUuid(size int) string {
	str, err := uuid.GenerateUUID()
	if err != nil {
		return ""
	}
	return gstr.SubStr(str, 0, size)
}

// GeneratePasswordHash 生成密码hash值
func GeneratePasswordHash(password string, salt string) string {
	md5 := md5.New()
	io.WriteString(md5, password)
	str := fmt.Sprintf("%x", md5.Sum(nil))
	s := sha256.New()
	io.WriteString(s, password+salt)
	str = fmt.Sprintf("%x", s.Sum(nil))
	return str
}

// IsPathExist 判断所给路径文件/文件夹是否存在
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// MakeMultiDir 调用os.MkdirAll递归创建文件夹
func MakeMultiDir(filePath string) error {
	if !IsPathExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// MakeFileOrPath
// @Description: 创建文件/文件夹
// @param path 目录
// @return *os.File
// @return error
func MakeFileOrPath(path string) (*os.File, error) {
	pathArr := strings.Split(path, "/")
	pathUrl := strings.Join(pathArr[:len(pathArr)-1], "/")
	if err := MakeMultiDir(pathUrl); err != nil {
		return nil, err
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file, nil
}

// WriteContentToFile
// @Description: 写文件
// @param filePath
// @return error
func WriteContentToFile(file *multipart.FileHeader, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	open, err := file.Open()
	if err != nil {
		return err
	}
	defer open.Close()
	fileBytes, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	if _, err := f.Write(fileBytes); err != nil {
		return err
	}
	return nil
}

// MakeTimeFormatDir
// @Description: 创建时间格式的目录 如：upload/{path}/2023-01-07/
// @param rootPath 根目录
// @param pathName 子目录名称
// @param timeFormat 时间格式 如：2006-01-02、20060102
// @return string
// @return error
func MakeTimeFormatDir(rootPath, pathName, timeFormat string) (string, error) {
	filePath := "upload/"
	if pathName != "" {
		filePath += pathName + "/"
	}
	filePath += time.Now().Format(timeFormat) + "/"
	if err := MakeMultiDir(rootPath + filePath); err != nil {
		return "", err
	}
	return filePath, nil
}

// Cmd
// @Description: 直接在当前目录使用并返回结果
// @param commandName
// @param params
// @return string
// @return error
func Cmd(commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

// PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetFileMD5
// @Description: 获取文件的MD5
// @param path 本地文件地址
// @param isRmote 是否网络地址
// @return MD5
// @return err
func GetFileMD5(path string) (MD5 string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(md5Hash.Sum(nil)), nil
}

// IsWindow7 获取windows版本号
func IsWindow7() bool {
	info, _ := host.Info()
	if strings.Contains(info.Platform, "Windows") && strings.Contains(info.Platform, "7") {
		return true
	}
	return false
}

// ArrayChunk 数组分组
func ArrayChunk[T any](arr []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		chunks = append(chunks, arr[i:end])
	}
	return chunks
}

// GetRuntimeUserHomeDir
// @Description: get runtime user home directory
// @return runtimeDir
// @author cx
func GetRuntimeUserHomeDir() (runtimeDir string) {
	switch runtime.GOOS {
	case "windows":
		runtimeDir = os.Getenv("APPDATA")
	case "darwin":
		runtimeDir, _ = os.UserHomeDir()
	default:
		runtimeDir, _ = os.UserHomeDir()
	}
	return
}

// GetStructColumnName 获取结构体中的字段名称 _type: 1: 获取tag字段值 2：获取结构体字段值
func GetStructColumnName(s interface{}, _type int) ([]string, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return []string{}, fmt.Errorf("interface is not a struct")
	}
	t := v.Type()
	var fields []string
	for i := 0; i < v.NumField(); i++ {
		var field string
		if _type == 1 {
			field = t.Field(i).Tag.Get("json")
			if field == "" {
				tagSetting := schema.ParseTagSetting(t.Field(i).Tag.Get("gorm"), ";")
				field = tagSetting["COLUMN"]
			}
		} else {
			field = t.Field(i).Name
		}
		fields = append(fields, field)
	}
	return fields, nil
}
