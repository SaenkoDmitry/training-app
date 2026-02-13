import {useEffect, useState} from "react";
import "../styles/EditableValue.css";

type Props = {
    fact: number;       // факт
    planned?: number;   // план
    suffix: string;
    completed?: boolean;
    onSave?: (v: number) => void;
};

export default function EditableValue({
                                          fact,
                                          planned,
                                          suffix,
                                          completed,
                                          onSave,
                                      }: Props) {
    const [local, setLocal] = useState(fact || 0);

    // если fact меняется извне, синхронизируем локальный state
    useEffect(() => {
        setLocal(fact || 0);
    }, [fact]);

    // ===== после выполнения просто текст =====
    if (completed) {
        if (!planned || planned === fact) {
            return <span className="value-text">{fact} {suffix}</span>;
        }

        return (
            <span className="value-text done-highlight">
                {planned} → {fact} {suffix}
            </span>
        );
    }

    const handleSave = () => {
        if (onSave) onSave(local);
    };

    return (
        <div className="editable-wrapper">
            <input
                type="number"
                value={local || ""}
                placeholder={planned?.toString()}
                className="edit-input"
                onChange={e => setLocal(+e.target.value)}
                onKeyDown={e => {
                    if (e.key === "Enter") handleSave();
                }}
                onBlur={handleSave}
            />
            <span className="input-suffix">{suffix}</span>
        </div>
    );
}
