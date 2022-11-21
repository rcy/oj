import { Link } from "react-router-dom";
import MySpaceMemberships from "./MySpaceMemberships";

export default function SpacesIndex() {
  return (
    <div className="flex flex-col">
      <div className="flex bg-green-200">
        <div className="text-6xl">Explore</div>
      </div>
      <section className="pb-10">
        <MySpaceMemberships />
        <Link className="text-blue-600" to="/spaces/explore">
          explore
        </Link>
      </section>
    </div>
  );
}
