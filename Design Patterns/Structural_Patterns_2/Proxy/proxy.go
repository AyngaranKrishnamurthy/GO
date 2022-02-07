package proxy

import (
	"fmt"
)

type User struct { //A type *struct* with a member called ID of type int32
	ID int32
}

type UserFiner interface { //An *interface* implemented by database and proxy
	FindUser(id string) (User, error)
}

//We are using a type of slice as with this we can implement the UserFinder interface but with slice we can't
type UserList []User //It is a type of a *slice* of users. It is the Database Representation

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}

	return User{}, fmt.Errorf("User %d could not be found\n", id)
}

func (t *UserList) addUser(newUser User) { //Adds user to cache if found in DB
	*t = append(*t, newUser)
}

type UserListProxy struct { //Proxy
	MockedDatabase *UserList //DB
	StackCache     UserList
	StackSize      int
	CacheUsed      bool //true if cache was used
}

func (u *UserListProxy) addUserToStack(user User) { //if stack is full removes first element and adds new element
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	//Search for the object in the cache list first
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.CacheUsed = true
		return user, nil
	}

	//Object is not in the cache list. Search in the heavy list
	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	//Adds the new user to the stack, removing the last if necessary
	u.addUserToStack(user)

	fmt.Println("Returning user from database")
	u.CacheUsed = false
	return user, nil
}
