import { useParams } from "react-router-dom"
import { usePersonPageDataQuery } from "../../generated-types"

export default function PersonPage() {
  const params = useParams()
  const q = usePersonPageDataQuery({ variables: { id: params.id } })

  return (
    <div className="p-2 flex items-center">
      <img src={q.data?.person?.avatarUrl} />
      <h1 className="px-2 text-6xl">{q.data?.person?.name}</h1>
    </div>
  )       
}
