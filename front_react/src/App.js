import './App.css';
import { useState } from 'react'
import Header from './components/Header.js'

class User {
  constructor(id, name)
  {
    this.id = id;
    this.name = name;
  }
}

function App() {
  const [user, setUser] = useState(null);
  
  return (
    <div className="App">
      <Header user={user} />
      {user === null ? <div>Logged off</div>
      : <div>Logged in</div>
    }
    </div>
  );
}

export default App;
