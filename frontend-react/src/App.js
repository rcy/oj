import logo from './logo.svg';
import LoggedInApp from './LoggedInApp.js'
import { useQuery } from '@apollo/client';
import { CURRENT_USER } from './queries.js';

function App() {
  const { loading, error, data } = useQuery(CURRENT_USER)

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
          {data.currentUser.name}'s Family
          <div style={{ float: 'right'}}>
            <a href="/auth/logout">logout</a>
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
