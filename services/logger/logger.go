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
	appCrashedMessage            = Event{1, "Application crashed"}

	cannotPingMongoDBMessage     = Event{2, "Cannot ping MongoDB"}
	cannotGetUserMessage         = Event{3, "Cannot get user"}
	cannotCreateUserMessage      = Event{4, "Cannot create user"}
	cannotUpdateUserMessage      = Event{5, "Cannot update user"}
	cannotGenerateTokenMessage   = Event{6, "Cannot generate token"}
	cannotCreatePasswordMessage  = Event{7, "Cannot create password"}
	cannotComparePasswordMessage = Event{8, "Cannot compare password"}
	cannotUpdatePasswordMessage  = Event{9, "Cannot update password"}

	didCreateUserMessage         = Event{10, "Created user"}
	didDeleteUserMessage         = Event{11, "Delete user"}
	didUpdateUserMessage         = Event{12, "Update user"}

	didChangePasswordMessage     = Event{13, "Change password user"}

	didLoginMessage              = Event{14, "User logged in"}
)

func (l *StandardLogger) ApplicationCrashed(reason string) {
  l.WithField("reason", reason).Errorf(appCrashedMessage.message)
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
	l.WithField("id", id).Infof(didCreateUserMessage.message, id)
}

func (l *StandardLogger) DidDeleteUser(id string) {
	l.WithField("id", id).Infof(didDeleteUserMessage.message, id)
}

func (l *StandardLogger) DidUpdateUser(id string) {
	l.WithField("id", id).Infof(didUpdateUserMessage.message, id)
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
