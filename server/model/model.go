package model

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
    "github.com/spf13/cast"
    "gorm.io/driver/mysql"
    "gorm.io/gen"
    "gorm.io/gorm"
    gormlogger "gorm.io/gorm/logger"
    "log"
    "nginx-ui/pkg/logger"
    "nginx-ui/pkg/settings"
    "nginx-ui/server/model/soft_delete"
    "os"
    "time"
)

const (
    StateEnable   = 1
    StateDisabled = 2
)

var (
    db  *gorm.DB
    rdb *redis.Client
)

type BaseModel struct {
    ID        uint                   `gorm:"primary_key" json:"id"`
    CreatedAt int64                  `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt int64                  `json:"updated_at" gorm:"autoUpdateTime"`
    DeletedAt *soft_delete.DeletedAt `gorm:"column:deleted_at;default:0" json:"deleted_at"`
}

func GenerateAllModel() []any {
    return []any{
        ConfigBackup{},
        Auth{},
        AuthToken{},
        Cert{},
        ChatGPTLog{},
        Site{},
        DnsCredential{},
        Environment{},
    }
}

func logMode() gormlogger.Interface {
    switch settings.ServerSettings.RunMode {
    case gin.ReleaseMode:
        return gormlogger.Default.LogMode(gormlogger.Warn)
    default:
        fallthrough
    case gin.DebugMode:
        return gormlogger.Default.LogMode(gormlogger.Info)
    }
}

func InitRedis() *redis.Client {
    rdb = redis.NewClient(&redis.Options{
        Addr:     settings.RedisSettings.Addr,
        Password: settings.RedisSettings.Password,
        DB:       settings.RedisSettings.Db,
    })

    return rdb
}

func Init() *gorm.DB {
    dsn := settings.DbSettings.Dsn
    fmt.Println("DbSettings:", settings.DbSettings)
    fmt.Println("dsn:", dsn)
    _default := gormlogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gormlogger.Config{
        SlowThreshold: 200 * time.Millisecond,
        LogLevel:      gormlogger.Info,
        Colorful:      true,
    })
    dbConfig := &gorm.Config{}
    dbConfig.Logger = _default
    var err error
    db, err = gorm.Open(mysql.Open(dsn), dbConfig)
    if err != nil {
        logger.Fatal(err.Error())
    }

    // Migrate the schema
    //err = db.AutoMigrate(GenerateAllModel()...)
    //if err != nil {
    //	logger.Fatal(err.Error())
    //}

    return db
}

func GetDb() *gorm.DB {
    return db
}

func orderAndPaginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        sort := c.DefaultQuery("sort", "desc")
        order := c.DefaultQuery("order_by", "id") +
            " " + sort

        page := cast.ToInt(c.Query("page"))
        if page == 0 {
            page = 1
        }
        pageSize := settings.ServerSettings.PageSize
        reqPageSize := c.Query("page_size")
        if reqPageSize != "" {
            pageSize = cast.ToInt(reqPageSize)
        }
        offset := (page - 1) * pageSize

        return db.Order(order).Offset(offset).Limit(pageSize)
    }
}

func totalPage(total int64, pageSize int) int64 {
    n := total / int64(pageSize)
    if total%int64(pageSize) > 0 {
        n++
    }
    return n
}

type Pagination struct {
    Total       int64 `json:"total"`
    PerPage     int   `json:"per_page"`
    CurrentPage int   `json:"current_page"`
    TotalPages  int64 `json:"total_pages"`
}

type DataList struct {
    Data       any        `json:"data"`
    Pagination Pagination `json:"pagination,omitempty"`
}

func GetListWithPagination(models any,
    c *gin.Context, totalRecords int64) (result DataList) {

    page := cast.ToInt(c.Query("page"))
    if page == 0 {
        page = 1
    }

    result = DataList{}

    result.Data = models

    pageSize := settings.ServerSettings.PageSize
    reqPageSize := c.Query("page_size")
    if reqPageSize != "" {
        pageSize = cast.ToInt(reqPageSize)
    }

    result.Pagination = Pagination{
        Total:       totalRecords,
        PerPage:     pageSize,
        CurrentPage: page,
        TotalPages:  totalPage(totalRecords, pageSize),
    }

    return
}

type Method interface {
    // FirstByID Where("id=@id")
    FirstByID(id int) (*gen.T, error)
    // DeleteByID update @@table set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=@id
    DeleteByID(id int) error
}
