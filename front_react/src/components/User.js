import { User } from '../App'

export default function UserComponent(props) {
    const user = props.user;

    return (
    <div>
        <div>Logged in as {user.name}</div>
        <div>You are level {user.xp}</div>
    </div> );
}