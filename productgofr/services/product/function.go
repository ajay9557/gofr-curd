package product

func validateID(id int) bool {
	if id > 0 {
		return true
	}
	return false
}
