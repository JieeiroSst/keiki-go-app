package mysql

import  (
	"github.com/JIeeiroSst/go-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	mutex sync.Mutex
	instance *MysqlConn

)

type MysqlConn struct {
	db *gorm.DB
}

type Config struct {
	DSN string
}

func GetMysqlConnInstance(cf Config) *MysqlConn {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			dsn := cf.DSN
			// dsn:="root:1234@tcp(localhost:3306)/db?parseTime=True"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			instance = &MysqlConn{
				db: db,
			}
			_ = db.AutoMigrate(&models.Book{}, &models.User{})
		}
	}
	return instance
}

func NewMysqlRepo(cf Config) *MysqlConn {
	return &MysqlConn{
		db:GetMysqlConnInstance(cf).db,
	}
}