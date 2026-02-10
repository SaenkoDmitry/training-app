import { api } from "./client";

export const getExerciseGroups = () =>
    api("/api/exercise-groups");

export const getExerciseTypesByGroup = (group: string) =>
    api(`/api/exercise-groups/${group}`);
