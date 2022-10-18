import { useParams } from 'react-router-dom'

export default function() {
  const { id: spaceId } = useParams()

  return <div>{spaceId}</div>
}
