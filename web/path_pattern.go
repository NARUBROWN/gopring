package web

import "strings"

// Spring의 AntPathMatcher
// pattern = /users/{id}
// actual = "/users/10"
func matchPatch(pattern, actual string) (map[string]string, bool) {
	pParts := strings.Split(pattern, "/")
	aParts := strings.Split(actual, "/")

	// /로 분리했을때, 다르면 거짓값 반환
	if len(pParts) != len(aParts) {
		return nil, false
	}

	vars := map[string]string{}

	for i := range pParts {
		// {로 시작한다면 PathVariable
		if strings.HasPrefix(pParts[i], "{") {
			key := strings.Trim(pParts[i], "{}")
			vars[key] = aParts[i]
			continue
		}

		// PathVariable이 아닌 구간이라 서로 다르면 거짓값 반환
		if pParts[i] != aParts[i] {
			return nil, false
		}
	}

	return vars, true
}
