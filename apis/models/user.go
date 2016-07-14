package models

import (
    "fmt"
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

type A struct {
    Username    string
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
func (u *User) Add() {
    data, _ := json.Marshal(u)
    if u.isUser() == false {
        dataSource.Set(u.key(), string(data), nil)
    }
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
