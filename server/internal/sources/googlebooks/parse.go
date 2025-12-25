package googlebooks

func filterVolumes(items []GoogleBooksVolumeItem, language string) []GoogleBooksVolumeItem {
	var filteredItems []GoogleBooksVolumeItem
	for _, item := range items {
		if item.VolumeInfo.Language == language {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
