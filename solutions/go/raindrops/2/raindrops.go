package raindrops

import "strconv"

func Convert(number int) string {
	base := []struct{
		k int
		v string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	res := ""

	for _, s := range base {
		if number % s.k == 0 {
			res += s.v
		}
	}

	if res == "" {
		return strconv.Itoa(number)
	}

	return res

}
