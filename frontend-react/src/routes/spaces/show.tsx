import { useParams } from 'react-router-dom'
import { useSpaceQuery } from '../../generated-types'
import Chat from '../../components/Chat'
import MembersList from './MembersList'

export default function SpaceShow() {
  const { id } = useParams()
  //const personId = useContext(PersonIdContext)
  const spaceQueryResult = useSpaceQuery({ variables: { id } })

  if (spaceQueryResult.error) {
    return <div>{JSON.stringify(spaceQueryResult.error)}</div>
  }

  // show the space members
  // show the space messages
  // customize the space with colors and images

  return (
    <div>
      <header className=''>
        <h1 className='text-7xl bg-yellow-300'>{spaceQueryResult.data?.space?.name}</h1>
        <p className='px-5 bg-orange-300'>
          {spaceQueryResult.data?.space?.description}
        </p>
      </header>

      <MembersList spaceId={id} />

      <hr/>

      <div className='flex justify-around px-5'>
        {id && <Chat spaceId={id} />}
      </div>
    </div>
  )
}
