type Trie struct {
	serial   string
	children map[string]*Trie
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	root := &Trie{children: make(map[string]*Trie)}
	for _, path := range paths {
		cur := root
		for _, folder := range path {
			if _, ok := cur.children[folder]; !ok {
				cur.children[folder] = &Trie{children: make(map[string]*Trie)}
			}
			cur = cur.children[folder]
		}
	}

	freq := map[string]int{}

	var serialize func(*Trie) string
	serialize = func(node *Trie) string {
		if len(node.children) == 0 {
			node.serial = ""
			return ""
		}
		v := make([]string, 0, len(node.children))
		for name, child := range node.children {
			sub := serialize(child)
			v = append(v, name+"("+sub+")")
		}
		sort.Strings(v)
		node.serial = strings.Join(v, "")
		freq[node.serial]++
		return node.serial
	}
	serialize(root)

	var res [][]string
	var path []string

	var collect func(*Trie)
	collect = func(node *Trie) {
		if freq[node.serial] > 1 {
			return
		}
		if len(path) > 0 {
			res = append(res, append([]string(nil), path...))
		}
		for name, child := range node.children {
			path = append(path, name)
			collect(child)
			path = path[:len(path)-1]
		}
	}
	collect(root)
	return res
}