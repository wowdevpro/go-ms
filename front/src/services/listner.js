import axios from "axios";

export default class Listener {
    _apiBase = 'http://localhost:' + (process.env.REACT_APP_LISTENER_PORT ?? "8081");

    get = async (url) => {
        return await axios.get(this._apiBase + url).then(value => {
            return value;
        });
    };

    getAllMessages = async () => {
        return await this.get("/messages/all");
    };
}