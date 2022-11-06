import md5 from 'md5'
import { useContext, useEffect, useState } from 'react'
import Button from '../../Button'
import { PersonIdContext } from '../../contexts'
import { useCurrentFamilyMembershipQuery, useSetPersonAvatarMutation } from '../../generated-types'

// return a hash based on a given string and index
function getDailyHash(str, index) {
  return md5(`${str}${index}`)
}

const types = ['monsterid', 'wavatar', 'robohash', 'identicon', 'retro']

export default function () {
  const cfmResult = useCurrentFamilyMembershipQuery()
  const personId = useContext(PersonIdContext)
  const [dtype, selectType] = useState(types[0])
  const [mutation] = useSetPersonAvatarMutation()
  const [hashes, setHashes] = useState([])

  async function handleSelect(avatarUrl) {
    await mutation({ variables: { personId, avatarUrl }})
    cfmResult.refetch()
  }

  function shuffle() {
    setHashes([...Array(32).keys()].map((k) => getDailyHash(personId, k)))
  }

  useEffect(shuffle, [])

  return (
    <div className="flex flex-col gap-5">
      <div className='text-5xl'>Change your profile picture</div>
      <div className="flex gap-1">
        {types.map(d => (
          <Button onClick={() => selectType(d)} key={d}>{d}</Button>
        ))}
      </div>
      <div className="flex flex-wrap gap-2">
        {hashes.map(h => (
          <div key={h} className="border-solid border-4 hover:border-black">
            <GravImage h={h} d={dtype} s={80} onSelect={handleSelect} />
          </div>
        ))}
      </div>
    </div>
  )
}

function GravImage({ h, d, s, onSelect }) {
  const url = `https://www.gravatar.com/avatar/${h}?f=y&d=${d}`

  function handleClick(ev) {
    ev.preventDefault()
    return onSelect(url)
  }

  return (
    <img width={s} height={s} key={h} src={`${url}&s=${s}`} onClick={handleClick} />
  )
}
