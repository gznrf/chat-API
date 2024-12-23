package personalData

/*func (h *Handler) emailVerifyHandler(w http.ResponseWriter, r *http.Request){
	verificationCode, err addCodeToDB(h.dbHandler, input.Username)
	if err != nil {
		if err := deleteUser(h.dbHandler, verificationCode, input, hashedPassword); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := sendEmail(input.Email, verificationCode); err != nil {
		if err := deleteUser(h.dbHandler, verificationCode, input, hashedPassword); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
*/
/*
type EmailCodeInput struct {
	Code     string `json:"code"`
	Username string `json:"username"`
}

func (h *Handler) signUpHandler(w http.ResponseWriter, r *http.Request) {
	var input EmailCodeInput

	if err := utils.DecodeJson(r, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if input.Code == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("code is required"))
		return
	}

	if !isCodeCorrect(h.dbHandler, input.Code, input.Username) {
		utils.WriteError(w, http.StatusBadRequest, errors.New("code is already used"))
		return
	}

	err := utils.WriteJson(w, http.StatusCreated)
}

func isCodeCorrect(dbHandler *gorm.DB, inputCode, username string) bool {
	var currentCode string

	if err := dbHandler.Table("codes").Select("code").
		Joins("JOIN codes_x_users ON codes_x_users.from_id = codes.id").
		Joins("JOIN users ON passwords_x_users.to_id = users.id").
		Where("users.name = ? ", username).
		First(&currentCode).Error; err != nil {
		return false
	}

	if currentCode != inputCode {
		return false
	}

	return true
}*/
/*
func generateVerificationCode() string {
	const (
		charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		codeLength = 6
	)

	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}*/
/*
func sendEmail(toEmail, body string) error {
	sender := viper.GetString("email.sender")
	password := viper.GetString("email.password")
	smtpHost := viper.GetString("email.smtpHost")
	smtpPort := viper.GetString("email.smtpPort")

	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", toEmail, "asd", body))

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", smtpHost, smtpPort), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to the server: %w", err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}

	auth := smtp.PlainAuth("", sender, password, smtpHost)
	if err = c.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	if err = c.Mail(sender); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	if err = c.Rcpt(toEmail); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("failed to get data writer: %w", err)
	}
	_, err = w.Write(message)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close data writer: %w", err)
	}

	c.Quit()

	return nil
}

func addCodeToDB(dbHandler *gorm.DB, username string) (string, error) {
	var verificationCode string
	var userId, codeId int64

	verificationCode = generateVerificationCode()

	if err := dbHandler.Table("users").Select("id").Where("name = ?", username).Scan(&userId).Error; err != nil {
		return "", err
	}
	if err := dbHandler.Table("codes").Select("id").Where("code = ?", verificationCode).Scan(&codeId).Error; err != nil {
		return "", err
	}

	if err := dbHandler.Create(&codesXUsers.CodesXUsers{FromId: codeId, ToId: userId}).Error; err != nil {
		return "", err
	}

	return verificationCode, nil

}*/
