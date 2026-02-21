import {useEffect} from "react";
import {useAuth} from "../context/AuthContext.tsx";
import {useNavigate} from "react-router-dom";

const AuthYandex = () => {
    const navigate = useNavigate();
    const { refreshUser } = useAuth();

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);

        const code = params.get("code");
        const returnedState = params.get("state");

        const savedState = localStorage.getItem("oauth_state");

        console.log("returnedState", returnedState, "savedState", savedState);
        if (!code || !returnedState || returnedState !== savedState) {
            navigate("/profile");
            return;
        }

        localStorage.removeItem("oauth_state");

        fetch("/api/yandex/login", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({ code }),
        })
            .then(res => res.json())
            .then(data => {
                localStorage.setItem("token", data.token);
                refreshUser();
                navigate("/");
            });
    }, []);

    return null;
};

export default AuthYandex;
