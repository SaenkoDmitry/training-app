import {api} from "./client.ts";


export const startTimer = (workoutID: number, seconds: number) =>
    api<TimerDTO>(`/api/timers/start`, {
        method: "POST",
        body: JSON.stringify({
            workout_id: workoutID,
            seconds: seconds,
        }),
    });

export const cancelTimer = (timerID: number) =>
    api(`/api/timers/cancel/${timerID}`, {
        method: "POST",
    });
