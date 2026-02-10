import {api} from "./client";

export const getProgram = (programId: number) =>
    api(`/api/programs/${programId}`);

export const createDay = (programId: number, name: string) =>
    api(`/api/programs/${programId}/days/${encodeURIComponent(name)}`, {method: "POST"});

export const deleteDay = (programId: number, dayId: number) =>
    api(`/api/programs/${programId}/days/${dayId}`, {method: "DELETE"});
