package account

//Login 账户登录
//类型，数据库中用户名是否存在，校验密码是否正确
func Login(account, password string) {
	//account 可能有三种类型
}

//判断账号类型
func getAccountType(account string) AccountType {
	if isMobile:=is.CNMobile(account){
		return AccountType.PHONE_NUMBER
	}

	if isEmail:= is.Email(account){
		return AccountType.EMAIL_ADDRESS
	}

	namePattern:=`^\w{4,32}$`
	nameReg:= regexp.MustCompile(namePattern)
	if isUserName:= nameReg.MatchString(account){
		return AccountType.USER_NAME
	}
	return AccountType.INVALID
}

//
