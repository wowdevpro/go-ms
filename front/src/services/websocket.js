import {renderMessage} from "../utils";
import Listener from "./listner";

const startWebsocket = () => {
    const connection = new WebSocket("ws://localhost:" + (process.env.REACT_APP_LISTENER_PORT ?? "8081") + "/ws");
    connection.onmessage = function (event) {
        renderMessage(event.data);
    };
    connection.onopen = function () {
        const listenerService = new Listener();

        listenerService.getAllMessages().then(res => {
            if(res.data != null) {
                for (let message of res.data) {
                    renderMessage(message);
                }
            }
        });
    };
}

export default startWebsocket;