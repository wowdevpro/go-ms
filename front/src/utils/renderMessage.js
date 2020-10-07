const renderMessage = data => {
    const messageData = JSON.parse(data);

    const message = document.createElement('p');
    message.innerText = messageData[0] + ": " + messageData[1];

    document.getElementById("chat-messages").appendChild(message);
}

export default renderMessage;