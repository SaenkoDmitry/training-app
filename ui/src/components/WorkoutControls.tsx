import Button from "./Button.tsx";

interface WorkoutControlsProps {
    onPrev: () => void;
    onNext: () => void;
    onFinish: () => void;
    disablePrev?: boolean;
    disableNext?: boolean;
}


export default function WorkoutControls({ onPrev, onNext, onFinish, disablePrev, disableNext }: WorkoutControlsProps) {
    return (
        <div className="controls">
            <Button
                variant="ghost"
                onClick={onPrev}
                disabled={disablePrev}
            >
                ⬅ Назад
            </Button>
            <Button
                variant="ghost"
                onClick={onNext}
                disabled={disableNext}
            >
                Вперед ➡
            </Button>
            <Button
                variant="primary"
                onClick={onFinish}
            >
                🏁 Завершить все
            </Button>
        </div>
    );
}
