var logoutBtn = document.getElementById("logout-button");

logoutBtn.addEventListener("click",() => {
    localStorage.removeItem('jwt_token')
    location.href = "/login"
})

function getToken() {
    return localStorage.getItem('jwt_token');
}

function initialize() {
    console.log('called')
    token = getToken()
    console.log(token);
    if (token == null || token == '') {
        location.href = '/login'
    } else {
        startMakingCalls(token);
    }
}

function startMakingCalls(token) {

}

initialize();