package services

type SampleService interface {
	SayHello(id int64) (string, error)
}

type sampleService struct {
	db *DB
}

func (s *sampleService) SayHello(id int64) (string, error) {
	message, err := s.db.GetMessage(uint64(id))
	if err != nil {
		return "", err
	}
	return message, nil
}

func NewService(db *DB) SampleService {
	return &sampleService{db: db}
}
