package SimpleGenerator

import "sort"

// //////////////////////////////////////////////////////////////

func sortMapKey[T any](mp map[string]T) []string {
	var listBuf []string
	for key := range mp {
		listBuf = append(listBuf, key)
	}
	sort.Strings(listBuf)

	return listBuf
}
