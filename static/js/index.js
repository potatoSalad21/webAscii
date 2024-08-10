async function getMedia() {
    try {
        const stream = await navigator.mediaDevices.getUserMedia({
            video: true
        });
        const video = document.querySelector("#video");
        video.srcObject = stream;

        await new Promise((resolve) => {
            video.addEventListener("loadedmetadata", () => {
                resolve();
            });
        });

        console.log(video.videoWidth, "x", video.videoHeight);

    } catch (e) {
        console.error("Cannot get webcam input:", e);
        throw(e);
    }
}

async function captureFrame() {
    const video = document.querySelector("#video");
    const canvas = document.querySelector("#frameCanvas");
    const ctx = canvas.getContext("2d");

    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;
    console.log(canvas.width, "x", canvas.height)
    ctx.drawImage(video, 0, 0, canvas.width, canvas.height);

    return new Promise((resolve, reject) => {
        canvas.toBlob(blob => {
            if (blob) {
                blob.arrayBuffer().then(buffer => {
                    resolve(buffer);
                }).catch(reject);
            } else {
                reject(new Error("Failed to create blob from canvas"));
            }
        }, "image/png");
    });
}

function sendData(ws, imgData) {
    if (ws.readyState !== WebSocket.CLOSED) {
        console.log(imgData);
        ws.send(imgData);
    } else {
        console.error("ERROR: socket is not open.");
    }
}

const btn = document.querySelector("#start")
btn.addEventListener("click", () => {
    const ws = new WebSocket("ws://127.0.0.1:8080/ws");
    setInterval(() => {
        getMedia().then(() => {
            captureFrame().then((data) => {
                sendData(ws, data);
            });
        });
    }, 1000);
});
