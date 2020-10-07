const initialState = {
    screen: "login",
    token: "",
};

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case 'SET_TOKEN':
            return {
                ...state,
                screen: "chat",
                token: action.payload
            };

        case 'RESET_TO_LOGIN':
            return {
                ...state,
                screen: "login",
                token: "",
            };

        default:
            return state;
    }
};

export default reducer;