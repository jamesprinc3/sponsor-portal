package sponsor

import (
	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/model"
)

type Service struct {
	staticFiles string

	Auth auth.Auth
	s3   *model.S3

	model.UserReader
	model.CVReader
}

func New(staticFiles string) (*Service, error) {
	service := Service{
		staticFiles: staticFiles,
	}
	service.setupAuth()

	return &service, nil
}

func (s *Service) SetupDatabase(dbConfig config.Database) error {
	db, err := model.NewDB(dbConfig)
	if err != nil {
		return err
	}

	s.UserReader = model.NewUserReader(db)
	s.CVReader = model.NewCVReader(db)

	return nil
}

func (s *Service) SetupStorer(s3Config config.S3) {
	s.s3 = model.NewS3(s3Config.Aws, s3Config.Bucket, s3Config.Prefix)
}
