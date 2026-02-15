import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { GetActiveProgramForUser } from "../api/programs";
import { startWorkout } from "../api/workouts";
import DayTypeCard from "../components/DayTypeCard";
import "../styles/workout.css";
import {FolderKanban, Loader} from "lucide-react";
import Button from "../components/Button.tsx";

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

    const isEmpty = days.length === 0;

    return (
        <div className="page stack">
            <h1>Выбор дня</h1>

            {isEmpty && <div>
                <div style={{marginBottom: 24}}>Сначала настройте программу и добавьте дни!</div>

                <Button variant={"active"} onClick={() => {
                    navigate(`/programs`);
                }}><FolderKanban size={18}/>Настроить программу</Button>
            </div>}

            {loading && <Loader/>}

            {!isEmpty && days.map(day => (
                <DayTypeCard key={day.id} day={day} onClick={() => handleStart(day.id)} />
            ))}
        </div>
    );
}
