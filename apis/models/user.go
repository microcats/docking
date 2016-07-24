package models

import (
    "fmt"
    "errors"
    "encoding/json"

)

type User struct {
    Username    string  `json:"username"`
    Password    string  `json:"password"`
    Email       string  `json:"email"`
    Enable      bool    `json:"enable"`
    CreateAt    uint64  `json:"createAt"`
    ModifyAt    uint64  `json:"modifyAt"`
    LastLoginAt uint64  `json:"lastLoginAt"`
    IsDeleted   bool    `json:"isDeleted"`
}

func NewUser(username string, password string, email string) *User {
    u := new(User)
    u.Username = username
    u.Password = password
    u.Email = email
    return u
}

// get user key
func (u *User) key() string {
    return fmt.Sprintf("/user/%s", u.Username)
}

// To add a user.
func (u *User) Add() (bool, error) {
    json, err := json.Marshal(u)
    if err != nil {
        return false, errors.New("Invalid User Data.")
    }

    if u.isUser() == true {
        return false, errors.New("Username already exists.")
    }

    if _, err = dataSource.Set(u.key(), string(json), nil); err != nil {
        return false, errors.New("Database Conection Error.")
    }

    return true, nil
}

func (u *User) Get() (User, error) {
    var user User
    resopnse, err := dataSource.Get(u.key(), nil)
    err = json.Unmarshal([]byte(resopnse.Node.Value), &user)
    return user, err
}

/*
// Modify the user information.
func (u *User) modify() {
}

// Delete the user.
func (u *User) delete() {
}

// It enables the user.
func (u *User) enable() {
}

// Disable the user.
func (u *User) disable() {
}
*/
// The user exists.
func (u *User) isUser() bool {
    return dataSource.IsKey(u.key())
}
