import { useParams } from 'react-router-dom'
import { useSpaceQuery } from '../../generated-types'

export default function SpaceShow() {
  const { id } = useParams()
  const { loading, data } = useSpaceQuery({ variables: { id } })

  // show the space members
  // show the space messages
  // customize the space with colors and images

  return (
    <div>
      <h1>{data?.space?.name}</h1>

      <p>
        {data?.space?.description}
      </p>
    </div>
  )
}
