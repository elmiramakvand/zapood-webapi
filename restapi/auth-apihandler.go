package restapi

// type AuthModel struct {
// 	DB *gorm.DB
// }

// var jwtKey = []byte("Zapood_WebApi_2020")

// func NewAuthModel(db *gorm.DB) *AuthModel {
// 	return &AuthModel{
// 		DB: db,
// 	}
// }

// type LoginInfo struct {
// 	Password string `json:"password"`
// 	UserName string `json:"userName"`
// }

// type Claims struct {
// 	UserName string
// 	jwt.StandardClaims
// }

// func (authModel AuthModel) Login(w http.ResponseWriter, r *http.Request) {
// 	var loginInfo LoginInfo
// 	err := json.NewDecoder(r.Body).Decode(&loginInfo)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	var User []entities.User
// 	result := authModel.DB.Where("UserName = ? AND Password >= ?", loginInfo.UserName, loginInfo.Password).Find(&User)
// 	if result.Error != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	sqlRow := result.Row()
// 	var user entities.User
// 	u := entities.User{}
// 	err2 := sqlRow.Scan(&u.ID, &u.Name, &u.Family, &u.UserName, &u.Password)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 		return
// 	} else {
// 		user = u
// 	}
// 	expireTime := time.Now().Add(8 * time.Hour)
// 	claims := &Claims{
// 		UserName: loginInfo.UserName,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expireTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	stringToken, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Println(err)
// 		return
// 	}
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "JWTToken",
// 		Expires:  expireTime,
// 		Value:    stringToken,
// 		HttpOnly: true,
// 	})
// 	json.NewEncoder(w).Encode(user)
// 	return
// }

// func IsLogedin(r *http.Request) (LoginInfo, bool) {

// 	c, err := r.Cookie("JWTToken")
// 	if err != nil {
// 		return LoginInfo{}, false
// 	}

// 	tokenString := c.Value

// 	claims := &Claims{}

// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if !token.Valid {
// 		return LoginInfo{}, false
// 	}
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			return LoginInfo{}, false
// 		}
// 		return LoginInfo{}, false
// 	}

// 	return LoginInfo{
// 		UserName: claims.UserName,
// 	}, true
// }
