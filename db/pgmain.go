package db

import (
	"fmt"
	"sync"
	"time"

	guuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //database
var once sync.Once

// Postgre cocnnection start
func InitPostgre() {
	username := "postgres"
	password := "root"
	dbName := "NewDRD"
	dbHost := "localhost"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	if !db.HasTable(&Company{}) {
		db.CreateTable(&Company{})
	}
	if !db.HasTable(&AppToken{}) {
		db.CreateTable(&AppToken{})
	}
	if !db.HasTable(&ApiLog{}) {
		db.CreateTable(&ApiLog{})
	}
	if !db.HasTable(&ApiType{}) {
		db.CreateTable(&ApiType{})
	}
	if !db.HasTable(&Session{}) {
		db.CreateTable(&Session{})
	}
	if !db.HasTable(&AuditTrails{}) {
		db.CreateTable(&AuditTrails{})
	}
	db.Debug().AutoMigrate(&User{}, &Company{}, &AppToken{}, &ApiLog{}, &ApiType{}, &Session{}, &AuditTrails{})
	// db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration

	db.Debug().Model(&AppToken{}).AddForeignKey("company_id", "companies(ID)", "CASCADE", "CASCADE")
	db.Debug().Model(&ApiLog{}).AddForeignKey("app_token_id", "app_tokens(ID)", "CASCADE", "CASCADE")
	db.Debug().Model(&ApiLog{}).AddForeignKey("api_type_id", "api_logs(ID)", "CASCADE", "CASCADE")
	db.Debug().Model(&Session{}).AddForeignKey("user_id", "users(ID)", "CASCADE", "CASCADE")
	db.Debug().Model(&AuditTrails{}).AddForeignKey("user_id", "users(ID)", "CASCADE", "CASCADE")
	db.Debug().Model(&AuditTrails{}).AddForeignKey("company_id", "companies(ID)", "CASCADE", "CASCADE")

	company := Company{Name: "AIA", Email: "aia@aia.com", CreatedAt: time.Now()}
	db.FirstOrCreate(&company)
	apptoken := AppToken{Company: company}
	db.FirstOrCreate(&apptoken)
	fmt.Println("Successfully connected!")
}

//returns a handle to the DB object

func GetDb() *gorm.DB {
	// Initiate value if there is no instance
	once.Do(func() {
		InitPostgre()
	})
	return db
}

type User struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Gender       string
	Email        string
	KtpNumber    int64 `gorm:"unique;not null"`
	Address      string
	PhoneNumber  string
	DateOfBirth  time.Time
	Cityzenship  string
	PlaceOfBirth string
	IsActive     bool
	Signature    string
	Signature2   string
	Signature3   string
}
type Company struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Email     string
}
type AppToken struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Token     string
	IsActive  bool
	CompanyID string
	Company   Company
}
type ApiLog struct {
	ID                int `gorm:"primary_key"`
	Timestamp         time.Time
	Ttl               int
	ApiResponseStatus int
	AppTokenID        int
	ApiTypeID         int
	AppToken          AppToken
	ApiType           ApiType
}
type ApiType struct {
	ID          int `gorm:"primary_key"`
	Path        string
	Method      string
	Url         string
	Description string
}
type Session struct {
	ID        int `gorm:"primary_key"`
	Token     string
	CreatedAt time.Time
	Duration  int
	UserID    string
	User      User
}
type AuditTrails struct {
	Id        int `gorm:"primary_key"`
	Timestamp time.Time
	Activity  string
	UserID    string
	CompanyID string
	User      User
	Company   Company
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", guuid.New().String())
	scope.SetColumn("IsActive", true)
	return nil
}

func (company *Company) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", guuid.New().String())
	return nil
}

func (apptoken *AppToken) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Token", guuid.New().String())
	scope.SetColumn("IsActive", true)
	return nil
}

func (session *Session) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Token", guuid.New().String())
	return nil
}
