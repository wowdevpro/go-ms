import React from 'react';
import {connect} from 'react-redux';
import {compose} from '../../utils'

import 'bootstrap/scss/bootstrap.scss';
import "./app.css";
import Login from "../login/login";
import Chat from "../chat/chat";

const App = (props) => {
    return (
        <>
            <div className="container">
                <div className="row d-flex justify-content-center pt-5 pb-5">
                    <div className="col-md-6">
                        <div className="card card-bordered">
                            {props.screen === "login" &&
                                <Login/>
                            }
                            {props.screen === "chat" &&
                                <Chat/>
                            }
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
};

const mapStateToProps = ({screen}) => {
    return {screen};
};

export default compose(
    connect(mapStateToProps)
)(App);
