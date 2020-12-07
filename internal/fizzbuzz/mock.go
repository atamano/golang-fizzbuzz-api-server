package fizzbuzz

type mockService struct{}

func (s mockService) Compute(params postRequest) string {
	return "mock"
}

//NewMockService for testing fizzbuz
func NewMockService() Service {
	return mockService{}
}
