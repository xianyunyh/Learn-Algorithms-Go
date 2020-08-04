package lcoffer

func permutation2(s string) []string {
	var ret []string
	var m = make(map[string]struct{})
	bfs2(s, "", m)

	for key := range m {
		ret = append(ret, key)
	}

	return ret
}

func bfs2(remained string, now string, m map[string]struct{}) {
	if len(remained) == 0 {
		m[now] = struct{}{}
		return
	}
	for i := 0; i < len(remained); i++ {
		s := remained[0:i] + remained[i+1:]
		bfs2(s, now+string(remained[i]), m)
	}
}
