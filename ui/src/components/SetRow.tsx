import EditableValue from "./EditableValue.tsx";
import "../styles/SetRow.css";

export default function SetRow({ set, index, onDelete, onComplete, onChange }) {
    return (
        <div className={`set-row ${set.completed ? "done" : ""}`}>
            <div className="set-index">{index + 1}</div>

            <div className="set-values">
                {set.reps > 0 && (
                    <EditableValue
                        fact={set.fact_reps}
                        planned={set.reps}
                        suffix="повт."
                        completed={set.completed}
                        onSave={(v) =>
                            onChange(set.id, v, set.fact_weight, set.fact_minutes, set.fact_meters)
                        }
                    />
                )}
                {set.weight > 0 && (
                    <EditableValue
                        fact={set.fact_weight}
                        planned={set.weight}
                        suffix="кг"
                        completed={set.completed}
                        onSave={(v) =>
                            onChange(set.id, set.fact_reps, v, set.fact_minutes, set.fact_meters)
                        }
                    />
                )}
                {set.minutes > 0 && (
                    <EditableValue
                        fact={set.fact_minutes}
                        planned={set.minutes}
                        suffix="минут(ы)"
                        completed={set.completed}
                        onSave={(v) =>
                            onChange(set.id, set.fact_reps, set.fact_weight, v, set.fact_meters)
                        }
                    />
                )}
                {set.meters > 0 && (
                    <EditableValue
                        fact={set.fact_meters}
                        planned={set.meters}
                        suffix="метр(ы)"
                        completed={set.completed}
                        onSave={(v) =>
                            onChange(set.id, set.fact_reps, set.fact_weight, set.fact_minutes, v)
                        }
                    />
                )}
            </div>

            <div className="set-actions">
                <button className="icon-btn complete" onClick={onComplete}>✓</button>
                <button className="icon-btn delete" onClick={onDelete}>✕</button>
            </div>
        </div>
    );
}
