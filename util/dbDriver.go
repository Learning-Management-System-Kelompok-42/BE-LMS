package util

import (
	"context"
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/record"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseDriver database driver enum
type DatabaseDriver string

const (
	// MongoDB Database driver
	MongoDB DatabaseDriver = "mongodb"

	// MySQL Database driver
	MySQL DatabaseDriver = "mysql"

	// PostgreSQL Database driver
	PostgreSQL DatabaseDriver = "postgres"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	// for MySQL
	MySQL *gorm.DB

	// for PostgreSQL
	PostgreSQL *gorm.DB

	// for MongoDB
	MongoDB     *mongo.Database
	mongoClient *mongo.Client
}

func NewConnectionDB(config *config.AppConfig) *DatabaseConnection {
	var dbConnection DatabaseConnection

	switch config.Database.Driver {
	case "mysql":
		dbConnection.Driver = MySQL
		dbConnection.MySQL = NewMySQLConnection(config)
	case "postgres":
		dbConnection.Driver = PostgreSQL
		dbConnection.PostgreSQL = NewPostgreSQLConnection(config)
	case "mongodb":
		dbConnection.Driver = MongoDB
		dbConnection.mongoClient = NewMongoDBConnection(config)
		dbConnection.MongoDB = dbConnection.mongoClient.Database(config.Database.Name)
	default:
		panic("Database driver not supported")
	}

	return &dbConnection
}

func NewMySQLConnection(config *config.AppConfig) *gorm.DB {
	var uri string
	uri = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(
		&record.Specialization{},
		&record.Company{},
		&record.User{},
		&record.Course{},
		&record.RequestCourse{},
		&record.Faq{},
		&record.UserCourse{},
		&record.Certificate{},
		&record.UserModule{},
		&record.Module{},
		&record.SpecializationCourse{},
		&record.Material{},
		&record.Quiz{},
		&record.Option{},
	); err != nil {
		fmt.Println("error migrate = ", err.Error())
		panic(err)
	}

	return db
}

func NewPostgreSQLConnection(config *config.AppConfig) *gorm.DB {
	var uri string
	uri = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Address,
		config.Database.Port,
		config.Database.Username,
		config.Database.Name,
		config.Database.Password,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func NewMongoDBConnection(config *config.AppConfig) *mongo.Client {
	var uri string
	uri = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client
}

func (db *DatabaseConnection) Close() {
	switch db.Driver {
	case "mysql":
		db, _ := db.MySQL.DB()
		db.Close()
	case "postgres":
		db, _ := db.PostgreSQL.DB()
		db.Close()
	case "mongodb":
		db.mongoClient.Disconnect(context.Background())
	default:
		panic("Database driver not supported")
	}
}
