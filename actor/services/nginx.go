package services

import (
    "fmt"
    amodel "nginx-ui/actor/model"
    "nginx-ui/pkg/logger"
    "nginx-ui/pkg/nginx"
    "nginx-ui/pkg/settings"
    "nginx-ui/pkg/utils"
    "nginx-ui/server/model"
    "os"
    "regexp"
    "strconv"
)

func ServerConfig(id uint) (conf string, err error) {
    cfg := nginx.NginxProxy{}
    cfg.Site, err = model.GetSiteById(uint(id))

    if err != nil {
        return
    }

    conf = cfg.BuildConfig()

    return
}

// ServerConfigReload 重载配置
// 1. 下载配置
// 2. 更新配置
// 3. 回调
func ServerConfigReload(force, restart bool, requestId string) (err error) {
    restarted := false
    siteUpdated, err := ServerConfigDownloading(force)
    if err != nil {
        return
    }
    output := ""
    if siteUpdated > 0 || force {
        if restart {
            output, err = nginx.Restart()
        } else {
            output, err = nginx.Reload()
        }
        if err != nil {
            logger.Errorf("nginx.reload failed err:%v", err)
            report(requestId, output, siteUpdated, restarted)
            return
        }
        restarted = true
    } else {
        output = "no site updates"
    }

    report(requestId, output, siteUpdated, restarted)
    return
}

func report(requestId, content string, siteUpdated uint, restarted bool) (err error) {
    url := settings.NginxSettings.MasterUrl + "/api/node/report"
    headers := map[string]string{
        "x-request-id": requestId,
    }
    data := map[string]any{
        "content":      content,
        "node_id":      settings.NginxSettings.Via,
        "restarted":    2,
        "site_updated": siteUpdated,
    }

    if restarted {
        data["restarted"] = 1
    }
    resp, err := client.R().SetHeaders(headers).
        SetBody(data).Post(url)

    if err != nil {
        return
    }
    fmt.Printf("resp: %s\n", resp.Body())
    return
}

func ServerConfigDownloading(force bool) (siteUpdated uint, err error) {
    url := settings.NginxSettings.MasterUrl + "/api/node/updates"

    var result amodel.SitePageResult
    _, err = client.R().SetResult(&result).Get(url)

    if err != nil {
        return
    }

    //fmt.Printf("resp: %s\n", resp.Body())
    fmt.Println("resultCode:", result.Code)
    fmt.Println("sites count:", len(result.Data.List))
    for _, row := range result.Data.List {
        hasUpdated, errW := DumpConfig(row, force)
        if errW != nil {
            logger.Errorf("writeFile: %s", errW)
            continue
        }

        logger.Infof("write :%t", hasUpdated)
        if hasUpdated {
            siteUpdated++
        }
    }

    return
}

// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cache
func DumpConfig(site model.Site, force bool) (updated bool, err error) {
    lastUpdate := site.UpdatedAt
    fileName := settings.NginxConfigDir() + site.Name + ".conf"
    lastModified := readLastModified(fileName)
    hasUpdated := lastUpdate > lastModified
    fmt.Printf("hasUpdate: %t, file:%d, db:%d\n", hasUpdated, lastModified, lastUpdate)
    if !force && !hasUpdated {
        return false, nil
    }

    fmt.Printf("filename %s\n", fileName)

    proxy := nginx.NewNginxProxy(site)
    conf := proxy.BuildConfig()
    err = utils.Byte2File([]byte(conf), fileName)
    if err != nil {
        logger.Infof("write config(%d_%s) failed", site.ID, site.Name)
        return false, err
    }
    logger.Infof("dumpConfig file :%s", fileName)
    return hasUpdated, nil
}

var modifiedExp = regexp.MustCompile(`lastModified:\s(?P<modified>\d+?)\n`)

func readLastModified(fileName string) (updated int64) {
    exists := utils.FileExist(fileName)
    if !exists {
        return 0
    }

    content, err := os.ReadFile(fileName)
    if err != nil {
        return 0
    }

    r2 := modifiedExp.FindAllSubmatch(content, 1)
    if len(r2) == 0 {
        return 0
    }

    fmt.Printf("r2 %s\n", r2[0][1])
    updated, err = strconv.ParseInt(string(r2[0][1]), 10, 64)
    if err != nil {
        return 0
    }
    return
}

func ServerRestart() {

}
