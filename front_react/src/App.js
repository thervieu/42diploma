import './App.css';
import { useState } from 'react'
import Header from './components/Header.js'
import UserComponent from './components/User.js'

export class User {
  constructor(id, name, xp) {
    this.id = id;
    this.name = name;
    this.xp = xp
  }
}

function App() {
  const [user, setUser] = useState(null);

  return (
    <div className="App">
      <Header user={user} setUser={setUser} />
      {user === null ? <div>Please Sign in using the top right button (42Auth).</div>
        :
        <div>
          <UserComponent user={user} />
        </div>
      }
    </div>
  );
}

export default App;
