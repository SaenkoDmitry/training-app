import Button from "./Button";
import { useNavigate } from "react-router-dom";

type Props = {
    day: WorkoutDayTypeDTO;
    programId: number;
    onDelete: (dayId: number) => Promise<void>;
};

export default function DayCard({ day, programId, onDelete }: Props) {
    const navigate = useNavigate();

    return (
        <div
            className="card row"
            style={{ cursor: "pointer", padding: "0.5rem", display: "flex", alignItems: "center", justifyContent: "space-between" }}
        >
            <div
                onClick={() => navigate(`/programs/${programId}/days/${day.id}`)}
                style={{ flex: 1 }}
            >
                {day.name}
            </div>

            <Button
                variant="danger"
                onClick={async (e) => {
                    e.stopPropagation(); // чтобы клик по кнопке не открывал страницу дня
                    await onDelete(day.id);
                }}
            >
                🗑
            </Button>
        </div>
    );
}
