import {
  CheckAndGetLatestServerVersion,
  DoUpgrade,
  GetDownloadUrlInfo,
  RestartApplication
} from "../../../wailsjs/go/version/Service.js";

export class UpgradeService {
  /**
   * constructor
   * @returns {Promise<void>}
   */
  constructor() {
    this.isUpdate = false; // is update
    this.lastVersionInfo == {}; // last version info
    this.downloadFileInfo = {}; // download info include: file size, estimated download time.
    this.upgradeInfo = {}; //
  }

  /**
   * check update
   * @returns {Promise<void>}
   */
  async checkUpdate() {
    const checkUpdate = await CheckAndGetLatestServerVersion();
    console.log('checkUpdate', checkUpdate);
    this.lastVersionInfo = checkUpdate.data.versionInfo;
    this.isUpdate = checkUpdate.data.needUpdate;
  }

  /**
   * get download info
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
   * download app
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

// install app
  async installApp() {
    console.log('installApp')
    if (!this.isUpdate) {
      return false;
    }
    await RestartApplication(this.upgradeInfo.file_path)
  }
}