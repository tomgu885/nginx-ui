package services

import (
    "fmt"
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

// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cache
func DumpConfig(site model.Site, force bool) (updated bool, err error) {
    lastUpdate := site.UpdatedAt
    fileName := settings.NginxConfigDir() + site.Name + ".conf"
    lastModified := readLastModified(fileName)
    if !force && lastModified >= lastUpdate {
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
    return true, nil
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

func ServerUpdate() {

}
