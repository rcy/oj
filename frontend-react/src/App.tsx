import LoggedInApp from './LoggedInApp'
import { useCurrentUserQuery } from './generated-types';
import Button from './Button';

function App() {
  const { loading, error, data } = useCurrentUserQuery();

  console.log({ data })

  if (error) {
    return <p>{JSON.stringify(error, null, 2)}</p>
  }

  if (loading) {
    return null
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
