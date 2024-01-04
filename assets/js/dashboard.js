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

async function startMakingCalls(token) {
    await getProfilePic(token);
    await getMessages(token);
}

var profilePic = document.getElementById('profile_img');

async function getProfilePic(token) {
    const url = '/api/v1/users/profile-pic';

    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token,
        },
    };

    const response = await fetch(url, options)
    if (!response.ok) {
        // throw new Error(`Failed to fetch image (${response.status} ${response.statusText})`);
        return
    }

    const jsonData = await response.json()
    console.log(jsonData)

    profilePic.src = '/img/' + jsonData['profile_pic_id']
}

var messagesDiv = document.getElementsByClassName("messages")[0];

function shareMessage(id) {
    console.log(id)
}

async function getMessages(token) {
    const url = '/api/v1/messages';

    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token,
        },
    };

    const response = await fetch(url, options)

    const jsonData = await response.json();
    console.log(jsonData);

    if (jsonData.messages.length == 0) {
        const messageDiv = document.createElement('div');
        messageDiv.classList.add('message-item');

        const para = document.createElement('p');
        para.textContent = "no messages yet :(";
        messageDiv.appendChild(para);
        messagesDiv.appendChild(messageDiv);
    }

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

function openFileUploader() {
    document.getElementById('fileInput').click();
}

async function handleFileSelect() {
    const selectedFile = document.getElementById('fileInput').files[0];

    if (selectedFile) {
        console.log(selectedFile)
        const formData = new FormData();
        formData.append('upload_file', selectedFile);

        const url = '/api/v1/users/upload';
        token = getToken();
        const options = {
            method: 'POST',
            headers: {
                'Authorization': 'Bearer ' + token,
            },
            body: formData,
        };

        const response = await fetch(url, options)

        const jsonData = await response.json();
        console.log(jsonData);

        if (jsonData['is_success']){
            location.href = "/dashboard"
        }
    }
}

initialize();