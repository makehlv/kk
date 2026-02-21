package variable

func (s *VariableService) Add(key, value string) error {
	return s.repositories.Variable.Add(key, value)
}
