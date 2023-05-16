package userroutes

import (
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
	responsemodels "github.com/ldehner/fiber-rental-api/models/response"
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

func CreateResponseUser(userModel storemodels.User) responsemodels.User {
	return responsemodels.User{
		Id:          userModel.Id,
		FirstName:   userModel.FirstName,
		LastName:    userModel.LastName,
		Phone:       userModel.Phone,
		Mail:        userModel.Mail,
		Country:     userModel.Country,
		City:        userModel.City,
		Street:      userModel.Street,
		Housenumber: userModel.Housenumber,
		Apartment:   userModel.Apartment,
	}
}

func CreateStoreUser(userModel requestmodels.User) storemodels.User {
	return storemodels.User{
		Id:          userModel.Id,
		FirstName:   userModel.FirstName,
		LastName:    userModel.LastName,
		Phone:       userModel.Phone,
		Mail:        userModel.Mail,
		Country:     userModel.Country,
		City:        userModel.City,
		Street:      userModel.Street,
		Housenumber: userModel.Housenumber,
		Apartment:   userModel.Apartment,
	}
}
