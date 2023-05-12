package cupgrade

import (
	"chatcat/backend/config"
	"chatcat/backend/pkg/cgit"
	"chatcat/backend/pkg/clog"
	"fmt"
	"github.com/hashicorp/go-version"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Info struct {
	Cfg *config.Conf
}

type Latest struct {
	Version     string `json:"version"`     // 版本号
	VersionDes  string `json:"versionDes"`  // 版本描述
	Url         string `json:"url"`         // 安装包地址
	Md5Hash     string `json:"md5Hash"`     // hash值
	Size        int64  `json:"size"`        // 大小
	ReleaseDate string `json:"releaseDate"` // 版本时间
}

func New(cfg *config.Conf) *Info {
	return &Info{
		Cfg: cfg,
	}
}

// GetLastVersionInfo
// @Description: 获取版本信息
// @receiver i
// @return *Latest
func (i *Info) GetLastVersionInfo() (versionInfo Latest) {
	releaseList, err := cgit.GetGithubReleaseList(i.Cfg.Github.Owner, i.Cfg.Github.Repo)
	if err != nil {
		return
	}
	if len(releaseList) == 0 {
		return
	}
	versionInfo.Version = *releaseList[0].TagName
	versionInfo.VersionDes = *releaseList[0].Body
	versionInfo.Url += i.Cfg.App.UpgradeUrl + i.Cfg.App.Version + "/" + i.newPackageName()
	versionInfo.ReleaseDate = releaseList[0].PublishedAt.String()
	return
}

// newPackageName
// @Description: 创建新包名
// @receiver i
// @return string
func (i *Info) newPackageName() string {
	var packageName = i.Cfg.App.AppName
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64": //64位AMD
			packageName += "-amd64-installer.exe"
		case "arm64": //64位ARM
			packageName += "-arm64-installer.exe"
		case "386": //32位Intel
			packageName += "-win32-installer.exe"
		}
	case "darwin": // 暂时只支持 amd64
		switch runtime.GOARCH {
		case "amd64":
			packageName += "-amd64-installer.dmg"
		case "arm64":
			packageName += "-arm64-installer.dmg"
		}
	}
	return packageName
}

// GetDownloadUrlInfo
// @Description: 获取下载资源信息
// @receiver i
// @param url
// @return map[string]interface{}
// @author cx
func (i *Info) GetDownloadUrlInfo(url string) (map[string]interface{}, error) {
	const DownloadSpeed = 5.0 // 下载速度为 5MB/s
	// 发送 HEAD 请求获取文件大小和文件类型等信息
	res, err := http.Head(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// 获取文件大小
	filesize := res.ContentLength
	// 计算预估下载时间
	downloadTime := float64(filesize) / (1024.0 * 1024.0 * DownloadSpeed)
	return map[string]interface{}{
		"file_size":     filesize,
		"download_time": downloadTime,
	}, nil
}

// downloadLatestVersion
// @Description: 下载最新版本
// @receiver i
// @param downloadUrl
// @return string
// @return error
// @author cx
func (i *Info) downloadLatestVersion(downloadUrl string) (string, error) {
	var databaseHome string

	resp, err := http.Get(downloadUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	downloadPath := strings.Split(downloadUrl, "/")
	packageName, err := url.QueryUnescape(downloadPath[len(downloadPath)-1])
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		databaseHome = os.Getenv("APPDATA")
	} else {
		databaseHome, _ = os.UserHomeDir()
	}
	pathExecutable, _ := os.Executable()
	_, pathAppname := filepath.Split(pathExecutable)
	filePath := databaseHome + "/" + pathAppname + "/" + packageName
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return filePath, err
}

// RestartApplication
// @Description: 执行安装应用
// @receiver i
// @return error
// @author cx
func (i *Info) RestartApplication(filePath string) error {
	clog.PrintInfo("RestartApplication", filePath)
	// 启动新的进程
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "start "+filePath)
		return cmd.Start()
	}
	return fmt.Errorf("请自行安装 新版本存放目录：%s", filePath)
}

// DoUpgrade
// @Description: 更新操作
// @receiver i
// @return *Latest
// @return error
func (i *Info) DoUpgrade(versionInfo Latest) (string, error) {
	if versionInfo.Version == "" {
		return "", fmt.Errorf("版本获取失败")
	}
	localVer, _ := version.NewVersion(i.Cfg.App.Version)
	remoteVer, _ := version.NewVersion(versionInfo.Version)
	if !localVer.LessThan(remoteVer) {
		return "", fmt.Errorf("当前版本已是最新版本")
	}
	filePath, err := i.downloadLatestVersion(versionInfo.Url)
	if err != nil {
		return "", fmt.Errorf("下载最新版本失败 err: %s", err.Error())
	}
	return filePath, nil
}
