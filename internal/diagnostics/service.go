package diagnostics

import (
	"log/slog"

	"github.com/arielcr/c64-diagnostic/internal/constants"
)

type Reader interface {
	GetStep(diagnostic string, step string) (Step, error)
	GetDiagnosticMeta() (DiagnosticMeta, error)
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

func (s *Service) GetDiagnostic(status DiagnosticStatus) (Diagnostic, error) {
	step, err := s.Reader.GetStep(status.Diagnostic, status.Step)

	if err != nil {
		return Diagnostic{}, err
	}

	return s.buildDiagnostic(status, step)
}

func (s *Service) buildDiagnostic(status DiagnosticStatus, step Step) (Diagnostic, error) {
	if status.Result == "" {
		return Diagnostic{
			Step: step,
		}, nil
	}

	if status.Result == constants.Yes {
		if step.Success.Diagnostic != "" {
			return Diagnostic{
				Finish:      true,
				Description: step.Success.Diagnostic,
				Step:        step,
			}, nil
		}

		nextStep, err := s.Reader.GetStep(step.Success.Next.Category, step.Success.Next.Step)
		if err != nil {
			return Diagnostic{}, err
		}

		return Diagnostic{
			Step: nextStep,
		}, nil
	}

	if status.Result == constants.No {
		if step.Error.Diagnostic != "" {
			return Diagnostic{
				Finish:      true,
				Description: step.Error.Diagnostic,
				Step:        step,
			}, nil
		}

		nextStep, err := s.Reader.GetStep(step.Error.Next.Category, step.Error.Next.Step)
		if err != nil {
			return Diagnostic{}, err
		}

		return Diagnostic{
			Step: nextStep,
		}, nil
	}

	return Diagnostic{}, nil
}

func (s *Service) GetMetaData() (DiagnosticMeta, error) {
	meta, err := s.Reader.GetDiagnosticMeta()
	if err != nil {
		return DiagnosticMeta{}, err
	}

	return meta, nil
}
