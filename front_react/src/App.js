import './App.css';
import background from './assets/background.png';
import { useEffect, useState } from 'react'
import { useLocation, useHistory } from 'react-router-dom'
import Header from './components/Header.js'
import CurrentLevel from './components/CurrentLevel.js'
import ProjectsForm from './components/ProjectsForm.js'
import ProjectsList from './components/ProjectsList.js'
import Calculate from './components/Calculate.js'
import Box from '@material-ui/core/Box';
import useWindowDimensions from './hooks/WindowHook.js'

export class User {
  constructor(name, level) {
    this.name = name;
    this.level = level;
  }
}

export class Project {
  constructor(name, xp, percentage, checked) {
    this.name = name;
    this.xp = xp;
    this.percentage = percentage;
    this.checked = checked;
  }
}

async function process42ApiRedirect(code) {
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

  const jsonData = await response.json();

  return jsonData;
}

async function set42User(setUser, setProjectsDoable, code) {
  let UserData = await process42ApiRedirect(code);

  if (UserData) {
    setUser(new User(UserData.login, UserData.level));
    setProjectsDoable(UserData.projects);
  }
}

function App() {
  const { height, width } = useWindowDimensions();
  const [user, setUser] = useState(null);
  const [projects, setProjects] = useState(null);
  const [projectsDoable, setProjectsDoable] = useState(null);
  let history = useHistory();
  const { search } = useLocation();

  useEffect(() => {
    var searchParams = new URLSearchParams(search);

    // if we catch an auth redirect from 42 api
    let code = searchParams.get("code");
    if (code) {
      set42User(setUser, setProjectsDoable, code);
      history.replace("/");
    }
  }, [search, history]);

  return (
    // <div className="App" style={{ backgroundImage: `url(${background})` }}>
    <div className="App">
      <Header user={user} setUser={setUser} setProjectsDoable={setProjectsDoable} />
      {user === null ? <div> Please Sign in using the top right button (42Auth). </div>
        :
        <div>
          <CurrentLevel user={user} />
          {projectsDoable !== null ?
            <div>
              <ProjectsForm user={user} projects={projects} setProjects={setProjects}
                projectsDoable={projectsDoable} />
              {projects !== null && projects.length > 0 ?
                <Box
                  sx={{
                    display: 'grid',
                    gridTemplateColumns: 'repeat(2, 1fr)'
                  }}
                >
                  <ProjectsList projects={projects} setProjects={setProjects} />
                  <Calculate level={user.level} projects={projects} />
                </Box>
                :
                <div> No projects chosen </div>}
            </div>
            :
            <div> No projects for now </div>}
        </div>
      }
    </div>
  );
}

export default App;