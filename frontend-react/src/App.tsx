import LoggedInApp from './LoggedInApp'
import { useCurrentUserQuery } from './generated-types';
import Button from './Button';

function App() {
  const { loading, error, data } = useCurrentUserQuery();

  if (loading) {
    return null
  }

  if (error) {
    return <p>error</p>
  }

  if (data?.currentUser) {
    return (
      <LoggedInApp />
    )
  }

  return (
    <div className="grid h-screen place-items-center">
      <a href={`/auth/login?from=${encodeURIComponent(window.location.href)}`}>
        <Button color="blue">login</Button>
      </a>
    </div>
  )
}

export default App;
