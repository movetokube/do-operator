package utils

func ZeroIntIfNil(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
func ZeroStringIfNil(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
