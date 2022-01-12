package taskmanager
import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"task-manager/Models"
)

func RunMigration(db *gorm.DB){
	db.AutoMigrate(&taskmanager.User{})
	db.AutoMigrate(&taskmanager.Task{})
	db.AutoMigrate(&taskmanager.ReviewersTask{})

}


func ConnectToDb()(*gorm.DB){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	RunMigration(db)
	return db 
}