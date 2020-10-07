import React, {useEffect, useState} from 'react';
import {connect} from 'react-redux';
import {compose} from '../../utils'
import Writer from "../../services/writer";
import {bindActionCreators} from "redux";
import {resetToLogin} from "../../actions";
import startWebsocket from "../../services/websocket";

const writerService = new Writer();

const Chat = (props) => {
    useEffect(() => {
        startWebsocket();
    }, []);

    const [message, setMessage] = useState("");

    function sendMessage() {
        if (message.length === 0) {
            return false;
        }

        writerService.send(props.token, message).then(res => {
            setMessage("");
        }).catch(err => {
            props.resetToLogin(err);
        });
    }

    return (
        <>
            <div className="card-header">
                <h4 className="card-title"><strong>Chat</strong></h4>
            </div>
            <div className="ps-container ps-theme-default ps-active-y" id="chat-content">
                <div className="media media-chat">
                    <div className="media-body" id="chat-messages">
                    </div>
                </div>
            </div>
            <div className="publisher bt-1 border-light">
                <input
                    className="publisher-input" onChange={e => setMessage(e.target.value)} value={message} type="text"
                    placeholder="Write something"/>

                <button type="button" onClick={sendMessage} className="btn-xs btn btn-success">Send</button>
            </div>
        </>
    );
};

const mapStateToProps = ({token}) => {
    return {token};
};

const mapDispatchToProps = (dispatch) => {
    return bindActionCreators({
        resetToLogin,
    }, dispatch);
};

export default compose(
    connect(mapStateToProps, mapDispatchToProps)
)(Chat);
