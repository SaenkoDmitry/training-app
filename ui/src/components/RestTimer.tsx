import { useEffect, useRef, useState } from "react";
import Button from "./Button";
import "../styles/RestTimer.css";

export default function RestTimer({
                                      seconds,
                                      onFinish,
                                  }: {
    seconds: number;
    onFinish?: () => void;
}) {
    const [time, setTime] = useState(seconds);
    const [running, setRunning] = useState(false);
    const [finished, setFinished] = useState(false);

    const intervalRef = useRef<number | null>(null);

    // старт / пауза
    useEffect(() => {
        if (!running) return;

        intervalRef.current = window.setInterval(() => {
            setTime((t) => t - 1);
        }, 1000);

        return () => {
            if (intervalRef.current) clearInterval(intervalRef.current);
        };
    }, [running]);

    // окончание — срабатывает только когда таймер реально идёт
    useEffect(() => {
        if (!running) return; // <--- добавили проверку
        if (time > 0) return;

        setRunning(false);
        setFinished(true);

        // вибрация
        navigator.vibrate?.([200, 100, 200]);

        onFinish?.();
    }, [time, running]);

    const toggle = () => setRunning((r) => !r);

    const reset = () => {
        setRunning(false);
        setTime(seconds);
        setFinished(false);
    };

    const format = (t: number) => {
        const m = Math.floor(t / 60);
        const s = t % 60;
        return `${m}:${s.toString().padStart(2, "0")}`;
    };

    return (
        <div className={`rest-timer ${finished ? "done" : ""}`}>
            <div className="timer-time">{format(Math.max(time, 0))}</div>

            <div className="timer-actions">
                <Button onClick={toggle}>
                    {running ? "⏸ Пауза" : "▶ Старт"}
                </Button>

                <Button variant="ghost" onClick={reset}>
                    🔄 Сброс
                </Button>
            </div>

            {finished && (
                <div className="timer-finished">Отдых закончен 💪</div>
            )}
        </div>
    );
}
