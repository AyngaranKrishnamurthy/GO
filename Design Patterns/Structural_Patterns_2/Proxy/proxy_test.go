package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) { //The test case creates a userlist of 10,000 user with random names.
	mockedDatabase := UserList{}
	rand.Seed(2342342) //we make use of random number generator with constant seed to get randomized results that are constant
	for i := 0; i < 10000; i++ {
		n := rand.Int31()
		mockedDatabase = append(mockedDatabase, User{ID: n})
	}

	proxy := UserListProxy{ //proxy with reference to SomeDB
		MockedDatabase: &mockedDatabase,
		StackSize:      2,
		StackCache:     UserList{},
	}

	knownIDs := [3]int32{mockedDatabase[3].ID, mockedDatabase[4].ID, mockedDatabase[5].ID} //random values from SomeDB to store to cache

	t.Run("FindUser - Empty Cache", func(t *testing.T) { //Embedded tests defined by t.Run(<description>, func(t *testing.T) )
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be one")
		}

		if proxy.CacheUsed == true {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object that is stored on it")
		}

		if !proxy.CacheUsed {
			t.Error("The user should have been returned from the cache")
		}
	})

	t.Run("FindUser - overflowing the stack", func(t *testing.T) {
		user1, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		user2, _ := proxy.FindUser(knownIDs[1])
		if proxy.CacheUsed {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		user3, _ := proxy.FindUser(knownIDs[2])
		if proxy.CacheUsed {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		for i := 0; i < len(proxy.StackCache); i++ { //checking if the first entry still exists
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User that should be gone was found")
			}
		}

		if len(proxy.StackCache) != 2 { //checking if the stack size is still the same
			t.Error("After inserting 3 users the cache should not grow" +
				" more than to two")
		}

		for _, v := range proxy.StackCache { //Checking if the last two queried users are still intact
			if v != user2 && v != user3 {
				t.Error("A non expected user was found on the cache")
			}
		}
	})
}
