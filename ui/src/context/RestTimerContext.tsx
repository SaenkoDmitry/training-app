import React, {createContext, useContext, useEffect, useRef, useState} from "react";
import {cancelTimer} from "../api/timers.ts";

const STORAGE_KEY = "rest_timer_end";

type RestTimerContextType = {
    seconds: number;
    remaining: number;
    running: boolean;
    start: (seconds: number, timerID: number) => void;
    pause: () => void;
    reset: () => void;
};

const RestTimerContext = createContext<RestTimerContextType | null>(null);

export const useRestTimer = () => {
    const ctx = useContext(RestTimerContext);
    if (!ctx) throw new Error("RestTimerProvider missing");
    return ctx;
};

export const RestTimerProvider = ({children}) => {
    const [seconds, setSeconds] = useState(0);
    const [remaining, setRemaining] = useState(0);
    const [endTime, setEndTime] = useState<number | null>(null);
    const [running, setRunning] = useState(false);
    const [timerID, setTimerID] = useState(0);

    const intervalRef = useRef<number | null>(null);

    // восстановление
    useEffect(() => {
        const saved = localStorage.getItem(STORAGE_KEY);
        if (saved) {
            const parsed = Number(saved);
            if (parsed > Date.now()) {
                setEndTime(parsed);
                setRunning(true);
            }
        }
    }, []);

    useEffect(() => {
        if (!running || !endTime) return;

        intervalRef.current = window.setInterval(() => {
            const diff = Math.max(0, Math.floor((endTime - Date.now()) / 1000));
            setRemaining(diff);

            if (diff <= 0) finish();
        }, 500);

        return () => {
            if (intervalRef.current) clearInterval(intervalRef.current);
        };
    }, [running, endTime]);

    const start = (secs: number, timerID: number) => {
        const newEnd = Date.now() + secs * 1000;
        setSeconds(secs);
        setRemaining(secs);
        setEndTime(newEnd);
        setTimerID(timerID);
        setRunning(true);
        localStorage.setItem(STORAGE_KEY, String(newEnd));
    };

    const pause = () => {
        if (!endTime) return;
        const diff = Math.max(0, Math.floor((endTime - Date.now()) / 1000));
        setRemaining(diff);
        setRunning(false);
        setEndTime(null);
        localStorage.removeItem(STORAGE_KEY);
    };

    const reset = () => {
        setRunning(false);
        setEndTime(null);
        setRemaining(0);
        setTimerID(0);
        console.log("timerID", timerID);
        if (timerID != 0) {
            cancelTimer(timerID);
        }
        localStorage.removeItem(STORAGE_KEY);
    };

    const finish = () => {
        reset();
        navigator.vibrate?.([300, 150, 300]);
        localStorage.removeItem("currentTimerID");
        setTimerID(0);

        window.dispatchEvent(new Event("rest_timer_finished"));
    };

    return (
        <RestTimerContext.Provider value={{
            seconds,
            remaining,
            running,
            start,
            pause,
            reset
        }}>
            {children}
        </RestTimerContext.Provider>
    );
};
