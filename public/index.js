const username = document.getElementById("username")
const crypto = document.getElementById("crypto")
const message = document.getElementById("message")
message.addEventListener('keypress', enter)
const output = document.getElementById("output")

const socket = new WebSocket("ws://localhost:8080/ws")

socket.onopen = () => output.innerHTML += "<p style='color:green;'>Status: Connected</p>\n"

socket.onmessage = e => {
    data = JSON.parse(e.data)

    const m = CryptoJS.AES.decrypt(data.message, crypto.value)
    let decryptedMessage = m.toString(CryptoJS.enc.Utf8)
    
    if (decryptedMessage === '') decryptedMessage = '******* encrypted message ********'

    output.innerHTML += "<p>@" + (data.username ? data.username : "anonymous") + ": " + decryptedMessage + "</p>"
    scrollToBottom()
}

function send() {

    const m = CryptoJS.AES.encrypt(message.value, crypto.value)
    const encryptedMessage = m.toString()

    socket.send(JSON.stringify({username: username.value, message: encryptedMessage}))
    message.value = ""
}

function scrollToBottom() {
    const obj = document.getElementById("output")
    if(!obj) return
    obj.scrollTop = obj.scrollHeight
}

function enter(e) {
    if (e.keyCode === 13) send()
}

        