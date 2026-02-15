import { useEffect, useState, useRef } from "react";
import { useRestTimer } from "../context/RestTimerContext";
import { useLocation } from "react-router-dom";

export default function FloatingRestTimer() {
    const { remaining, seconds, running } = useRestTimer();
    const location = useLocation();

    const [position, setPosition] = useState({ x: 20, y: 100 });
    const [blink, setBlink] = useState(false);

    const touchRef = useRef<{ startX: number; startY: number } | null>(null);

    // 🔹 читаем сохранённую позицию при монтировании
    useEffect(() => {
        const saved = localStorage.getItem("floatingTimerPosition");
        if (saved) setPosition(JSON.parse(saved));
    }, []);

    // 🔹 сохраняем позицию при изменении
    useEffect(() => {
        localStorage.setItem("floatingTimerPosition", JSON.stringify(position));
    }, [position]);

    const shouldRender = running && !location.pathname.startsWith("/sessions/");

    // 🔹 мигание последние 5 секунд
    useEffect(() => {
        if (!shouldRender) return;
        if (remaining <= 5 && remaining > 0) {
            const interval = setInterval(() => setBlink(prev => !prev), 500);
            return () => clearInterval(interval);
        } else {
            setBlink(false);
        }
    }, [remaining, shouldRender]);

    // 🔹 вибрация по завершению
    useEffect(() => {
        if (!shouldRender) return;
        if (remaining === 0 && running) {
            navigator.vibrate?.([300, 150, 300]);
        }
    }, [remaining, running, shouldRender]);

    if (!shouldRender) return null;

    const progress = seconds > 0 ? 1 - remaining / seconds : 0;
    const radius = 26;
    const circumference = 2 * Math.PI * radius;

    // 🔹 обработчики touch для iOS
    const onTouchStart = (e: React.TouchEvent) => {
        const touch = e.touches[0];
        touchRef.current = {
            startX: touch.clientX - position.x,
            startY: touch.clientY - position.y,
        };
    };

    const onTouchMove = (e: React.TouchEvent) => {
        if (!touchRef.current) return;
        const touch = e.touches[0];
        setPosition({
            x: touch.clientX - touchRef.current.startX,
            y: touch.clientY - touchRef.current.startY,
        });
    };

    const onTouchEnd = () => {
        touchRef.current = null;
    };

    return (
        <div
            style={{
                position: "fixed",
                top: position.y,
                left: position.x,
                zIndex: 9999,
                width: "64px",
                height: "64px",
                background: "#fff",
                borderRadius: "50%",
                boxShadow: "0 8px 24px rgba(0,0,0,0.18)",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                userSelect: "none",
                opacity: blink ? 0.4 : 1,
                transition: "opacity 0.3s",
                touchAction: "none",
            }}
            onTouchStart={onTouchStart}
            onTouchMove={onTouchMove}
            onTouchEnd={onTouchEnd}
        >
            <svg width="64" height="64">
                <circle r={radius} cx="32" cy="32" fill="none" />
                <circle
                    r={radius}
                    cx="32"
                    cy="32"
                    fill="none"
                    stroke="var(--color-primary)"
                    strokeWidth={6}
                    strokeDasharray={circumference}
                    strokeDashoffset={circumference * (1 - progress)}
                    strokeLinecap="round"
                    style={{ transition: "stroke-dashoffset 0.3s linear" }}
                />
                <text
                    x="32"
                    y="36"
                    textAnchor="middle"
                    fontSize="14"
                    fontWeight="600"
                    fill="#111"
                >
                    {remaining > 0 ? `${Math.floor(remaining / 60)}:${(remaining % 60).toString().padStart(2, "0")}` : ""}
                </text>
            </svg>
        </div>
    );
}
