import React, {useState} from 'react';
import {connect} from 'react-redux';
import {compose} from '../../utils'
import Auth from "../../services/auth";
import {bindActionCreators} from "redux";
import {setToken} from "../../actions";

const authService = new Auth();

const Login = (props) => {
    const [login, setLogin] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");

    function signIn() {
        if(!login || !password) {
            setError("Please fill login and password");
            return false
        }

        authService.login(login, password).then(res => {
            props.setToken(res.data);
        }).catch(err => {
            setError(err.response.data)
        });
    }

    return (
        <>
            <div className="card-header">
                <h4 className="card-title"><strong>Authentication</strong></h4>
            </div>
            <div className="container mt-3 mb-3">
                <div className="form-group">
                    <label htmlFor="exampleInputEmail1">Login</label>
                    <input type="text" onChange={e => setLogin(e.target.value)} value={login} className="form-control" id="exampleInputEmail1" />
                </div>
                <div className="form-group">
                    <label htmlFor="exampleInputPassword1">Password</label>
                    <input type="password" onChange={e => setPassword(e.target.value)} value={password} className="form-control" id="exampleInputPassword1" />
                </div>
                <div className="text-center">
                    <button type="button" onClick={signIn} className="btn btn-primary">Sign in</button>
                </div>

                {error &&
                    <p className="text-center mt-3" style={{color:"red"}}>{error}</p>
                }
            </div>
        </>
    );
};

const mapStateToProps = () => {
    return {};
};

const mapDispatchToProps = (dispatch) => {
    return bindActionCreators({
        setToken,
    }, dispatch);
};

export default compose(
    connect(mapStateToProps, mapDispatchToProps)
)(Login);
