package router

import (
    "github.com/gin-contrib/static"
    "github.com/gin-gonic/gin"
    "net/http"
    "nginx-ui/pkg/middleware"
    "nginx-ui/server/api"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    //gin.SetMode(gin.DebugMode)
    r.Use(recovery())
    r.Use(middleware.XRequestID())
    r.Use(cacheJs())

    r.Use(OperationSync())

    r.Use(static.Serve("/", mustFS("")))

    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "not found",
        })
    })

    root := r.Group("/api")
    {
        //root.GET("install", api.InstallLockCheck)
        //root.POST("install", api.InstallNginxUI)

        root.POST("/login", api.Login)
        root.DELETE("/logout", api.Logout)

        // api for nodes
        root.GET("/node/updates", api.Updates)
        root.POST("/node/report", api.Report)
        // backend auth
        w := root.Group("/", authRequired(), proxyWs())
        {
            // Analytic
            w.GET("analytic", api.Analytic)
            w.GET("analytic/intro", api.GetNodeStat)
            w.GET("analytic/nodes", api.GetNodesAnalytic)
            // pty
            w.GET("pty", api.Pty)
            // Nginx log
            w.GET("nginx_log", api.NginxLog)
            w.GET("sys_log", api.GetSysLog)
        }

        g := root.Group("/", authRequired(), proxy())
        {

            g.GET("analytic/init", api.GetAnalyticInit)

            g.GET("users", api.GetUsers)
            g.GET("user/:id", api.GetUser)
            g.POST("user", api.AddUser)
            g.POST("user/:id", api.EditUser)
            g.DELETE("user/:id", api.DeleteUser)

            // Transform NgxConf to nginx configuration
            //g.POST("ngx/build_config", api.BuildNginxConfig)
            // Tokenized nginx configuration to NgxConf
            //g.POST("ngx/tokenize_config", api.TokenizeNginxConfig)
            // Format nginx configuration code
            //g.POST("ngx/format_code", api.FormatNginxConfig)

            g.POST("nginx/reload", api.ReloadNginx)
            //g.POST("nginx/restart", api.RestartNginx)
            //g.POST("nginx/test", api.TestNginx)
            //g.GET("nginx/status", api.NginxStatus)

            //g.POST("domain/:name/enable", api.EnableDomain)
            //g.POST("domain/:name/disable", api.DisableDomain)
            //g.POST("domain/:name/advance", api.DomainEditByAdvancedMode)

            //g.DELETE("domain/:name", api.DeleteDomain)

            //g.POST("domain/:name/duplicate", api.DuplicateSite)
            //g.GET("domain/:name/cert", api.IssueCert)
            g.GET("/sites", api.GetSites)
            g.GET("/sites/:id", api.GetSite)
            g.POST("/sites/:id", api.EditSite)
            g.POST("/sites", api.CreateSite)
            g.GET("/sites/config", api.SiteConfig)
            g.GET("configs", api.GetConfigs)
            g.GET("config/*name", api.GetConfig)
            g.POST("config", api.AddConfig)
            g.POST("config/*name", api.EditConfig)

            //g.GET("backups", api.GetFileBackupList)
            //g.GET("backup/:id", api.GetFileBackup)

            g.GET("template", api.GetTemplate)
            g.GET("template/configs", api.GetTemplateConfList)
            g.GET("template/blocks", api.GetTemplateBlockList)
            g.GET("template/block/:name", api.GetTemplateBlock)
            g.POST("template/block/:name", api.GetTemplateBlock)

            //g.GET("certs", api.GetCertList)
            //g.GET("cert/:id", api.GetCert)
            //g.POST("cert", api.AddCert)
            //g.POST("cert/:id", api.ModifyCert)
            //g.DELETE("cert/:id", api.RemoveCert)

            // Add domain to auto-renew cert list
            //g.POST("auto_cert/:name", api.AddDomainToAutoCert)
            // Delete domain from auto-renew cert list
            //g.DELETE("auto_cert/:name", api.RemoveDomainFromAutoCert)
            //g.GET("auto_cert/dns/providers", api.GetDNSProvidersList)
            //g.GET("auto_cert/dns/provider/:code", api.GetDNSProvider)

            // DNS Credential
            //g.GET("dns_credentials", api.GetDnsCredentialList)
            //g.GET("dns_credential/:id", api.GetDnsCredential)
            //g.POST("dns_credential", api.AddDnsCredential)
            //g.POST("dns_credential/:id", api.EditDnsCredential)
            //g.DELETE("dns_credential/:id", api.DeleteDnsCredential)

            g.POST("nginx_log", api.GetNginxLogPage)

            // cdn_nodes
            g.GET("cdn_nodes", api.GetCdnNodes)
            g.GET("cdn_node/:id", api.GetCdnNode)
            g.POST("cdn_node", api.CreateCdnNode)
            g.POST("cdn_node/:id", api.UpdateCdnNode)
            g.DELETE("cdn_node/:id", api.DeleteCdnNode)

            // Settings
            g.GET("settings", api.GetSettings)
            g.POST("settings", api.CreateSetting)
            g.DELETE("settings/:id", api.DeleteSetting)

            // Upgrade
            //g.GET("upgrade/release", api.GetRelease)
            //g.GET("upgrade/current", api.GetCurrentVersion)
            //g.GET("upgrade/perform", api.PerformCoreUpgrade)

            // ChatGPT
            //g.POST("chat_gpt", api.MakeChatCompletionRequest)
            //g.POST("chat_gpt_record", api.StoreChatGPTRecord)

            // Environment
            //g.GET("environments", api.GetEnvironmentList)
            //envGroup := g.Group("environment")
            //{
            //	envGroup.GET("/:id", api.GetEnvironment)
            //	envGroup.POST("", api.AddEnvironment)
            //	envGroup.POST("/:id", api.EditEnvironment)
            //	envGroup.DELETE("/:id", api.DeleteEnvironment)
            //}

            // node
            g.GET("node", api.GetCurrentNode)

            // translation
            g.GET("translation/:code", api.GetTranslation)
        } // authRequired()
    }

    return r
}
