const uuid = require('uuid');

async function Login() {
    let STATE = uuid.v4();
    let REDIRECT_URL = 'http://127.0.0.1:3001';
    let CLIENT_ID;
    if (process.env.REACT_APP_CLIENT_ID)
        CLIENT_ID = process.env.REACT_APP_CLIENT_ID;
    else
        throw new Error("CLIENT_ID environment variable is not set");
    window.location.href = `https://api.intra.42.fr/oauth/authorize?client_id=${CLIENT_ID}&redirect_uri=${REDIRECT_URL}&scope=public&state=${STATE}&response_type=code`;
}

function setAll(setUser, setProjectsDoable)
{
    setUser(null);
    setProjectsDoable(null);
}

export default function Header(props) {
    const user = props.user;

    if (user === null)
        return <button onClick={Login}>Sign in</button>
    return <button onClick={() => setAll(props.setUser, props.setProjectsDoable)}>Sign out</button>
}