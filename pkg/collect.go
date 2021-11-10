package pkg

func CollectToArray(collectedMap map[string]bool) []string {
	var keys []string

	for k := range collectedMap {
		keys = append(keys, k)
	}
	return keys
}
