import logo from './logo.svg';
import './App.css';
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
    return <p>loading...</p>
  }

  if (error) {
    return <p>error</p>
  }

  if (data.currentUser) {
    return (
      <p>
        Logged in as {data.currentUser.name} <a href="/auth/logout">logout</a>
      </p>
    )
  }

  return (
    <p>
      <a href="/auth/google">login</a>
    </p>
  )
}

export default App;
