import logo from './logo.svg';
import LoggedInApp from './LoggedInApp.js'
import { useQuery } from '@apollo/client';
import { CURRENT_USER } from './queries.js';

function App() {
  const { loading, error, data } = useQuery(CURRENT_USER)

  function logout(ev) {
    ev.preventDefault()
    console.log('logout')
    sessionStorage.clear()
    window.location = '/auth/logout'
  }

  if (loading) {
    return null
  }

  if (error) {
    return <p>error</p>
  }

  if (data.currentUser) {
    return (
      <>
        <header>
          <a href="/">{data.currentUser.name}'s Family</a>
          <div style={{ float: 'right'}}>
            <a href="#logout" onClick={logout}>logout</a>
          </div>
          <hr/>
        </header>
        <main>
          <LoggedInApp />
        </main>
      </>
    )
  }

  return (
    <p>
      <a href={`/auth/login?from=${encodeURIComponent(window.location)}`}>login</a>
    </p>
  )
}

export default App;
