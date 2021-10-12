import { User } from '../App'

export default function Header(props) {
    const user = props.user;

    if (user === null)
        return <button onClick={() => props.setUser(new User('1', 'thervieu', '10.02'))}>Sign in</button>
    return <button onClick={() => props.setUser(null)}>Sign out</button>
}