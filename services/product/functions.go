package product

func validateID(id int) bool {
	if id < 0 {
		return false
	}
	return true
}
