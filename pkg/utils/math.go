package utils

// CalculatePageSkip calculates the offset for pagination
//
// Parameters:
//   - page: page number
//   - limit: limit per page
//
// Returns:
//   - offset: offset for pagination
func CalculatePageSkip(page int, limit int) int {
	return (page - 1) * limit
}
