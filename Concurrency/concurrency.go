var userTotal map[string]int

func init() {
	userTotal = make(map[string]int)
}

func CreateUser(cid string) {
	go func() {
		userTotal[cid]++
	}()
}

func DeleteUser(cid string) {
	go func() {
		userTotal[cid]--
	}()
}

func GetUserTotal(cid string) int {
	return userTotal[cid]
}

// —-------------------------------
// Can you mention some of the possible issues here?
