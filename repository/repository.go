package repository

import (
	"github.com/sam-explorex/demo_math_grpc/database"
	"github.com/sam-explorex/demo_math_grpc/entities"
)

// create log for any operation func CreateEmails(email *entities.Emails) (entities.Emails, error) {
func CreateLogs(log entities.Logs) (entities.Logs, error) {
	err := database.Db.Create(&log).Error
	if err != nil {
		return entities.Logs{}, err

	}
	return log, nil
}
