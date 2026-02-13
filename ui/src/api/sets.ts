import {api} from "./client";

export const addSet = (exerciseID: number) =>
    api(`/api/sets/${exerciseID}`, {
        method: "POST"
    });

export const deleteSet = (id: number) =>
    api(`/api/sets/${id}`, {
        method: "DELETE"
    });

export const completeSet = (id: number) =>
    api(`/api/sets/${id}/complete`, {
        method: "POST",
    });

export const changeSet = (id: number, reps, weight, minutes, meters: number) => // обычно либо reps!=0&weight!=0, либо minutes!=0, либо meters!=0
    api(`/api/sets/${id}/change`, {
        method: "POST",
        body: JSON.stringify({
            fact_reps: reps,
            fact_weight: weight,
            fact_minutes: minutes,
            fact_meters: meters,
        }),
    });
