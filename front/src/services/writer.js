import axios from "axios";

export default class Writer {
    _apiBase = 'http://' + (process.env.REACT_APP_HOST_NAME ?? "localhost") + ':' + (process.env.REACT_APP_WRITER_PORT ?? "8082");
    _config = {
        headers: {'Content-Type': 'application/x-www-form-urlencoded'}
    };

    post = async (url, params) => {
        return await axios.post(this._apiBase + url, params, this._config).then(value => {
            return value;
        });
    };

    send = async (token, message) => {
        const params = new URLSearchParams();
        params.append('token', token);
        params.append('message', message);

        return await this.post("/message/new", params);
    };
}