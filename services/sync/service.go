package sync

import (
	"go.uber.org/zap"
)

type Service struct {
	filePath     string
	Logger       *zap.Logger
	DealErrorLog *DealErrorLog
}

// 返回一个新实例
func NewSerivce(filePath string) *Service {
	return &Service{
		filePath:     filePath,
		DealErrorLog: NewDealErrorLog(filePath),
		Logger:       zap.NewNop(),
	}
}

func (s *Service) Open() error {
	s.Logger.Info("Start sync service", zap.String("filePath", s.filePath))
	go s.DealErrorLog.TailLog()
	return nil
}

// Close closes the underlying listener.
func (s *Service) Close() error {
	s.Logger.Info("Close service")
	return nil
}

// WithLogger sets the logger for the service.
func (s *Service) WithLogger(log *zap.Logger) {
	s.Logger = log.With(zap.String("service", "sync"))
	s.DealErrorLog.Logger = s.Logger
}
