package account

type AccountType int8

const{
	PHONE_NUMBER 	AccountType =0
	USER_NAME 		AccountType =1
	EMAIL_ADDRESS 	AccountType =2
	INVALID			AccountType =3
}

func(accountType AccountType)String()string{
switch(accountType){
case PHONE_NUMBER:return "PHONE_NUMBER"
case USER_NAME:return "USER_NAME"
case EMAIL_ADDRESS:return "EMAIL_ADDRESS"
default:return "INVALID account type"
}
}