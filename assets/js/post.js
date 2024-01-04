import { successFlash,errorFlash } from "./flash.js";

function getUserIdFromPath() {
    const dynamicPath = window.location.pathname;

    const id = dynamicPath.split("/")[2]

    return id
}

var postBtn = document.getElementById("post-button");

postBtn.addEventListener("click", async () => {
    const id = getUserIdFromPath();
    const url = '/api/v1/messages';

    const data = {
        user_id: id,
        message: 'asdasda',
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
        await successFlash('message sent');
    }else{
        errorFlash(resData['err'])
    }
})