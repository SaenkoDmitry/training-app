type Props = {
    day: WorkoutDayTypeDTO;
    onClick: () => void;
};

export default function DayTypeCard({ day, onClick }: Props) {
    return (
        <div onClick={onClick} className={"stack card"}>
            <h3>{day.name}</h3>
            {/*<p>{day.preset}</p>*/}
        </div>
    );
}
