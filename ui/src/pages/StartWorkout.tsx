import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { GetActiveProgramForUser } from "../api/programs";
import { startWorkout } from "../api/workouts";
import DayTypeCard from "../components/DayTypeCard";
import "../styles/workout.css";
import {Loader} from "lucide-react";

export default function StartWorkout() {
    const [days, setDays] = useState<WorkoutDayTypeDTO[]>([]);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        const load = async () => {
            const programs = await GetActiveProgramForUser();
            setDays(programs.day_types);

            setLoading(false);
        };

        load();
    }, []);

    const handleStart = async (dayId: number) => {
        const res = await startWorkout(dayId);
        navigate(`/sessions/${res.workout_id}`);
    };

    return (
        <div className="page stack">
            <h1>Выбери день</h1>

            {loading && <Loader/>}

            {days.map(day => (
                <DayTypeCard key={day.id} day={day} onClick={() => handleStart(day.id)} />
            ))}
        </div>
    );
}
