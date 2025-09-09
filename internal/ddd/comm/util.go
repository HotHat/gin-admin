package comm

import "strconv"

func StrToID(s string) (ID, error) {

	u64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		// Print an error and skip this element if conversion fails
		//fmt.Printf("Error converting '%s': %v\n", s, err)
		return 0, err
	}

	return ID(u64), nil
}

func IDToStr(id ID) string {
	return strconv.FormatUint(uint64(id), 10)
}

func StrArrToID(s []string) []ID {
	uintSlice := make([]ID, 0, len(s))

	// Loop through each string in the array
	for _, s := range s {
		u64, err := StrToID(s)
		if err != nil {
			continue
		}
		// Convert the uint64 to a uint and append to our new slice
		uintSlice = append(uintSlice, u64)
	}

	return uintSlice
}

func IDArrToStr(ids []ID) []string {
	uintSlice := make([]string, 0, len(ids))

	// Loop through each string in the array
	for _, s := range ids {
		uintSlice = append(uintSlice, IDToStr(s))
	}

	return uintSlice
}
