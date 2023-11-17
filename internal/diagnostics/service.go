package diagnostics

import (
	"log/slog"

	"github.com/arielcr/c64-diagnostic/internal/constants"
)

type Reader interface {
	GetStep(diagnostic string, step int) (Step, error)
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

func (s *Service) GetNextStep(status DiagnosticStatus) (Step, error) {
	s.Logger.Info("Service > GetNextStep", "status", status)

	step, err := s.Reader.GetStep(status.Diagnostic, status.Step)
	if err != nil {
		return Step{}, err
	}

	if status.Result == "" {
		return step, nil
	}

	if status.Result == constants.Yes {
		nextStep, err := s.Reader.GetStep(step.Success.Next.Category, step.Success.Next.Step)
		if err != nil {
			return Step{}, err
		}
		return nextStep, nil
	}

	if status.Result == constants.No {
		nextStep, err := s.Reader.GetStep(step.Error.Next.Category, step.Error.Next.Step)
		if err != nil {
			return Step{}, err
		}
		return nextStep, nil
	}

	return Step{}, nil
}
