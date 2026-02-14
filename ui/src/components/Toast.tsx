import {useEffect} from "react";

export default function Toast({message, onClose}) {
    useEffect(() => {
        const t = setTimeout(onClose, 2500);
        return () => clearTimeout(t);
    }, []);

    return <div className="toast">{message}</div>;
}
