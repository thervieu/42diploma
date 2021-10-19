import './App.css';
import { useEffect, useState } from 'react'
import { useLocation, useHistory } from 'react-router-dom'
import Header from './components/Header.js'
import CurrentLevel from './components/CurrentLevel.js'
import Projects from './components/Projects.js'
import BarChart from './components/BarChart.js'

export class User {
  constructor(name, xp) {
    this.name = name;
    this.xp = xp;
  }
}

export class Project {
  constructor(name, xp) {
    this.name = name;
    this.xp = xp;
    this.percentage = 100;
  }
}

async function process42ApiRedirect(code){
  const data = {
    code: code
  };
  const response = await fetch("http://localhost:3000/auth", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });
  if (!response.ok)
    return null;
  console.log("response");
  console.log(response);
  console.log("response after");
  const jsonData = await response.json();
  console.log("jsonData");
  console.log(jsonData);
  return jsonData;
}

async function set42User(setUser, setProjectsDone, code) {
  let UserData = await process42ApiRedirect(code);

  if (UserData) {
    setUser(new User(UserData.Login, UserData.Level));
    setProjectsDone(UserData.Projects);
  }
}

function App() {
  const [user, setUser] = useState(null);
  const [level, setLevel] = useState(null);
  const [projectsDone, setProjectsDone] = useState(null);
  let history = useHistory();
  const { search } = useLocation();

  useEffect(() => {
    var searchParams = new URLSearchParams(search);

    // if we catch an auth redirect from 42 api
    let code = searchParams.get("code");
    if (code) {
      set42User(setUser, setProjectsDone, code);
      history.replace("/");
    }
  }, [search, history]);

  return (
    <div className="App">
      <Header user={user} setUser={setUser} />
      {user === null ? <div> Please Sign in using the top right button (42Auth). </div>
        :
        <div>
          <CurrentLevel user={user} />
          <Projects user={user} setLevel={setLevel} projectsDone={projectsDone} />
          <BarChart user={user} level={level} />
        </div>
      }
    </div>
  );
}

export default App;