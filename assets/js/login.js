import { successFlash,errorFlash } from "./flash.js";

var loginBtn = document.getElementById("login-button");

var email = document.getElementById("email");
var password = document.getElementById("password");


loginBtn.addEventListener("click", async () => {
    const url = '/api/v1/users/login';

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
        await successFlash('login successful');
        localStorage.setItem('jwt_token',resData['jwt_token'])
        location.href = "/dashboard"
    }else{
        errorFlash(resData['err'])
    }
})
