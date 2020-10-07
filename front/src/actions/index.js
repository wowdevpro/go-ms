export const setToken = (data) => {
    return {
        type: 'SET_TOKEN',
        payload: data,
    };
};

export const resetToLogin = (data) => {
    return {
        type: 'RESET_TO_LOGIN',
        payload: data,
    };
};