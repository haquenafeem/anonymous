export async function successFlash(message) {
    return new Promise(resolve => {
        const flash = document.createElement('div');
        flash.id = "flash-card";

        const paragraph = document.createElement('p')
        paragraph.textContent = 'Success : ' + message;

        flash.appendChild(paragraph);
        document.body.appendChild(flash);

        setTimeout(function () {
            flash.style.display = 'none';
            resolve();
        }, 3000);
    })
}

export function errorFlash(message) {
    const flash = document.createElement('div');
    flash.id = "flash-card";

    const paragraph = document.createElement('p')
    paragraph.textContent = 'Failed : ' + message;

    flash.appendChild(paragraph);
    document.body.appendChild(flash);
}