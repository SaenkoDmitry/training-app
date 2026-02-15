import React, {useEffect, useRef, useState} from 'react';
import {useNavigate} from 'react-router-dom';
import WorkoutCard from '../components/WorkoutCard';
import {useAuth} from '../context/AuthContext';
import '../styles/App.css';
import Button from "../components/Button.tsx";
import {deleteWorkout, getWorkouts} from "../api/workouts.ts";
import {Loader, Play, Trash2} from "lucide-react";

const LIMIT = 10;

const Home: React.FC = () => {
    const {user} = useAuth();
    const [workouts, setWorkouts] = useState<Workout[]>([]);
    const [pagination, setPagination] = useState<Pagination | null>(null);
    const [loading, setLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);

    const offsetRef = useRef(0); // храним текущий offset
    const loaderRef = useRef<HTMLDivElement>(null);
    const navigate = useNavigate();

    // ---------------- DELETE WORKOUT ----------------
    const handleDelete = async (id: number) => {
        if (!confirm("Вы уверены, что хотите удалить тренировку?")) return;

        await deleteWorkout(id);
        setWorkouts(prev => prev.filter(w => w.id !== id));
    };

    // ---------------- FETCH WORKOUTS ----------------
    const fetchWorkouts = async () => {
        if (loading || !hasMore) return;

        setLoading(true);
        try {
            const nextOffset = offsetRef.current; // берем актуальный offset
            const data: ShowMyWorkoutsResult = await getWorkouts(nextOffset, LIMIT);

            setWorkouts(prev => [...prev, ...data.items]);
            setPagination(data.pagination);

            offsetRef.current += data.items.length; // обновляем offset
            setHasMore(offsetRef.current < data.pagination.total);
        } finally {
            setLoading(false);
        }
    };

    // ---------------- INFINITE SCROLL ----------------
    useEffect(() => {
        if (!loaderRef.current || !hasMore) return;

        const observer = new IntersectionObserver((entries) => {
            if (entries[0].isIntersecting && !loading) {
                fetchWorkouts();
            }
        });

        observer.observe(loaderRef.current);
        return () => observer.disconnect();
    }, [user, hasMore, loading]);

    const isEmpty = pagination && pagination.total === 0;

    return (
        <div className="page stack">
            <h1>Мои тренировки</h1>

            <Button
                variant="active"
                onClick={() => navigate('/start')}
            >
                <Play/> Начать новую тренировку
            </Button>

            {!isEmpty && <div style={{display: 'flex', flexDirection: 'column', gap: 12}}>
                {workouts.map((w, idx) => (
                    <div
                        key={w.id}
                        onClick={() => navigate(`/workouts/${w.id}`)}
                        className="workout-item"
                    >
                        <WorkoutCard w={w} idx={idx + 1}/>

                        <div className="workout-actions">
                            {!w.completed && (
                                <Button
                                    variant="active"
                                    onClick={(e) => {
                                        navigate(`/sessions/${w.id}`);
                                        e.stopPropagation();
                                    }}
                                >
                                    <Play size={14}/>
                                </Button>
                            )}

                            <Button
                                variant="danger"
                                onClick={(e) => {
                                    e.stopPropagation();
                                    handleDelete(w.id);
                                }}
                            >
                                <Trash2 size={14}/>
                            </Button>
                        </div>
                    </div>
                ))}
            </div>}

            {loading && <Loader/>}

            {isEmpty && (
                <div>
                    <div style={{marginTop: 18, fontSize: 18}}>У вас пока нет ни одной тренировки.</div>
                    <h2>Пора начать! 💪</h2>
                </div>
            )}

            {/* IntersectionObserver смотрит сюда */}
            <div ref={loaderRef} style={{height: 20}}/>

            {pagination && pagination.total > 0 && (
                <p>
                    {Math.min(offsetRef.current, pagination.total)} из {pagination.total}
                </p>
            )}
        </div>
    );
};

export default Home;
