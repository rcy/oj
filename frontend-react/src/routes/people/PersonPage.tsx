import { useParams } from "react-router-dom"
import { usePersonPageDataQuery } from "../../generated-types"
import ChatSection from "./ChatSection";

export default function PersonPage() {
  const params = useParams()
  const q = usePersonPageDataQuery({ variables: { id: params.id } })
  const pagePerson = q.data?.person;

  return (
    <div className="flex flex-col">
      <header className="p-2 flex items-center">
        <img src={pagePerson?.avatarUrl} />
        <h1 className="px-2 text-6xl">{pagePerson?.name}</h1>
      </header>

      <ChatSection pagePerson={pagePerson} />
    </div>
  )
}
