package logger

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

var instance *StandardLogger

func GetLogger() *StandardLogger {
	if instance == nil {
		instance = NewLogger()
	}
	return instance
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}
	standardLogger.SetReportCaller(true)

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	appCrashedMessage               = Event{1, "APPLICATION_CRASHED"}

	cannotLoadEnvironmentFile       = Event{2, "CANNOT_LOAD_ENVIRONMENT_FILE"}

	cannotPingMongoDBMessage        = Event{3, "CANNOT_PING_MONGODB"}
	cannotGetMongoDBInstanceMessage = Event{4, "CANNOT_GET_MONGODB_INSTANCE"}
	cannotMigrateDatabaseMessage    = Event{5, "CANNOT_MIGRATE_DATABASE"}
  
	cannotGetUserMessage            = Event{6, "CANNOT_GET_USER"}
	cannotCreateUserMessage         = Event{7, "CANNOT_CREATE_USER"}
	cannotUpdateUserMessage         = Event{8, "CANNOT_UPDATE_USER"}
	cannotGenerateTokenMessage      = Event{9, "CANNOT_GENERATE_TOKEN"}
	cannotCreatePasswordMessage     = Event{10, "CANNOT_CREATE_PASSWORD"}
	cannotComparePasswordMessage    = Event{11, "CANNOT_COMPARE_PASSWORD"}
	cannotUpdatePasswordMessage     = Event{12, "CANNOT_UPDATE_PASSWORD"}

	didCreateUserMessage            = Event{13, "CREATED_USER"}
	didDeleteUserMessage            = Event{14, "DELETE_USER"}
	didUpdateUserMessage            = Event{15, "UPDATE_USER"}

	didChangePasswordMessage        = Event{15, "CHANGE_USER_PASSWORD"}

	didLoginMessage                 = Event{16, "USER_LOGGED_IN"}
)

func (l *StandardLogger) ApplicationCrashed(reason string) {
  l.WithField("reason", reason).Errorf(appCrashedMessage.message)
}

func (l *StandardLogger) CannotLoadEnvironmentFile(reason string) {
  l.WithField("reason", reason).Errorf(cannotLoadEnvironmentFile.message)
}

func (l *StandardLogger) CannotMigrateDatabase(reason string) {
  l.WithField("reason", reason).Errorf(cannotMigrateDatabaseMessage.message)
}

func (l *StandardLogger) CannotGetMongoDBInstance(reason string) {
	l.WithField("reason", reason).Errorf(cannotGetMongoDBInstanceMessage.message)
}

func (l *StandardLogger) CannotPingMongoDB(reason string) {
	l.WithField("reason", reason).Errorf(cannotPingMongoDBMessage.message)
}

func (l *StandardLogger) CannotCreatePassword(reason string) {
	l.WithField("reason", reason).Errorf(cannotCreatePasswordMessage.message)
}

func (l *StandardLogger) CannotUpdateUser(reason string) {
	l.WithField("reason", reason).Errorf(cannotUpdateUserMessage.message)
}

func (l *StandardLogger) CannotCreateUser(reason string) {
	l.WithField("reason", reason).Errorf(cannotCreateUserMessage.message)
}

func (l *StandardLogger) CannotGetUser(reason string) {
	l.WithField("reason", reason).Errorf(cannotGetUserMessage.message)
}

func (l *StandardLogger) CannotGenerateToken(reason string) {
	l.WithField("reason", reason).Errorf(cannotGenerateTokenMessage.message)
}

func (l *StandardLogger) CannotComparePassword(reason string) {
	l.WithField("reason", reason).Errorf(cannotComparePasswordMessage.message)
}

func (l *StandardLogger) CannotChangePassword(reason string) {
	l.WithField("reason", reason).Errorf(cannotUpdatePasswordMessage.message)
}

func (l *StandardLogger) DidCreateUser(id string) {
	l.WithField("id", id).Infof(didCreateUserMessage.message)
}

func (l *StandardLogger) DidDeleteUser(id string) {
	l.WithField("id", id).Infof(didDeleteUserMessage.message)
}

func (l *StandardLogger) DidUpdateUser(id string) {
	l.WithField("id", id).Infof(didUpdateUserMessage.message)
}

func (l *StandardLogger) DidChangePassword(id string) {
	l.WithField("id", id).Infof(didChangePasswordMessage.message)
}

func (l *StandardLogger) DidLogin(id string, ip string) {
	l.WithFields(logrus.Fields{
		"user": id,
		"ip":   ip,
	}).Infof(didLoginMessage.message)
}
