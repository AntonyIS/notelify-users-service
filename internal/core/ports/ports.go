package ports

import "github.com/AntonyIS/notelify-users-service/internal/core/domain"

type UserService interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ReadUserWithId(user_id string) (*domain.User, error)
	ReadUserWithGithubId(user_id string) (*domain.User, error)
	ReadUserWithLinkedinId(user_id string) (*domain.User, error)
	ReadUserWithEmail(email string) (*domain.User, error)
	ReadUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(user_id string) (string, error)
	DeleteAllUsers() (string, error)
}

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ReadUserWithId(user_id string) (*domain.User, error)
	ReadUserWithGithubId(user_id string) (*domain.User, error)
	ReadUserWithLinkedinId(user_id string) (*domain.User, error)
	ReadUserWithEmail(email string) (*domain.User, error)
	ReadUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(user_id string) (string, error)
	DeleteAllUsers() (string, error)
}

type LoggingService interface {
	SendLog(LogEntry domain.LogMessage)
	LogDebug(LogEntry domain.LogMessage)
	LogInfo(LogEntry domain.LogMessage)
	LogWarning(LogEntry domain.LogMessage)
	LogError(LogEntry domain.LogMessage)
}
