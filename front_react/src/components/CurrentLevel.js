export default function CurrentLevel(props) {
    return (
    <div>
        <div>Logged in as {props.user.name}</div>
        <div>You are level {props.user.level}</div>
    </div> );
}