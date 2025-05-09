package config
import (
	"log"
	"os"
   
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
   )
   
   var DB *gorm.DB
   
   func ConnectDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	 log.Fatal("Failed to connect to database!")
	}
   }
   func Create[T any](data *T) error{
	return DB.Create(data).Error
   }

   func Delete[T any](data *T) error{
	return DB.Delete(data).Error
   }

   func FindOne[T any] (query *T)(*T , error){
	var data T 
	err:=DB.First(&data,query).Error
	if err !=nil {
		return nil,err
	}
	return &data, nil
   }

   func Update[T any](data *T) error {
	return DB.Save(data).Error
}

   func FindAll[T any](query *T)([]*T, error){
	var data []*T
	err:=DB.Find(&data,query).Error
	return data, err
   }

   func FindWithLimit[T any](query *T , limit int)([]*T, error){
	var data []*T
	err:= DB.Limit(limit).Find(&data, query).Error
    return data, err   
}

   func Count[T any](query *T)(int64, error){
	var count int64
	var model T 
	err:=DB.Model(&model).Where(query).Count(&count).Error
	return count , err
   }