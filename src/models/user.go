package models

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    userModel, err := UnmarshalUserModel(bytes)
//    bytes, err = userModel.Marshal()

type UserModel struct {
	ID        int64  `json:"id"`
	UserName  string `json:"UserName"`
	Passeord  string `json:"Passeord"`
	Firstname string `json:"Firstname"`
	LastName  string `json:"LastName"`
}
