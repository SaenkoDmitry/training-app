import {api} from "./client.ts";

export const deleteMeasurement = (id: number) =>
    api(`/api/measurements/${id}`, {method: "DELETE"});

export const getMeasurements = (offset, limit: number) =>
    api<FindWithOffsetLimitMeasurement>(`/api/measurements?offset=${offset}&limit=${limit}`);
