import LoggedInApp from './LoggedInApp'
import LoggedOutApp from './LoggedOutApp'
import { useCurrentUserQuery } from './generated-types';

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
    <LoggedOutApp />
  )
}

export default App;
