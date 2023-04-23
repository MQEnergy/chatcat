import {
  CheckAndGetLatestServerVersion,
  DoUpgrade,
  GetDownloadUrlInfo,
  RestartApplication
} from "../../../wailsjs/go/version/Service.js";

export class UpgradeService {
  /**
   * 启动升级
   * @returns {Promise<void>}
   */
  constructor() {
    this.isUpdate = false; // 是否可更新
    this.lastVersionInfo == {}; // 最新版本信息
    this.downloadFileInfo = {}; // 下载信息 包含：文件大小 和 预估下载时间
    this.upgradeInfo = {}; //
  }

  /**
   * 检查是否可更新
   * @returns {Promise<void>}
   */
  async checkUpdate() {
    const checkUpdate = await CheckAndGetLatestServerVersion();
    console.log('checkUpdate', checkUpdate);
    this.lastVersionInfo = checkUpdate.data.versionInfo;
    this.isUpdate = checkUpdate.data.needUpdate;
  }

  /**
   * 获取下载信息
   * @returns {Promise<*|boolean>}
   */
  async getDownloadInfo() {
    console.log('getDownloadInfo')
    if (!this.isUpdate) {
      return false;
    }
    const downloadUrl = this.lastVersionInfo.url;
    const downloadFileInfo = await GetDownloadUrlInfo(downloadUrl);
    this.downloadFileInfo = downloadFileInfo.data;
  }

  /**
   * 下载应用
   * @returns {Promise<*|boolean>}
   */
  async downloadApp() {
    console.log('downloadApp')
    if (!this.isUpdate) {
      return false;
    }
    await this.getDownloadInfo();
    if (this.downloadFileInfo.code == -1) {
      return false;
    }
    const upgradeInfo = await DoUpgrade();
    if (upgradeInfo.code == -1) {
      return false;
    }
    this.upgradeInfo = upgradeInfo.data;
  }

// 安装应用
  async installApp() {
    console.log('installApp')
    if (!this.isUpdate) {
      return false;
    }
    await RestartApplication(this.upgradeInfo.file_path)
  }
}