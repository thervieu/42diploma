export default function Header(props) {
    const user = props.user;

    if (user === null)
        return <h1>Sign in</h1>
    return <h1>Signed in as {user.name}</h1>;
}