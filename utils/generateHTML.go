package utils

import (
    "os"
    "fmt"
)

func GenerateHTML() error {
    err := os.Mkdir("htmls", 0777)
    if err != nil && !os.IsExist(err) {
        return err
    }
    names := make([]string, 8)
    names[0] = "emailVerified"
    names[1] = "emailVrfctRequestInvalid"
    names[2] = "passwordReset404"
    names[3] = "passwordResetRequestInvalid"
    names[4] = "emailVrfct404"
    names[5] = "emailVrfctTokenExpired"
    names[6] = "passwordReset"
    names[7] = "passwordResetTokenExpired"

    HTMLcontent := createHTMLcontent()

    for _, name := range names {
        fileName := "htmls/" + name + ".html"
        file, err := os.Create(fileName)
        if err != nil {
            return err
        }
        _, err = file.Write([]byte(HTMLcontent[name]))
        if err != nil {
            return err
        }
    }
    return nil
}

func createHTMLcontent() map[string]string {
    apiHost := os.Getenv("API_HOST")
    apiURL := "\"" + apiHost + "/api/user/handle-pr" + "\""

    return map[string]string {
"emailVerified": `
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Email verified</title>
  <style>
    :root {
      color-scheme: light dark;
    }
    body {
      background-color: #2c2f33;
    }
    .success {
      font-size: 18px;
      color: #f6f7fa;
      width: fit-content;
      margin: 3rem auto;
    }
  </style>
</head>

<body>
  <div class="success" >Success. Your email address has been verified.</div>
</body>

</html>`,
"emailVrfctRequestInvalid": `<!doctype html>
 <html lang="en">
 <head>
   <meta charset="UTF-8"/>
   <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
   <title>Invalid</title>
   <style>
     :root {
       color-scheme: light dark;
     }
     body {
       background-color: #2c2f33;
     }
     .success {
       font-size: 18px;
       color: #f6f7fa;
       width: fit-content;
       margin: 3rem auto;
     }
   </style>
 </head>

 <body>
   <div class="success" >Email verification failed: invalid link.</div>
 </body>

 </html>`,
"passwordReset404": `
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Not found</title>
  <style>
    :root {
      color-scheme: light dark;
    }
    body {
      background-color: #2c2f33;
    }
    .success {
      font-size: 18px;
      color: #f6f7fa;
      width: fit-content;
      margin: 3rem auto;
    }
  </style>
</head>

<body>
  <div class="success" >Resource not found. 404</div>
</body>

</html>`,
"passwordResetRequestInvalid": `
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Invalid</title>
  <style>
    :root {
      color-scheme: light dark;
    }
    body {
      background-color: #2c2f33;
    }
    .success {
      font-size: 18px;
      color: #f6f7fa;
      width: fit-content;
      margin: 3rem auto;
    }
  </style>
</head>

<body>
  <div class="success" >Password reset failed: invalid link.</div>
</body>

</html>`,
"emailVrfct404": `
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Not found</title>
  <style>
    :root {
      color-scheme: light dark;
    }
    body {
      background-color: #2c2f33;
    }
    .success {
      font-size: 18px;
      color: #f6f7fa;
      width: fit-content;
      margin: 3rem auto;
    }
  </style>
</head>
<body>
  <div class="success" >Resource not found. 404</div>
</body>
</html>`,
"emailVrfctTokenExpired": `
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Expired</title>
  <style>
    :root {
      color-scheme: light dark;
    }
    body {
      background-color: #2c2f33;
    }
    .success {
      font-size: 18px;
      color: #f6f7fa;
      width: fit-content;
      margin: 3rem auto;
    }
  </style>
</head>
<body>
  <div class="success" >Verification link has expired.</div>
</body>
</html>`,
"passwordResetTokenExpired": `
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Expired</title>
  <style>
    :root {
      color-scheme: light dark;
    }
    body {
      background-color: #2c2f33;
    }
    .success {
      font-size: 18px;
      color: #f6f7fa;
      width: fit-content;
      margin: 3rem auto;
    }
  </style>
</head>
<body>
  <div class="success" >The password reset link has expired.</div>
</body>
</html>`,
"passwordReset": fmt.Sprintf(`
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>reset</title>
  <style>
    :root {
      color-scheme: light dark;
    }
    body {
      background-color: #2c2f33;
    }
    #container {
      width: 450px;
      height: 250px;
      border: 1px #f6f7fa solid;
      color: #f6f7fa;
      margin: 3rem auto;
      padding: 1rem;
    }
    form {
      margin-top: 1rem;
    }
    #reset-password-prom {
      width: 400px;
      height: 30px;
      margin: 0 auto;
      text-align: center;
      color: #bf660d;
    }
    #new-password-container {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: end;
      margin-top: 10px;
    }
    #confirm-password-container {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: end;
      margin-top: 10px;
    }
    #new-password-prom, #confirm-password-prom {
      width: 300px;
      height: 25px;
      font-size: 15px;
      text-align: center;
      color: #bf660d;
    }
    #new-password, #confirm-password {
      width: 300px;
      height: 15px;
    }
    #submit {
      display: block;
      width: 70px;
      height: 30px;
      margin: 0 auto;
      margin-top: 30px;
      border-radius: 8px;
      border: none;
      background-color: #0d997c;
    }
    #submit:hover {
      cursor: pointer;
      background-color: #11bd9a;
    }
  </style>
</head>

<body>
  <div id="container" >
    <form id="password-form">
      <div id="reset-password-prom"></div>
      <div id="new-password-container">
        <label for="new-password">New password</label>
        <div>
          <div id="new-password-prom"></div>
          <input type="password" id="new-password" name="new-password"/>
        </div>
      </div>
      <div id="confirm-password-container">
        <label for="confirm-password"> Confirm password</label>
        <div>
          <div id="confirm-password-prom"></div>
          <input type="password" id="confirm-password" name="confirm-password"/>
        </div>
      </div>
      <input type="submit" id="submit"/>
    </form>
  </div>
  <script>
    let nPassword = "";
    let cPassword = "";
    const passwordForm = document.getElementById("password-form");
    const resetPasswordProm = document.getElementById("reset-password-prom");
    const newPasswordProm = document.getElementById("new-password-prom");
    const newPassword = document.getElementById("new-password");
    const confirmPasswordProm = document.getElementById("confirm-password-prom");
    const confirmPassword = document.getElementById("confirm-password");
    const submit = document.getElementById("submit");

    passwordForm.addEventListener("submit", handleSubmit);
    newPassword.addEventListener("change", handleNewPasswordChange);
    confirmPassword.addEventListener("change", handleConfirmPasswordChange);

    const passValidConst = {
      INITIAL_PROM: "",
      ZERO_LENGTH: "Please enter password",
      ZERO_LENGTH_CONF_PASS: "",
      SHORTER_THAN_SIX: "Please use at least 8 characters",
      ALL_NUMBER: "Please do not use numbers only",
      NO_NUMBERS: "Please include numbers", 
      NO_UPPER_CASE: "Please include uppercase letters",
      OK: "",
      SUCCESS: "SUCCESS",
      WRONG_PASSWORD: "Password is incorrect",
      INVALID_CHARACTER: "Please use numbers, letters, and $, _, !, -, @",
      PASSWORDS_NOT_SAME: "Passwords must be same",
      NEW_PASSWORDS_NOT_SAME: "New passwords must be same",
    }
    function validatePassword(password) {
      if(password.length  === 0){
        return "ZERO_LENGTH";
      }
      if(password.length < 8){
        return "SHORTER_THAN_SIX";
      }

      const regex = new RegExp("^[0-9a-zA-Z@$_!-]{8,}$");
      const rofRegex = regex.test(password);
      if(rofRegex === true) {
        
        const notAllNum = new RegExp("[^0-9]");
        const rofNotAllNum = notAllNum.test(password);
        if(!rofNotAllNum){
          return "ALL_NUMBER";
        }

        const containsLetter = new RegExp("[^a-zA-Z@$_!-]");
        const rofContainsLetter = containsLetter.test(password);
        if(!rofContainsLetter){
          return "NO_NUMBERS"; 
        }

        const containsUpp = new RegExp("[A-Z]");
        const rofContainsUpp = containsUpp.test(password);
        if(!rofContainsUpp){
          return "NO_UPPER_CASE";
        }
        return "OK";
      }else {
        return "INVALID_CHARACTER"; 
      }
    }
    function handleNewPasswordChange(e) {
      nPassword = e.target.value.trim();
    }
    function handleConfirmPasswordChange(e) {
      cPassword = e.target.value.trim();
    }
    async function handleSubmit(e) {
      e.preventDefault();
      const search = window.location.search;
      const token = search.slice(7, 37).trim();
      const idStr = search.slice(41).trim();
      const id = parseInt(idStr);
      const rofNewPassword = validatePassword(nPassword);
      const rofConfirmPassword = validatePassword(cPassword);
      if (rofNewPassword !== "OK" || rofConfirmPassword !== "OK") {
        newPasswordProm.innerText = passValidConst[rofNewPassword];
        confirmPasswordProm.innerText = passValidConst[rofConfirmPassword];
        return
      }
      if (nPassword !== cPassword) {
        confirmPasswordProm.innerText = passValidConst.PASSWORDS_NOT_SAME;
        return
      }
      const body = {password: nPassword, token, id};
      bodyJson = JSON.stringify(body);
      const request = new Request(%s, {
        method: "POST",
        mode: "cors",
        headers: {
          "Content-Type": "application/json"
        },
        body: bodyJson,
      });
      function clearInputs() {
          newPassword.value = "";
          confirmPassword.value = "";
      }
      try {
        const response = await fetch(request);
        switch (response.status) {
          case 200:
            resetPasswordProm.innerText = "Your password has been reset."
            clearInputs();
            break;
          case 408:
            resetPasswordProm.innerText = "The link has expired."
            clearInputs();
            break;
          case 400:
            resetPasswordProm.innerText = "The link is invalid."
            clearInputs();
            break;
          case 404:
            resetPasswordProm.innerText = "The link is invalid."
            clearInputs();
            break;
          default:
            resetPasswordProm.innerText = "Internal server error, please try later."
            break;
        }
      }catch (err) {
        resetPasswordProm.innerText = "Connection failed, please try later."
      }
    }
  </script>
</body>
</html>`, apiURL),
    }
}
