import { useParams } from "react-router-dom"
import { usePersonPageDataQuery } from "../../generated-types"

export default function PersonPage() {
  const params = useParams()
  const q = usePersonPageDataQuery({ variables: { id: params.id } })

  return (
    <div>
      <h1>{q.data?.person?.name}</h1>
      <img src={q.data?.person?.avatarUrl} />
    </div>
  )       
}
