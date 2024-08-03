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

function captureFrame() {
    const video = document.querySelector("#video");
    const canvas = document.querySelector("#frameCanvas");
    const ctx = canvas.getContext("2d");

    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;

    ctx.drawImage(video, 0, 0, canvas.width, canvas.heigth);
    const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height);
    console.log("Got the frame:", imageData);
}

function sendData() {
    const ws = new WebSocket("ws://127.0.0.1:8080/ws", "protocolOne");
    const enc = new TextEncoder()
    ws.addEventListener("open", () => {
        ws.send(enc.encode("DATAAAAAAAAAAAAAAAAA INCOMING"))
    });
}

const btn = document.querySelector("#start")
btn.addEventListener("click", () => {
    sendData()
    getMedia().then(() => {
        setInterval(captureFrame, 1000);
    });
});


// TODO: send image over websocks
