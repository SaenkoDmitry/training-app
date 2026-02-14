import {useEffect, useState} from "react";
import "../styles/EditableValue.css";

type Props = {
    fact: number;       // фактическое значение
    planned?: number;   // запланированное
    suffix: string;
    typeParam: string;
    completed?: boolean;
    onSave?: (v: number) => void;
};

export default function EditableValue({
                                          fact,
                                          planned,
                                          suffix,
                                          typeParam,
                                          completed,
                                          onSave,
                                      }: Props) {
    const [localStr, setLocalStr] = useState((fact || 0) as string);

    // синхронизация локального значения при изменении fact извне
    useEffect(() => {
        setLocalStr((fact || 0) as string);
    }, [fact]);

    const handleSave = () => {
        if (onSave) {
            // если ничего не введено, подставляем план
            const valueToSend = localStr
                ? parseFloat(localStr)           // конвертируем в число
                : planned || 0;                  // если пустое, берем план или 0
            onSave(valueToSend);
        }
    };

    const matchInt = (input: string): boolean => {
        return /^\d*$/.test(input)
    }

    const matchFloat = (input: string): boolean => {
        return /^\d*\.?\d?$/.test(input)
    }

    if (completed) {
        return (
            <div className="completed-row">
                <div className="value-text"
                     style={{
                         color: fact >= planned ? "var(--color-active)" : "var(--color-danger)"
                     }}
                >
                    {!planned || planned === fact
                        ? `${fact} ${suffix}`
                        : <span>{planned} → {fact} {suffix}</span>}
                </div>
            </div>
        );
    }

    return (
        <div className="editable-wrapper">
            <input
                type="text"
                inputMode="decimal"       // цифровая клавиатура на мобильных
                value={localStr || ""}    // хранить строку, не число
                placeholder={planned?.toString()}
                className="edit-input"
                onChange={e => {
                    let value = e.target.value;

                    if (typeParam == 'int' && (!matchInt(value) || parseInt(value) <= 0)) {
                        return;
                    } else if (typeParam == 'float' && (!matchFloat(value) || parseFloat(value) <= 0)) {
                        return;
                    }
                    setLocalStr(value);  // сохраняем как строку
                }}
                onKeyDown={e => {
                    if (e.key === "Enter") {
                        setLocalStr(localStr);  // конвертируем в число при сохранении
                        handleSave();
                    }
                }}
                onBlur={() => {
                    setLocalStr(localStr);    // конвертируем в число при потере фокуса
                    handleSave();
                }}
            />

            <span className="input-suffix">{suffix}</span>
        </div>
    );
}
