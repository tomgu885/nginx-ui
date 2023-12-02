package boot

import (
    "context"
    "fmt"
    "github.com/google/uuid"
    "mime"
    "nginx-ui/pkg/logger"
    "nginx-ui/pkg/settings"
    "nginx-ui/server/internal/analytic"
    "nginx-ui/server/model"
    "nginx-ui/server/query"
    "runtime"
)

func Kernel() {
    defer recovery()

    async := []func(){
        InitJsExtensionType,
        InitDatabase,
        InitNodeSecret,
    }

    syncs := []func(){
        analytic.RecordServerAnalytic,
    }

    for _, v := range async {
        v()
    }

    for _, v := range syncs {
        go v()
    }
}

func InitAfterDatabase() {
    syncs := []func(){
        InitAutoObtainCert,
        analytic.RetrieveNodesStatus,
        model.LoadSettFromDb,
    }

    for _, v := range syncs {
        go v()
    }
}

func recovery() {
    if err := recover(); err != nil {
        buf := make([]byte, 1024)
        runtime.Stack(buf, false)
        logger.Errorf("%s\n%s", err, buf)
    }
}

func InitDatabase() {
    rdb := model.InitRedis()
    pong, err := rdb.Ping(context.Background()).Result()
    if err != nil {
        panic("fail to ping redis")
    }
    fmt.Println("ping redis:", pong)
    db := model.Init()
    if "" != settings.ServerSettings.JwtSecret {

        query.Init(db)

        InitAfterDatabase()
    }
}

func InitNodeSecret() {
    if "" == settings.ServerSettings.NodeSecret {
        logger.Warn("NodeSecret is empty, generating...")
        settings.ServerSettings.NodeSecret = uuid.New().String()
        settings.ReflectFrom()

        err := settings.Save()
        if err != nil {
            logger.Error("Error save settings")
        }
        logger.Warn("Generated NodeSecret: ", settings.ServerSettings.NodeSecret)
    }
}

func InitJsExtensionType() {
    // Hack: fix wrong Content Type of .js file on some OS platforms
    // See https://github.com/golang/go/issues/32350
    _ = mime.AddExtensionType(".js", "text/javascript; charset=utf-8")
}

func InitAutoObtainCert() {
    //s := gocron.NewScheduler(time.UTC)
    //job, err := s.Every(30).Minute().SingletonMode().Do(cert.AutoObtain)
    //
    //if err != nil {
    //    logger.Fatalf("AutoCert Job: %v, Err: %v\n", job, err)
    //}
    //
    //s.StartAsync()
}
