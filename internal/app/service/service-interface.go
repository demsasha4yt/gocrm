package service

// Interface ...
type Interface interface {
	User() *UserService
}

// Service ...
type Service struct {
	userService *UserService
}

// New ...
func New() Interface {
	return &Service{}
}

// User ...
func (s *Service) User() *UserService {
	if s.userService != nil {
		return s.userService
	}

	s.userService = &UserService{
		service: s,
	}

	return s.userService
}
