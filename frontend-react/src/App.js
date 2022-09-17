import logo from './logo.svg';
import LoggedInApp from './LoggedInApp.js'
import { useQuery } from '@apollo/client';
import { CURRENT_USER } from './queries.js';
import Button from './Button.js';

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
      <LoggedInApp />
    )
  }

  return (
    <div className="grid h-screen place-items-center">
      <a href={`/auth/login?from=${encodeURIComponent(window.location)}`}>
        <Button>login</Button>
      </a>
    </div>
  )
}

export default App;
