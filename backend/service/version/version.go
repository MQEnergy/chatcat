package version

import (
	"chatcat/backend/pkg/cresp"
	"chatcat/backend/pkg/cupgrade"
	"chatcat/backend/service"
	"github.com/hashicorp/go-version"
)

type Service struct {
	App *service.App
}

func New(t *service.App) *Service {
	return &Service{
		App: t,
	}
}

// getLocalVersion Service implements Back end interface for version
// 获取本地版本方法（在应用程序初始化时就有版本号）
func (s *Service) getLocalVersion() string {
	return s.App.Cfg.App.Version
}

// needUpdate 是否要更新, 需要传参数
func (s *Service) needUpdate(local string, remote string) bool {
	localVer, localVerErr := version.NewVersion(local)
	if localVerErr != nil {
		return false
	}
	remoteVer, remoteVerErr := version.NewVersion(remote)
	if remoteVerErr != nil {
		return false
	}
	return localVer.LessThan(remoteVer)
}

// CheckAndGetLatestServerVersion 检测和获取最新版本
func (s *Service) CheckAndGetLatestServerVersion() *cresp.Response {
	localVersion := s.getLocalVersion()
	remoteVersionInfo := cupgrade.New(s.App.Cfg).GetLastVersionInfo()
	return cresp.Success(map[string]interface{}{
		"needUpdate":  s.needUpdate(localVersion, remoteVersionInfo.Version),
		"versionInfo": remoteVersionInfo,
	})
}

// DoUpgrade
// @Description: 执行下载
// @receiver s
// @return *cresp.Response
// @author cx
func (s *Service) DoUpgrade() *cresp.Response {
	remoteVersionInfo := cupgrade.New(s.App.Cfg).GetLastVersionInfo()
	filePath, err := cupgrade.New(s.App.Cfg).DoUpgrade(remoteVersionInfo)
	if err != nil {
		return cresp.Fail(err.Error())
	}
	return cresp.Success(map[string]string{
		"file_path":    filePath,
		"download_url": remoteVersionInfo.Url,
	})
}

// GetDownloadUrlInfo
// @Description: 获取下载连接信息
// @receiver s
// @param url 下载地址
// @return *cresp.Response
// @author cx
func (s *Service) GetDownloadUrlInfo(url string) *cresp.Response {
	info, err := cupgrade.New(s.App.Cfg).GetDownloadUrlInfo(url)
	if err != nil {
		return cresp.Fail("获取下载信息异常 err:" + err.Error())
	}
	return cresp.Success(info)
}

// RestartApplication
// @Description: 启动安装新应用
// @receiver s
// @return *cresp.Response
// @author cx
func (s *Service) RestartApplication(filePath string) *cresp.Response {
	if err := cupgrade.New(s.App.Cfg).RestartApplication(filePath); err != nil {
		return cresp.Fail(err.Error())
	}
	s.App.ExitSignalChan <- true
	return cresp.Success("")
}
