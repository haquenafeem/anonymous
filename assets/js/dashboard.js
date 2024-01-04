var logoutBtn = document.getElementById("logout-button");

logoutBtn.addEventListener("click", () => {
    localStorage.removeItem('jwt_token')
    location.href = "/login"
})

var qrBtn = document.getElementById("generate-qr");

qrBtn.addEventListener("click", () => {
    downloadQr();
})

async function downloadQr() {
    token = getToken();
    const url = '/api/v1/users/generate-qr-code';

    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token,
        },
    };

    const response = await fetch(url, options)
    if (!response.ok) {
        throw new Error(`Failed to fetch image (${response.status} ${response.statusText})`);
    }

    const file = await response.blob()
    const downloadLink = document.createElement('a');

    const blobUrl = URL.createObjectURL(file);

    downloadLink.href = blobUrl;
    downloadLink.download = 'qr_code.png'; // Change the filename as needed

    document.body.appendChild(downloadLink);

    downloadLink.click();

    document.body.removeChild(downloadLink);
}

function getToken() {
    return localStorage.getItem('jwt_token');
}



function initialize() {
    token = getToken()
    if (token == null || token == '') {
        location.href = '/login'
    } else {
        startMakingCalls(token);
    }
}

function startMakingCalls(token) {
    getProfilePic(token);
    getMessages(token);
}

function getProfilePic(token) {

}

var messagesDiv = document.getElementsByClassName("messages")[0];

function shareMessage(id) {
    console.log(id)
}

async function getMessages(token) {
    token = getToken()
    const url = '/api/v1/messages';

    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token,
        },
    };

    const response = await fetch(url, options)
    if (!response.ok) {
        throw new Error(`Failed to fetch image (${response.status} ${response.statusText})`);
    }

    const jsonData = await response.json();
    console.log(jsonData);


    jsonData.messages.forEach(msg => {
        console.log(msg)
        const messageDiv = document.createElement('div');
        messageDiv.classList.add('message-item');

        const para = document.createElement('p');
        para.textContent = msg.data;

        const shareBtn = document.createElement('button');
        shareBtn.textContent = 'Share';
        shareBtn.onclick = function () {
            shareMessage(msg.id);
        };


        messageDiv.appendChild(para);
        messageDiv.appendChild(shareBtn);

        messagesDiv.appendChild(messageDiv);
    });
}

initialize();