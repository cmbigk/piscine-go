package piscine

func Index(parent string, toFind string) int {
	if parent == "" || toFind == "" {
		return 0
	}
	for i := 0; i < len(parent); i++ {
		if i+len(toFind) > len(parent) {
			return -1
		}
		if parent[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
