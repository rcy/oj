import logo from './logo.svg';
import './App.css';
import LoggedInApp from './LoggedInApp.js'
import { useQuery, gql } from '@apollo/client';

const CURRENT_USER = gql`
  query CurrentUser {
    currentUser {
      id
      name
    }
  }
`;

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
          Logged in as {data.currentUser.name} <a href="/auth/logout">logout</a>
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
