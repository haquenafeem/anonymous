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
            'Authorization':'Bearer '+ token,
        },
    };

    const response = await fetch(url, options)
    if (!response.ok){
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

function getMessages(token){

}

initialize();