import { useEffect, useState, useRef } from "react";
import { useRestTimer } from "../context/RestTimerContext";
import { useLocation, useNavigate } from "react-router-dom";
import "../styles/FloatingRestTimer.css";

// 🔹 Константы размера и позиции круга
const TIMER_SIZE = 100; // размер SVG
const RADIUS = 45;      // радиус круга
const CENTER = TIMER_SIZE / 2;
const TEXT_Y = CENTER + 6; // вертикальное положение текста

export default function FloatingRestTimer() {
    const { remaining, seconds, running } = useRestTimer();
    const location = useLocation();
    const navigate = useNavigate();

    const [position, setPosition] = useState({ x: 20, y: 100 });
    const [blinking, setBlinking] = useState(false);
    const [mounted, setMounted] = useState(false);
    const touchRef = useRef<{ startX: number; startY: number } | null>(null);

    useEffect(() => {
        setMounted(true); // для плавного появления прогресса
    }, []);

    useEffect(() => {
        const saved = localStorage.getItem("floatingTimerPosition");
        if (saved) setPosition(JSON.parse(saved));
    }, []);

    useEffect(() => {
        localStorage.setItem("floatingTimerPosition", JSON.stringify(position));
    }, [position]);

    const shouldRender = running && !location.pathname.startsWith("/sessions/");

    // Пульс на последние 5 секунд
    useEffect(() => {
        if (!shouldRender) return;
        setBlinking(remaining > 0 && remaining <= 5);
    }, [remaining, shouldRender]);

    // Вибрация по завершению
    useEffect(() => {
        if (!shouldRender) return;
        if (remaining === 0 && running) {
            navigator.vibrate?.([300, 150, 300]);
        }
    }, [remaining, running, shouldRender]);

    if (!shouldRender || seconds <= 0) return null; // защищаем от "0 секунд"

    // 🔹 расчёт прогресса
    const safeProgress = Math.max(0, Math.min(1, 1 - remaining / seconds));
    const circumference = 2 * Math.PI * RADIUS;
    const strokeOffset = mounted ? circumference * (1 - safeProgress) : circumference;

    // touch для перемещения
    const onTouchStart = (e: React.TouchEvent) => {
        const touch = e.touches[0];
        touchRef.current = { startX: touch.clientX - position.x, startY: touch.clientY - position.y };
    };
    const onTouchMove = (e: React.TouchEvent) => {
        if (!touchRef.current) return;
        const touch = e.touches[0];
        setPosition({ x: touch.clientX - touchRef.current.startX, y: touch.clientY - touchRef.current.startY });
    };
    const onTouchEnd = () => { touchRef.current = null; };

    const handleClick = () => {
        const link = localStorage.getItem("floatingTimerLink");
        if (link) navigate(link);
    };

    const minutes = Math.floor(remaining / 60);
    const secs = (remaining % 60).toString().padStart(2, "0");

    return (
        <div
            className={`floating-rest-timer ${blinking ? "blinking" : ""}`}
            style={{ top: position.y, left: position.x }}
            onTouchStart={onTouchStart}
            onTouchMove={onTouchMove}
            onTouchEnd={onTouchEnd}
            onClick={handleClick}
        >
            <svg width={TIMER_SIZE} height={TIMER_SIZE}>
                <circle r={RADIUS} cx={CENTER} cy={CENTER} />
                <circle
                    className="progress"
                    r={RADIUS}
                    cx={CENTER}
                    cy={CENTER}
                    strokeDasharray={circumference}
                    strokeDashoffset={strokeOffset}
                />
                <text x={CENTER} y={TEXT_Y} textAnchor="middle" className="timer-text">
                    {remaining > 0 ? `${minutes}:${secs}` : ""}
                </text>
            </svg>
        </div>
    );
}
