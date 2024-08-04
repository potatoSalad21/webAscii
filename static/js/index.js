async function getMedia() {
    try {
        const stream = await navigator.mediaDevices.getUserMedia({
            video: true
        });
        const video = document.querySelector("video");
        video.srcObject = stream;
    } catch (e) {
        console.error("Cannot get webcam input:", e);
    }
}

async function captureFrame() {
    const video = document.querySelector("#video");
    const canvas = document.querySelector("#frameCanvas");
    const ctx = canvas.getContext("2d");

    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;

    ctx.drawImage(video, 0, 0, canvas.width, canvas.heigth);

    return new Promise((resolve) => {
        canvas.toBlob((blob) => {
            blob.arrayBuffer().then((buffer) => {
                resolve(buffer);
            });
        }, "image/png");
    });
}

function sendData(imgData) {
    const ws = new WebSocket("ws://127.0.0.1:8080/ws");

    ws.addEventListener("open", () => {
        ws.send(imgData);
    });
}

const btn = document.querySelector("#start")
btn.addEventListener("click", () => {
    getMedia().then(() => {
        setInterval(async () => {
            const data = await captureFrame();
            if (data) {
                sendData(data);
            } else {
                console.error("Couldn't capture frame");
            }
        }, 1000);
    });
});
