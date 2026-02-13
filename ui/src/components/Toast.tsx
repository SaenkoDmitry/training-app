import {useEffect} from "react";
import '../styles/Toast.css';

export default function Toast({message, onClose}) {
    useEffect(() => {
        const t = setTimeout(onClose, 2500);
        return () => clearTimeout(t);
    }, []);

    return <div className="toast">{message}</div>;
}
