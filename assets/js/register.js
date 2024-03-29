import { successFlash,errorFlash } from "./flash.js";

var registerBtn = document.getElementById("register-button");

var email = document.getElementById("email");
var password = document.getElementById("password");


registerBtn.addEventListener("click", async () => {
    const url = '/api/v1/users/register';

    const data = {
        email: email.value,
        password: password.value,
    };

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    };

    const response = await fetch(url, options)
    const resData = await response.json()

    console.log(resData)

    if (resData['is_success']){
        await successFlash('registration successful');
        location.href = "/login"
    }else{
        errorFlash(resData['err'])
    }
})

