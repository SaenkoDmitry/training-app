import { useEffect } from "react";

const AuthSuccess = () => {
    useEffect(() => {
        const params = new URLSearchParams(window.location.search);
        const token = params.get("token");

        if (token) {
            localStorage.setItem("token", token);
            window.location.replace("/");
        }
    }, []);

    return null;
};

export default AuthSuccess;
