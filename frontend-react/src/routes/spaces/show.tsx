import { useParams } from 'react-router-dom'
import { useSpaceMembershipsBySpaceIdQuery, useSpaceQuery } from '../../generated-types'
import Chat from '../../components/Chat'
import Debug from '../../components/Debug'
import MembersList from './MembersList'
import { useContext } from 'react'
import { PersonIdContext } from '../../contexts'

export default function SpaceShow() {
  const { id } = useParams()
  //const personId = useContext(PersonIdContext)
  const spaceQueryResult = useSpaceQuery({ variables: { id } })

  // show the space members
  // show the space messages
  // customize the space with colors and images

  return (
    <div>
      <header className='pb-5'>
        <h1 className='p-5 text-5xl bg-yellow-300'>{spaceQueryResult.data?.space?.name}</h1>
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
