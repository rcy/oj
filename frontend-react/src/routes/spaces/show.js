import { useParams } from 'react-router-dom'

export default function() {
  const { id: spaceId } = useParams()

  // show the space members
  // show the space messages
  // customize the space with colors and images

  return <div>{spaceId}</div>
}
