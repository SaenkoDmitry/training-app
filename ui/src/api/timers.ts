import {api} from "./client.ts";


export const startTimer = (workoutID: number, seconds: number) =>
    api(`/api/timers/start`, {
        method: "POST",
        body: JSON.stringify({
            workout_id: workoutID,
            seconds: seconds,
        }),
    });
