package diagnostics

import (
	"log/slog"
)

type Reader interface {
	Query(status DiagnosticStatus) (Diagnostic, error)
}

type Service struct {
	Reader Reader
	Logger *slog.Logger
}

func NewService(reader Reader, logger *slog.Logger) *Service {
	newService := Service{
		Reader: reader,
		Logger: logger,
	}

	return &newService
}

func (s *Service) Query(status DiagnosticStatus) (Diagnostic, error) {

	result, _ := s.Reader.Query(status)

	return result, nil
}
