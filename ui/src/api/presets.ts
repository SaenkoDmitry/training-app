import { api } from "./client";

export const parsePreset = (preset: string) =>
    api<PresetListDTO>("/api/presets/parse", {
        method: "POST",
        body: JSON.stringify({ preset }),
    });

export const savePreset = (dayTypeId: number, newPreset: string) =>
    api("/api/presets/save", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            day_type_id: dayTypeId,
            new_preset: newPreset,
        }),
    });
