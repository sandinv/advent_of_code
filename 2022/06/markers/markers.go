package markers

type Signal map[rune]struct{}

const marker_length = 4

func FindMarker(signals string) int {

	for i := 0; i < len(signals); i++ {
		signal := Signal{}
		repeated := false
		for j := i; j < (i + marker_length); j++ {
			item := rune(signals[j])
			if _, ok := signal[item]; ok {
				repeated = true
				break
			} else {
				signal[item] = struct{}{}
			}
		}
		if !repeated {
			return i + marker_length
		}
	}
	return 0

}
