package entities

// User is a user entity
type User struct {
	ID        uint32
	Username  string
	FirstName string
	LastName  string
	Age       uint8
}

// AtLeastTwice checks if users list contains a user who is at least twice as old as any other person in the list.
// Time comlexity: O(N)
// Space complexity: O(1)
func AtLeastTwice(users []User) bool {
	if len(users) < 2 {
		return false
	}

	min := users[0].Age
	max := users[0].Age
	for _, u := range users[1:] {
		if u.Age < min {
			min = u.Age
		}
		if u.Age > max {
			max = u.Age
		}

		if min*2 <= max {
			return true
		}
	}
	return false
}

// ExactlyTwice checks if users list contains a user who is execly twice as old as any other person in the list.
// Time complexity: O(N)
// Space complexity: O(N)
func ExactlyTwice(users []User) bool {
	if len(users) < 2 {
		return false
	}

	ages := make(map[uint8]struct{}, 100)
	for _, u := range users {
		if _, ok := ages[u.Age*2]; ok {
			return true
		} else if _, ok := ages[u.Age/2]; ok {
			return true
		}
		ages[u.Age] = struct{}{}
	}
	return false
}

// ConstrainedExactlyTwice checks if users list contains a user who is execly twice as old as any other person in the list with condition that
// the user ages are in ragne 18-80.
// Time complexity: O(N)
// Space complexity: O(N)
func ConstrainedExactlyTwice(users []User) bool {
	if len(users) < 2 {
		return false
	}

	ages := make(map[uint8]struct{})
	for _, u := range users {
		switch {
		case u.Age <= 40:
			if _, ok := ages[u.Age*2]; ok {
				return true
			}
			fallthrough
		case u.Age >= 36:
			if _, ok := ages[u.Age/2]; ok {
				return true
			}
			fallthrough
		default:
			ages[u.Age] = struct{}{}
		}
	}
	return false
}
