import axios from "axios";

export default class Auth {
    _apiBase = 'http://localhost:' + (process.env.REACT_APP_AUTH_PORT ?? "8080");
    _config = {
        headers: {'Content-Type': 'application/x-www-form-urlencoded'}
    };

    post = async (url, params) => {
        return await axios.post(this._apiBase + url, params, this._config).then(value => {
            return value;
        });
    };

    login = async (login, password) => {
        const params = new URLSearchParams();
        params.append('login', login);
        params.append('password', password);

        return await this.post("/login", params);
    };
}