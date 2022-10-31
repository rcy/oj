import { useContext, useEffect, useState } from 'react'
import Button from '../../Button'
import { PersonIdContext } from '../../contexts'
import { useCurrentFamilyMembershipQuery, useSetPersonAvatarMutation } from '../../generated-types'

function hashCode(str) {
  let hash = 0;
  for (let i = 0, len = str.length; i < len; i++) {
    let chr = str.charCodeAt(i);
    hash = (hash << 5) - hash + chr;
    hash |= 0; // Convert to 32bit integer
  }
  return hash;
}

// return a hash based on a given uuid, index that is unique per day
function getDailyHash(uuid, index) {
  const today = Math.floor(Date.now() / 1000 / 60 / 60 / 24)
  const hash = `${uuid.split('-')[0]}${today}${index}`
  console.log({ uuid, index, today, hash })
  return hash;
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
    setHashes([...Array(100).keys()].map((k) => getDailyHash(personId, k)))
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
