package normalize

import "net"

// Address returns addr with the passed default port appended if there is not
// already a port specified.
func Address(addr, defaultPort string) (a string, e error) {
	var p string
	a, p, e = net.SplitHostPort(addr)
	if log.E.Chk(e) || p == "" {
		return net.JoinHostPort(a, defaultPort), e
	}
	return net.JoinHostPort(a, p), e
}

// Addresses returns a new slice with all the passed peer addresses normalized
// with the given default port, and all duplicates removed.
func Addresses(addrs []string, defaultPort string) (a []string, e error) {
	for i := range addrs {
		addrs[i], e = Address(addrs[i], defaultPort)
	}
	return addrs, e
}

// RemoveDuplicateAddresses returns a new slice with all duplicate entries in
// addrs removed.
func RemoveDuplicateAddresses(addrs []string) (result []string) {
	result = make([]string, 0, len(addrs))
	seen := map[string]struct{}{}
	for _, val := range addrs {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = struct{}{}
		}
	}
	return result
}
